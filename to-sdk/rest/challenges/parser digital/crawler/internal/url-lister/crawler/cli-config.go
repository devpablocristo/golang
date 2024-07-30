package crawler

import (
	"context"

	"github.com/spf13/cobra"
)

type CliConfig struct {
	rootCmd *cobra.Command
	cmd     *cobra.Command
	crawler CrawlerPort
}

type CliConfigPort interface {
	GetRootCmd() *cobra.Command
	GetCmd() *cobra.Command
}

func NewCliConfig(ctx context.Context, cs CrawlerPort) CliConfigPort {
	cc := &CliConfig{
		crawler: cs,
	}

	rc := &cobra.Command{
		Use: "crawl",
		Run: func(cmd *cobra.Command, args []string) {
			startURL := args[0]
			cc.crawler.Crawl(ctx, startURL)
		},
	}

	cd := &cobra.Command{
		Use:   "crawl",
		Short: "Start web Crawlering",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			startURL := args[0]
			cc.crawler.Crawl(ctx, startURL)
		}}

	cc.rootCmd = rc
	cc.cmd = cd

	return cc

}

func (c *CliConfig) GetRootCmd() *cobra.Command {
	return c.rootCmd
}

func (c *CliConfig) GetCmd() *cobra.Command {
	return c.cmd
}
