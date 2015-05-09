package command

import (
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type InfoCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *InfoCommand) Help() string {
	helpText := `
Usage: godo-cli info [options]

Options:
  -id=int The id of the droplet (required)

e.g.
  godo-cli info -id=droplet id
`
	return strings.TrimSpace(helpText)
}

func (c *InfoCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	var id = 0
	cmdFlags.IntVar(&id, "id", 0, "")

	err := cmdFlags.Parse(args)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to parse %v", err))
		return -1
	}

	dropletRoot, _, err := c.Client.Droplets.Get(id)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	droplet := dropletRoot.Droplet
	c.Ui.Output(fmt.Sprintf("%s (status: %s, region :%s, id: %d, image id:%d size:%s)\n",
		droplet.Name, droplet.Status, droplet.Region.Slug, droplet.ID, droplet.Image.ID, droplet.Size.Slug))

	return 0
}

func (c *InfoCommand) Synopsis() string {
	return fmt.Sprintf("Show a droplet's information.")
}
