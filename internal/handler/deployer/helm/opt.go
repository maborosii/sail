package helm

// helm install ale-case-service . --set image.repository=harbor.minstone.com:5002/app/ale-case-service,image.tag=1.2.4.2  -n qa

import (
	"context"
	"sail/global"

	"go.uber.org/zap"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

func RunInstall(ctx context.Context, args []string, settings *cli.EnvSettings, client *action.Install) (*release.Release, error) {
	global.Logger.Debug("Original chart version: ", zap.String("version", client.Version))
	if client.Version == "" && client.Devel {
		debug("setting version to >0.0.0-0")
		client.Version = ">0.0.0-0"
	}

	// 解析args中的name和路径
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
		warning("This chart is deprecated")
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
