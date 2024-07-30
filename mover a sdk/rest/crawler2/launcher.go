package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	uccrawler "github.com/devpablocristo/crawler/internal/url-lister"
	crawler "github.com/devpablocristo/crawler/internal/url-lister/crawler"
)

type launcherPort interface {
	setup(ctx context.Context)
	stop(ctx context.Context) error
	init(ctx context.Context) error
}

type components struct {
	htmlParser crawler.HTMLParserPort
	service    crawler.CrawlerPort
	cliCfg     crawler.CliConfigPort
	cli        crawler.CliPort
	running    bool
}

type launcher struct {
	launcher components
}

func newLauncher() launcherPort {
	return &launcher{}
}

func (c *launcher) stop(ctx context.Context) error {
	return nil
}

func (c *launcher) init(ctx context.Context) error {
	if c.launcher.running {
		return errors.New("the web crawling service is already running")
	}

	c.setup(ctx)

	if err := c.launcher.cli.Execute(ctx); err != nil {
		return fmt.Errorf("error executing the CLI: %w", err)
	}

	c.launcher.running = true
	return nil
}

func (c *launcher) setup(ctx context.Context) {
	c.launcher.htmlParser = crawler.NewHTMLParser()
	c.launcher.service = uccrawler.NewCrawler(c.launcher.htmlParser)
	c.launcher.cliCfg = crawler.NewCliConfig(ctx, c.launcher.service)
	c.launcher.cli = crawler.NewCli(c.launcher.cliCfg)
}

func startCrawler(ctx context.Context) {
	crawlerSrv := newLauncher()
	if err := crawlerSrv.init(ctx); err != nil {
		log.Printf("error starting Crawler service: %v", err)
	}
}
