package helm

import (
	"context"
	"fmt"
	"log"
	"os"
	bd "sail/internal/handler/deployer"
	"time"

	// "go.uber.org/zap"
	gutils "github.com/Laisky/go-utils"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

type HelmCfg struct {
	// 默认参数
	settings *cli.EnvSettings
	cfg      *action.Configuration

	InstallCli   *action.Install
	UninstallCli *action.Uninstall
}

type DeployerByHelm struct {
	*bd.BaseDeployer
	*HelmCfg
}

func (h *HelmCfg) SetSetting(ns string) {
	h.settings = cli.New()
	h.settings.SetNamespace(ns)
}
func (h *HelmCfg) SetCfg(ns string) {
	h.cfg = new(action.Configuration)
	if err := h.cfg.Init(h.settings.RESTClientGetter(), ns, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		gutils.Logger.Fatal(err)
		log.Printf("%+v", err)
		os.Exit(1)
	}
}

func (h *HelmCfg) SetInstallClient() {
	h.InstallCli = action.NewInstall(h.cfg)

}
func (h *HelmCfg) SetUninstallClient() {
	h.UninstallCli = action.NewUninstall(h.cfg)
}

func (dh *DeployerByHelm) Install(ctx context.Context, b images.BuildInfo) error {
	ctxTimeOut, cancel := context.WithTimeout(ctx, 120*time.Second)

	defer cancel()

	commonArgs := []string{b.GetName(), b.GetComposePath()}
	installArgs := map[string]interface{}{
		"image.repository": fmt.Sprintf("%s/%s/%s", b.GetRepository(), b.GetProject(), b.GetName()),
		"image.tag":        b.GetImageTags()[0],
	}

	release, err := customhelm.RunInstall(ctxTimeOut, commonArgs, dh.InstallCli, installArgs, dh.settings.Namespace(), os.Stdout)
	if err != nil {
		fmt.Println(err)
		panic("install failed")
	}
	fmt.Println(release.Name, "install successful")
	return nil

}

func (dh *DeployerByHelm) Uninstall(b images.BuildInfo) error {
	err := customhelm.RunUninstall(b.GetName(), dh.UninstallCli, os.Stdout)
	if err != nil {
		panic("uninstall failed")
	}
	fmt.Println(b.GetName(), "uninstall successful")
	return nil

}

func (dh *DeployerByHelm) Run(ctx context.Context) {
	var (
		ok        bool
		buildInfo images.BuildInfo
		err       error
	)

DEPLOY_LOOP:
	for {
		select {
		case <-ctx.Done():
			break DEPLOY_LOOP
		case buildInfo, ok = <-dh.inChan:
			if !ok {
				break DEPLOY_LOOP
			}
		}
		// main logic
		// TODO：需要解耦
		err = dh.Uninstall(buildInfo)
		if err != nil {
			panic(err)
		}
		err = dh.Install(ctx, buildInfo)
		if err != nil {
			panic(err)
		}
		dh.outChan <- buildInfo
	}
}
