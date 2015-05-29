package command

import (
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type PowerCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *PowerCommand) Help() string {
	helpText := `
	Usage: godo-cli power [options] 
  
Options:
  -id=int The id of the droplet
  -name=string The name of the droplet
  -off [default, exclusive cannot be used inconjunction with -on]
  -on [exclsuive cannot be used in conjunction with -off]

e.g.
  godo-cli power -id=droplet id -off
  godo-cli power -name=droplet name -on
   
`
	return strings.TrimSpace(helpText)
}

func (c *PowerCommand) PowerOffById(id int) error {
	_, _, err := c.Client.DropletActions.PowerOff(id)

	if err != nil {
		return err
	}

	return nil
}

func (c *PowerCommand) PowerOffByName(name string) (int, error) {

	util := GodoUtil{Client: c.Client}
	droplet, err := util.GetDropletByName(name)

	if err != nil {
		return -1, err
	}
	return droplet.ID, c.PowerOffById(droplet.ID)
}

func (c *PowerCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	var id = 0
	var name = ""
	cmdFlags.IntVar(&id, "id", 0, "")
	cmdFlags.StringVar(&name, "name", "", "")
	// cmdFlags.BoolVar(&on, "on", true, "")

	err := cmdFlags.Parse(args)

	c.Ui.Output(fmt.Sprintf("Testing before func call\n"))
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to parse %v", err))
		return -1
	}

	if name == "" && id == 0 {
		c.Help()
		return -1
	}

	if name != "" {
		id, err = c.PowerOffByName(name)
		c.Ui.Output(fmt.Sprintf("PowerOffByName\n"))
	} else {
		if id != 0 {
			err = c.PowerOffById(id)
			c.Ui.Output(fmt.Sprintf("PowerOffById\n"))
		}
	}

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("Queuing power off for %d ...done\n", id))

	return 0
}

func (c *PowerCommand) Synopsis() string {
	return fmt.Sprintf("Power off a droplet.")
}
