package command

import (
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type RegionsCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *RegionsCommand) Help() string {
	helpText := `
Usage: godo-cli regisons [options]
Options:
  Todo
`
	return strings.TrimSpace(helpText)
}

func (c *RegionsCommand) Run(args []string) int {
	opt := &godo.ListOptions{}
	regions, _, err := c.Client.Regions.List(opt)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	for _, region := range regions {
		c.Ui.Output(fmt.Sprintf("slug: %s name: %s", region.Slug, region.Name))
	}

	return 0
}

func (c *RegionsCommand) Synopsis() string {
	return fmt.Sprintf("Show available droplet regions")
}
