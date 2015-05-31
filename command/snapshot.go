package command

import (
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type SnapshotCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *SnapshotCommand) Help() string {
	helpText := `
	Usage: godo-cli snapshot [options] 
  
Options:
  -id=int The id of the droplet
  -name=string The name of the droplet
  -snapshot=Name of snapshot [mandatory]

e.g.
  godo-cli snapshot -id=droplet id -snapshot=foo
  godo-cli snapshot -name=droplet name  -snapshot=foo
   
`
	return strings.TrimSpace(helpText)
}

func (c *SnapshotCommand) SnapshotById(id int, snapshot string) error {
	_, _, err := c.Client.DropletActions.Snapshot(id, snapshot)

	if err != nil {
		return err
	}

	return nil
}

func (c *SnapshotCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	var id = 0
	var name = ""
	var snapshot = ""
	cmdFlags.IntVar(&id, "id", 0, "")
	cmdFlags.StringVar(&name, "name", "", "")
	cmdFlags.StringVar(&snapshot, "snapshot", "", "")

	err := cmdFlags.Parse(args)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to parse %v", err))
		return -1
	}

	if name == "" && id == 0 && snapshot == "" {
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
	err = c.SnapshotById(id, snapshot)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("Queuing snapshot for %d ...done\n", id))

	return 0
}

func (c *SnapshotCommand) Synopsis() string {
	return fmt.Sprintf("Create a snapshot of a droplet.")
}
