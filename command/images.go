package command

import (
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type ImagesCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *ImagesCommand) Help() string {
	helpText := `
  Usage: godo-cli images [options] 
  
  Images command that are provided in the digitalocean.

Options:
  -list=string List distribution images
  :ex 
    -list=distro List distribution images
    -list=app    List application images
    -list=user   List user images

`
	return strings.TrimSpace(helpText)
}

func (c *ImagesCommand) Run(args []string) int {

	var typeFlag string
	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)
	cmdFlags.StringVar(&typeFlag, "list", "", "specify the type of list. distro/app/user")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	var images []godo.Image
	var err error
	opt := &godo.ListOptions{}

	switch typeFlag {
	case "distro":
		images, _, err = c.Client.Images.ListDistribution(opt)
	case "app":
		images, _, err = c.Client.Images.ListApplication(opt)
	case "user":
		images, _, err = c.Client.Images.ListUser(opt)
	default:
		images, _, err = c.Client.Images.List(opt)
	}

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	for _, image := range images {
		c.Ui.Output(fmt.Sprintf("%s (id: %d, distro: %s) %v",
			image.Name, image.ID, image.Distribution, image.Slug))
	}

	return 0
}

func (c *ImagesCommand) Synopsis() string {
	return fmt.Sprintf("Retrieve a list of your images")
}
