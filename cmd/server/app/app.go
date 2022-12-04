package app

import (
	"github.com/spf13/cobra"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/term"
	"k8s.io/klog/v2"

	"github.com/gaulzhw/go-server/cmd/server/app/options"
	"github.com/gaulzhw/go-server/pkg/features"
)

func NewServerCommand() *cobra.Command {
	opts := options.NewOptions()

	cmd := &cobra.Command{
		Use: "go-server",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliflag.PrintFlags(cmd.Flags())

			if err := opts.Complete(); err != nil {
				klog.Errorf("unable to complete options, %+v", err)
				return err
			}

			if err := opts.Validate(); err != nil {
				klog.Errorf("unable to validate options, %+v", err)
				return err
			}

			if err := run(opts); err != nil {
				klog.Errorf("unable to run server, %+v", err)
				return err
			}

			return nil
		},
	}

	fs := cmd.Flags()
	namedFlagSets := opts.Flags()
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	features.SetDefaultFeatureGates()

	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, namedFlagSets, cols)

	return cmd
}

func run(opts *options.Options) error {
	return nil
}
