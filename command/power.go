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
  -mode= off[default], on, cycle 

e.g.
  godo-cli power -id=droplet id -mode=cycle # power cycles droplet by id
  godo-cli power -name=droplet name # powers off droplet by name
   
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

func (c *PowerCommand) PowerOnById(id int) error {
	_, _, err := c.Client.DropletActions.PowerOn(id)

	if err != nil {
		return err
	}

	return nil
}

func (c *PowerCommand) PowerCycleById(id int) error {
	_, _, err := c.Client.DropletActions.PowerCycle(id)

	if err != nil {
		return err
	}

	return nil
}

func (c *PowerCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	var id = 0
	var name = ""
	var mode = ""

	cmdFlags.IntVar(&id, "id", 0, "")
	cmdFlags.StringVar(&name, "name", "", "")
	cmdFlags.StringVar(&mode, "power", "", "")

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
		id = droplet.ID
		if err != nil {
			c.Ui.Error(fmt.Sprintf("Failed to get ID for droplet %s: %v", name, err))
			return -1
		}

	}

	switch mode {
	case "on":
		err = c.PowerOnById(id)
	case "off":
		err = c.PowerOffById(id)
	case "cycle":
		err = c.PowerCycleById(id)
	default:
		c.Ui.Error(fmt.Sprintf("Unknown mode: %s", mode))
		return -1
	}

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("Queuing power %s for %d ...done\n", mode, id))

	return 0
}

func (c *PowerCommand) Synopsis() string {
	return fmt.Sprintf("Power off/on/cycle a droplet.")
}
