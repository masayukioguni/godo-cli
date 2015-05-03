package command

import (
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type DropletsCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *DropletsCommand) Help() string {
	helpText := `
Usage: godo-cli droplets [options]
Options:
  Todo
`
	return strings.TrimSpace(helpText)
}

func (c *DropletsCommand) Run(args []string) int {
	opt := &godo.ListOptions{}
	droplets, _, err := c.Client.Droplets.List(opt)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	for _, droplet := range droplets {
		fmt.Printf("%s (status: %s, region :%s, id: %d)\n",
			droplet.Name, droplet.Status, droplet.Region.Slug, droplet.ID)
	}

	return 0
}

func (c *DropletsCommand) Synopsis() string {
	return fmt.Sprintf("Show available droplet sizes")
}
