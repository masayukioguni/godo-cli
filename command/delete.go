package command

import (
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type DeleteCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *DeleteCommand) Help() string {
	helpText := `
	Usage: godo-cli delete [options] 
  
Options:
  -id=int The id of the droplet (required)

e.g.
  godo-cli delete -id=droplet id
`
	return strings.TrimSpace(helpText)
}

func (c *DeleteCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	var id = 0
	cmdFlags.IntVar(&id, "id", 0, "")

	err := cmdFlags.Parse(args)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to parse %v", err))
		return -1
	}

	_, err = c.Client.Droplets.Delete(id)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("Queuing destroy for %d ...done\n", id))

	return 0
}

func (c *DeleteCommand) Synopsis() string {
	return fmt.Sprintf("Destroy a droplet.")
}
