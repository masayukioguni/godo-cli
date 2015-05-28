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
  -id=int The id of the droplet
  -id=string The name of the droplet


e.g.
  droplet by id
    godo-cli info -id=droplet id
  droplet by name
    godo-cli info -name=name

`
	return strings.TrimSpace(helpText)
}

func (c *InfoCommand) DropletById(id int) (*godo.Droplet, error) {
	dropletRoot, _, err := c.Client.Droplets.Get(id)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return nil, err
	}

	return dropletRoot.Droplet, nil

}

func (c *InfoCommand) DropletByName(name string) (*godo.Droplet, error) {

	util := GodoUtil{Client: c.Client}
	droplet, err := util.GetDropletByName(name)

	if err != nil {
		return nil, err
	}

	return droplet, nil
}

func (c *InfoCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	var id = 0
	var name = ""
	cmdFlags.IntVar(&id, "id", 0, "")
	cmdFlags.StringVar(&name, "name", "", "")

	err := cmdFlags.Parse(args)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to parse %v", err))
		return -1
	}

	if name == "" && id == 0 {
		c.Help()
		return -1
	}

	var droplet *godo.Droplet
	if name != "" {
		droplet, err = c.DropletByName(name)
	} else {

		if id != 0 {
			droplet, err = c.DropletById(id)
		}
	}

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("%s (status: %s, region :%s, id: %d, image id:%d size:%s)\n",
		droplet.Name, droplet.Status, droplet.Region.Slug, droplet.ID, droplet.Image.ID, droplet.Size.Slug))

	return 0
}

func (c *InfoCommand) Synopsis() string {
	return fmt.Sprintf("Show a droplet's information.")
}
