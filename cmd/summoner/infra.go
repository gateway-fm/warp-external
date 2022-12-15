package summoner

import (
	"github.com/spf13/cobra"

	"github.com/gateway-fm/warp_external/internal"
)

// Cmd is
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "infra-summon",
		Short: "Generate infrastructure files",

		RunE: func(cmd *cobra.Command, args []string) error {
			if err := internal.SummonNewInfra(); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
