package command

import (
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type DestroyCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *DestroyCommand) Help() string {
	helpText := `
	Usage: godo-cli destroy [options] 
  
Options:
  -id=int The id of the droplet
  -name=string The name of the droplet

e.g.
  godo-cli destroy -id=droplet id
  godo-cli destroy -name=droplet name
   
`
	return strings.TrimSpace(helpText)
}

func (c *DestroyCommand) DestroyById(id int) error {
	_, err := c.Client.Droplets.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (c *DestroyCommand) DestroyByName(name string) (int, error) {

	util := GodoUtil{Client: c.Client}
	droplet, err := util.GetDropletByName(name)

	if err != nil {
		return -1, err
	}
	return droplet.ID, c.DestroyById(droplet.ID)
}

func (c *DestroyCommand) Run(args []string) int {
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
		id, err = c.DestroyByName(name)
	} else {
		if id != 0 {
			err = c.DestroyById(id)
		}
	}

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("Queuing destroy for %d ...done\n", id))

	return 0
}

func (c *DestroyCommand) Synopsis() string {
	return fmt.Sprintf("Destroy a droplet.")
}
