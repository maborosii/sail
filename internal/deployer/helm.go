package deployer

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"

	cm "sail/common"
	"sail/global"
)

type HelmCfg struct {
	// 默认参数
	settings *cli.EnvSettings
	cfg      *action.Configuration

	InstallCli   *action.Install
	UninstallCli *action.Uninstall
}

type DeployerByHelm struct {
	*BaseDeployer
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

func (dh *DeployerByHelm) Install(ctx context.Context, b cm.InMessage) error {
	name, err := b.GetName()
	if err != nil {
		global.Logger.Error("get app name occur error, info:", zap.Error(err))
		return err
	}

	chart, err := b.GetResource()
	if err != nil {
		global.Logger.Error("get chart name occur error, info:", zap.Error(err))
		return err
	}
	commonArgs := []string{name, chart}

	ctxTimeOut, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	release, err := RunInstall(ctxTimeOut, commonArgs, dh.settings, dh.InstallCli)
	if err != nil {
		global.Logger.Error("app install failed", zap.String("app", release.Name), zap.Error(err))
	}
	global.Logger.Info("app install successful", zap.String("app", release.Name))
	return nil

}

func (dh *DeployerByHelm) Uninstall(b cm.InMessage) error {
	name, err := b.GetName()
	if err != nil {
		global.Logger.Error("get app name occur error, info:", zap.Error(err))
		return err
	}
	err = RunUninstall(name, dh.UninstallCli)
	if err != nil {
		global.Logger.Error("app uninstall occur error", zap.Error(err))
	}
	global.Logger.Info("app uninstall successful", zap.String("app", name))
	return nil
}

func (dh *DeployerByHelm) Run(ctx context.Context) {
	var (
		ok      bool
		Message cm.InMessage
		err     error
	)

DEPLOY_LOOP:
	for {
		select {
		case <-ctx.Done():
			break DEPLOY_LOOP
		case Message, ok = <-dh.inChan:
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
		dh.outChan <- Message
	}
}

func RunInstall(ctx context.Context, args []string, settings *cli.EnvSettings, client *action.Install) (*release.Release, error) {
	global.Logger.Debug("Original chart version: ", zap.String("version", client.Version))
	if client.Version == "" && client.Devel {
		global.Logger.Debug("setting version to >0.0.0-0")
		// debug("setting version to >0.0.0-0")
		client.Version = ">0.0.0-0"
	}

	// 解析args中的app name和chart name
	name, chart, err := client.NameAndChart(args)
	if err != nil {
		return nil, err
	}
	client.ReleaseName = name

	// 若chart本地存在，则返回该chart的绝对路径
	// 若本地不存在，则远端拉取
	cp, err := client.ChartPathOptions.LocateChart(chart, settings)
	if err != nil {
		return nil, err
	}
	global.Logger.Debug("CHART PATH: ", zap.String("path", cp))

	if err != nil {
		return nil, err
	}
	chartRequested, err := loader.Load(cp)
	if err != nil {
		return nil, err
	}

	// 检查chart是否可安装
	if err := checkIfInstallable(chartRequested); err != nil {
		return nil, err
	}

	if chartRequested.Metadata.Deprecated {
		global.Logger.Warn("setting version to >0.0.0-0")
		// warning("This chart is deprecated")
	}

	client.Namespace = settings.Namespace()

	// 传入空的命令行values
	return client.RunWithContext(ctx, chartRequested, make(map[string]interface{}))
}

func RunUninstall(app string, client *action.Uninstall) error {
	res, err := client.Run(app)
	if err != nil {
		// global.Logger.Error("app uninstall occur error", zap.Error(err))
		return err
	}
	if res != nil && res.Info != "" {
		// fmt.Fprintln(out, res.Info)
		global.Logger.Info("app uninstall info", zap.String("info", res.Info))
	}

	// fmt.Fprintf(out, "release \"%s\" uninstalled\n", app)
	global.Logger.Info("app uninstall info", zap.String("app", app))
	return nil
}

// func debug(format string, v ...interface{}) {
// 	if true {
// 		format = fmt.Sprintf("[debug] %s\n", format)
// 		log.Output(2, fmt.Sprintf(format, v...))
// 	}
// }

// func warning(format string, v ...interface{}) {
// 	format = fmt.Sprintf("WARNING: %s\n", format)
// 	fmt.Fprintf(os.Stderr, format, v...)
// }
func checkIfInstallable(ch *chart.Chart) error {
	switch ch.Metadata.Type {
	case "", "application":
		return nil
	}
	return errors.Errorf("%s charts are not installable", ch.Metadata.Type)
}
