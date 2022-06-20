package helm

import (
	"fmt"
	"io"

	"helm.sh/helm/v3/pkg/action"
)

func RunUninstall(app string, client *action.Uninstall, out io.Writer) error {
	res, err := client.Run(app)
	if err != nil {
		return err
	}
	if res != nil && res.Info != "" {
		fmt.Fprintln(out, res.Info)
	}

	fmt.Fprintf(out, "release \"%s\" uninstalled\n", app)
	return nil
}
