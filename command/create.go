package command

import (
	"errors"
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/masayukioguni/godo-cli/config"
	"github.com/mitchellh/cli"
	"strconv"
	"strings"
)

type CreateCommand struct {
	Ui     cli.Ui
	Client *godo.Client
	Config *config.Config
}

func (c *CreateCommand) Help() string {
	helpText := `
  Usage: godo-cli create [options] 

Options:
  -name=string The name of the droplet (required)
  -size=string The size of the droplet (ex: 512mb)
  -region=string The region of the droplet (ex: nyc1)
  -image=int The image id of the droplet (ex: 9801950)
  -key=int  The ssh key id of the droplet (ex: ssh key id)

e.g.
  godo-cli create -name=godo-test -size=512mb -region=nyc1 -image=9801950 -key=xxxxx
`
	return strings.TrimSpace(helpText)
}

type CreateFlags struct {
	Name              string
	Image             string
	Size              string
	Region            string
	Key               string
	Backups           bool
	IPv6              bool
	PrivateNetworking bool
}

func (c *CreateCommand) parse(args []string) (*CreateFlags, error) {
	flags := new(CreateFlags)

	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	cmdFlags.StringVar(&flags.Name, "name", "", "")
	cmdFlags.StringVar(&flags.Size, "size", c.Config.Defaults.Size, "")
	cmdFlags.StringVar(&flags.Image, "image", c.Config.Defaults.Image, "")
	cmdFlags.StringVar(&flags.Region, "region", c.Config.Defaults.Region, "")
	cmdFlags.StringVar(&flags.Key, "key", c.Config.Defaults.Key, "")
	cmdFlags.BoolVar(&flags.Backups, "backups", c.Config.Defaults.Backups, "")
	cmdFlags.BoolVar(&flags.Backups, "ipv6", c.Config.Defaults.IPv6, "")
	cmdFlags.BoolVar(&flags.PrivateNetworking, "private_networking", c.Config.Defaults.PrivateNetworking, "")

	err := cmdFlags.Parse(args)

	if err != nil {
		return nil, err
	}

	if flags.Name == "" {
		return nil, errors.New("invalid name")
	}

	return flags, nil
}
func (c *CreateCommand) GetDropletCreateImage(text string) godo.DropletCreateImage {

	if idInt, err := strconv.Atoi(text); err == nil {
		return godo.DropletCreateImage{
			ID: idInt,
		}
	}

	return godo.DropletCreateImage{
		Slug: text,
	}

}

func (c *CreateCommand) Run(args []string) int {

	flags, err := c.parse(args)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to parse %v", err))
		return -1
	}

	key, err := strconv.Atoi(flags.Key)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to strconv.Atoi %v", err))
		return -1
	}

	image := c.GetDropletCreateImage(flags.Image)

	createRequest := &godo.DropletCreateRequest{
		Name:   flags.Name,
		Region: flags.Region,
		Size:   flags.Size,
		SSHKeys: []godo.DropletCreateSSHKey{
			{ID: key},
		},
		Image:             image,
		Backups:           flags.Backups,
		IPv6:              flags.IPv6,
		PrivateNetworking: flags.PrivateNetworking,
	}

	newDroplet, _, err := c.Client.Droplets.Create(createRequest)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("Queueing creation of droplet '%s' ...done\n", newDroplet.Droplet.Name))

	return 0
}

func (c *CreateCommand) Synopsis() string {
	return fmt.Sprintf("Create a droplet.")
}
