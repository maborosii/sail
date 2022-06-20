package helm

// helm install ale-case-service . --set image.repository=harbor.minstone.com:5002/app/ale-case-service,image.tag=1.2.4.2  -n qa

import (
	"context"
	"io"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/release"
)

func RunInstall(ctx context.Context, args []string, client *action.Install, values map[string]interface{}, ns string, out io.Writer) (*release.Release, error) {
	debug("Original chart version: %q", client.Version)
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
	cp, err := client.ChartPathOptions.LocateChart(chart, nil)
	if err != nil {
		return nil, err
	}

	debug("CHART PATH: %s\n", cp)

	if err != nil {
		return nil, err
	}

	// Check chart dependencies to make sure all are present in /charts
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

	client.Namespace = ns
	return client.RunWithContext(ctx, chartRequested, values)
}
