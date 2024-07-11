package crawler

import (
	"context"

	"github.com/spf13/cobra"
)

type cli struct {
	rootCmd *cobra.Command
	cmd     *cobra.Command
}

type CliPort interface {
	Execute(context.Context) error
}

func NewCli(cc CliConfigPort) CliPort {
	return &cli{
		rootCmd: cc.GetRootCmd(),
		cmd:     cc.GetCmd(),
	}
}

func (c *cli) Execute(ctx context.Context) error {
	c.rootCmd.AddCommand(c.cmd)
	if err := c.rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
