package command

import (
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"os"
	"os/exec"
	"strings"
)

type SSHCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *SSHCommand) Help() string {
	helpText := `
Usage: godo-cli account [options]
Options:
  Todo
`
	return strings.TrimSpace(helpText)
}

func (c *SSHCommand) GetDropletByID(id int) (*godo.Droplet, error) {
	util := GodoUtil{Client: c.Client}
	droplet, err := util.GetDropletByID(id)

	if err != nil {
		return nil, err
	}
	return droplet, nil

}

func (c *SSHCommand) GetDropletByName(name string) (*godo.Droplet, error) {
	util := GodoUtil{Client: c.Client}
	droplet, err := util.GetDropletByName(name)

	if err != nil {
		return nil, err
	}
	return droplet, nil
}

func (c *SSHCommand) Run(args []string) int {

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
		droplet, err = c.GetDropletByName(name)
	} else {
		if id != 0 {
			droplet, err = c.GetDropletByID(id)
		}
	}

	fmt.Printf("%v\n", droplet)

	s := fmt.Sprintf("root@%s", droplet.Networks.V4[0].IPAddress)

	cmd := exec.Command("ssh", s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	return 0
}

func (c *SSHCommand) Synopsis() string {
	return fmt.Sprintf("Show a account information.")
}
