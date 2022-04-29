package cmd

import (
	"github.com/goo-app/cli/internal"
	"github.com/spf13/cobra"
)

var (
	use = &cobra.Command{
		Use:   "use [app-name] [version]",
		Short: "Use application version (when already installed)",
		RunE:  doUse,
	}
)

func init() {
	root.AddCommand(use)
}

func doUse(_ *cobra.Command, argv []string) error {
	app, err := internal.GetApplication(argv[0])
	if err != nil {
		return err
	}
	return app.Use(argv[1])
}
