package command

import (
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type ShutdownCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *ShutdownCommand) Help() string {
	helpText := `
	Usage: godo-cli shutdown [options] 
  
Options:
  -id=int The id of the droplet
  -name=string The name of the droplet

e.g.
  godo-cli shutdown -id=droplet id 
  godo-cli shutdown -name=droplet name
   
`
	return strings.TrimSpace(helpText)
}

func (c *ShutdownCommand) ShutdownById(id int) error {
	_, _, err := c.Client.DropletActions.Shutdown(id)

	if err != nil {
		return err
	}

	return nil
}

func (c *ShutdownCommand) Run(args []string) int {
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

	if name != "" {
		util := GodoUtil{Client: c.Client}
		droplet, err := util.GetDropletByName(name)

		if err != nil {
			c.Ui.Error(fmt.Sprintf("Failed to get droplet id for %s: %v", name, err))
			return -1
		}

		id = droplet.ID
	}
	err = c.ShutdownById(id)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("Queuing shutdown for %d ...done\n", id))

	return 0
}

func (c *ShutdownCommand) Synopsis() string {
	return fmt.Sprintf("Shutdown a droplet.")
}
