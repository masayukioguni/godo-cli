package command

import (
	"flag"
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
  -q quiet mode (only displays dropet ids)
  .
`
	return strings.TrimSpace(helpText)
}

type DropletsFlags struct {
	Quiet bool
}

func (c *DropletsCommand) parse(args []string) (*DropletsFlags, error) {
	flags := new(DropletsFlags)

	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	cmdFlags.BoolVar(&flags.Quiet, "q", false, "")

	err := cmdFlags.Parse(args)

	if err != nil {
		return nil, err
	}

	return flags, nil
}

func (c *DropletsCommand) Run(args []string) int {
	flags, err := c.parse(args)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to parse %v", err))
		return -1
	}

	util := GodoUtil{Client: c.Client}
	droplets, err := util.GetDroplets()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	for _, droplet := range droplets {
		if flags.Quiet {
			fmt.Printf("%d\n", droplet.ID)
		} else {
			fmt.Printf("%s (ip: %s, status: %s, region :%s, id: %d)\n",
				droplet.Name, GetNetworksV4IPAddress(droplet.Networks), droplet.Status, droplet.Region.Slug, droplet.ID)
		}
	}
	return 0
}

func (c *DropletsCommand) Synopsis() string {
	return fmt.Sprintf("Retrieve a list of your droplets.")
}
