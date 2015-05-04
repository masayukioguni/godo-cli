package command

import (
	"errors"
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type CreateCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *CreateCommand) Help() string {
	helpText := `
  Usage: godo-cli create [options] 

Options:
  -name=string The name of the droplet (required)
  -size=string The size of the droplet (ex: 512mb)
  -region=string The region of the droplet (ex: nyc1)
  -image=string The image of the droplet (ex: 9801950)
  -keys=int  The ssh key id of the droplet (ex: ssh key id)

  :ex 
  godo-cli create -name=godo-test -size=512mb -region=nyc1 -image=9801950 -keys=xxxxx
`
	return strings.TrimSpace(helpText)
}

type CreateFlags struct {
	Name   string
	Image  string
	Size   string
	Region string
	Keys   int
}

func (c *CreateCommand) parse(args []string) (*CreateFlags, error) {
	flags := new(CreateFlags)

	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	cmdFlags.StringVar(&flags.Name, "name", "", "")
	cmdFlags.StringVar(&flags.Size, "size", "", "")
	cmdFlags.StringVar(&flags.Image, "image", "", "")
	cmdFlags.StringVar(&flags.Region, "region", "", "")
	cmdFlags.IntVar(&flags.Keys, "keys", 0, "")

	err := cmdFlags.Parse(args)

	if err != nil {
		return nil, err
	}

	if flags.Name == "" {
		return nil, errors.New("invalid name")
	}

	if flags.Size == "" {
		flags.Size = "512mb"
	}

	if flags.Image == "" {
		flags.Image = "9801950"
	}

	if flags.Region == "" {
		flags.Region = "nyc1"
	}

	if flags.Keys == 0 {
		flags.Keys = 0
	}

	return flags, nil
}

func (c *CreateCommand) Run(args []string) int {

	flags, err := c.parse(args)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to parse %v", err))
		return 1
	}

	createRequest := &godo.DropletCreateRequest{
		Name:   flags.Name,
		Region: flags.Region,
		Size:   flags.Size,
		SSHKeys: []godo.DropletCreateSSHKey{
			{ID: flags.Keys},
		},
		Image: godo.DropletCreateImage{
			Slug: flags.Image,
		},
	}

	newDroplet, _, err := c.Client.Droplets.Create(createRequest)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("%v ", newDroplet))

	return 0
}

func (c *CreateCommand) Synopsis() string {
	return fmt.Sprintf("Create a droplet.")
}
