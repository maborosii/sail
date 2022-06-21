package helm

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"

	cm "sail/common"
	"sail/global"
	bd "sail/internal/handler/deployer"
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

func (dh *DeployerByHelm) Install(ctx context.Context, b cm.Message) error {
	ctxTimeOut, cancel := context.WithTimeout(ctx, 120*time.Second)

	commonArgs := []string{b.GetName(), b.GetChart()}
	defer cancel()
	release, err := RunInstall(ctxTimeOut, commonArgs, dh.settings, dh.InstallCli)
	if err != nil {
		global.Logger.Error("app install failed", zap.String("app", release.Name), zap.Error(err))
	}
	global.Logger.Info("app install successful", zap.String("app", release.Name))
	return nil

}

func (dh *DeployerByHelm) Uninstall(b cm.Message) error {
	err := RunUninstall(b.GetName(), dh.UninstallCli)
	if err != nil {
		global.Logger.Error("app uninstall occur error", zap.Error(err))
	}
	fmt.Println(b.GetName(), "uninstall successful")
	return nil
}

func (dh *DeployerByHelm) Run(ctx context.Context) {
	var (
		ok      bool
		Message cm.Message
		err     error
	)

DEPLOY_LOOP:
	for {
		select {
		case <-ctx.Done():
			break DEPLOY_LOOP
		case Message, ok = <-dh.InChan:
			if !ok {
				break DEPLOY_LOOP
			}
		}
		// main logic
		// TODO：需要解耦
		err = dh.Uninstall(Message)
		if err != nil {
			panic(err)
		}
		err = dh.Install(ctx, Message)
		if err != nil {
			panic(err)
		}
		dh.OutChan <- Message
	}
}
