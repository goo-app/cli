package cmd

import (
	"github.com/goo-app/cli/api"
	"github.com/spf13/cobra"
)

var (
	upgrade = &cobra.Command{
		Use:   "upgrade [app-names...]",
		Short: "Upgrade specific applications",
		RunE:  doUpgrade,
	}
)

func init() {
	root.AddCommand(upgrade)
}

func doUpgrade(_ *cobra.Command, argv []string) error {
	return api.Upgrade(argv...)
}
