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
  Usage: godo-cli images [subcommand] [options] 

    Images command that are provided in the digitalocean.

Subcommand:
  list
  info
  update
  delete

List Options:
  -type=string List type

Delete Options:
  (required) -id=int image id

Update Options:
  (required) -id=int image id
  (required) -name=string new name

:ex
  *List distribution images
    godo-cli images list -type=distro
  *List application images
    godo-cli images list -type=app
  *List user images
    godo-cli images list -type=user
  *Infomation an image
    godo-cli images info -id=image_id
  *update an image name
    godo-cli images update -id=image_id -name=new-name
  *Delete an image
    godo-cli images delete -id=image_id

`
	return strings.TrimSpace(helpText)
}

func (c *ImagesCommand) getList(listFunc func(*godo.ListOptions) ([]godo.Image, *godo.Response, error)) ([]godo.Image, *godo.Response, error) {
	opt := &godo.ListOptions{}
	return listFunc(opt)
}

func (c *ImagesCommand) List(args []string) int {

	var typeFlag string
	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)
	cmdFlags.StringVar(&typeFlag, "type", "", "specify the type of list. distro/app/user")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	var images []godo.Image
	var err error
	opt := &godo.ListOptions{}

	switch typeFlag {
	case "distro":
		images, _, err = c.getList(c.Client.Images.ListDistribution)
	case "app":
		images, _, err = c.getList(c.Client.Images.ListApplication)
	case "user":
		images, _, err = c.getList(c.Client.Images.ListUser)
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

func (c *ImagesCommand) Delete(args []string) int {

	var imageID int

	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)
	cmdFlags.IntVar(&imageID, "id", -1, "")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if imageID == -1 {
		c.Ui.Error(fmt.Sprintf("required image id."))
		return -1
	}

	_, err := c.Client.Images.Delete(imageID)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	return 0
}

func (c *ImagesCommand) Update(args []string) int {

	var imageID int
	var name string

	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)
	cmdFlags.IntVar(&imageID, "id", -1, "")

	cmdFlags.StringVar(&name, "name", "", "")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if name == "" {
		c.Ui.Error(fmt.Sprintf("required name."))
		return -1
	}

	if imageID == -1 {
		c.Ui.Error(fmt.Sprintf("required image id."))
		return -1
	}

	updateRequest := godo.ImageUpdateRequest{
		Name: name,
	}

	image, _, err := c.Client.Images.Update(imageID, &updateRequest)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("%s (id: %d, distro: %s) %v",
		image.Name, image.ID, image.Distribution, image.Slug))

	return 0
}

func (c *ImagesCommand) Info(args []string) int {

	var imageID int

	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)
	cmdFlags.IntVar(&imageID, "id", -1, "")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if imageID == -1 {
		c.Ui.Error(fmt.Sprintf("required image id."))
		return -1
	}

	image, _, err := c.Client.Images.GetByID(imageID)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("%s (id: %d, distro: %s) %v",
		image.Name, image.ID, image.Distribution, image.Slug))

	return 0
}
func (c *ImagesCommand) Run(args []string) int {

	if len(args) < 1 {
		c.Ui.Output(c.Help())
		return -1
	}

	subcommand := args[0]
	newArgs := args[1:]
	var res int

	switch subcommand {
	case "list":
		res = c.List(newArgs)
	case "info":
		res = c.Info(newArgs)
	case "update":
		res = c.Update(newArgs)
	case "delete":
		res = c.Delete(newArgs)
	default:

	}

	return res
}

func (c *ImagesCommand) Synopsis() string {
	return fmt.Sprintf("Retrieve a list of your images")
}
