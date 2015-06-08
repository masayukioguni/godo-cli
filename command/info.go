package command

import (
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
	"time"
)

type InfoCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *InfoCommand) Help() string {
	helpText := `
Usage: godo-cli info [options]

Options:
  -id=int The id of the droplet
  -id=string The name of the droplet


e.g.
  droplet by id
    godo-cli info -id=droplet id
  droplet by name
    godo-cli info -name=name

`
	return strings.TrimSpace(helpText)
}

func (c *InfoCommand) DropletById(id int) (*godo.Droplet, error) {
	droplet, _, err := c.Client.Droplets.Get(id)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return nil, err
	}

	return droplet, nil

}

func (c *InfoCommand) DropletByName(name string) (*godo.Droplet, error) {

	util := GodoUtil{Client: c.Client}
	droplet, err := util.GetDropletByName(name)

	if err != nil {
		return nil, err
	}

	return droplet, nil
}

func (c *InfoCommand) Run(args []string) int {
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
		droplet, err = c.DropletByName(name)
	} else {

		if id != 0 {
			droplet, err = c.DropletById(id)
		}
	}

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("%s (status: %s, region :%s, id: %d, image id:%d size:%s)\n",
		droplet.Name, droplet.Status, droplet.Region.Slug, droplet.ID, droplet.Image.ID, droplet.Size.Slug))

	c.Ui.Output(fmt.Sprintf("Droplet history\n"))
	c.Ui.Output(fmt.Sprintf("===============\n"))

	util := GodoUtil{Client: c.Client}
	actions, err := util.GetDropletActions(droplet.ID) // func (u *GodoUtil) GetDropletActions(targetID int) ([]godo.Action, error) {¬
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	/*
	   godo.Action{ID:52715195, Status:"completed", Type:"power_on", StartedAt:godo.Timestamp{2015-06-07 14:08:03 +0000 UTC}, CompletedAt:godo.Timestamp{2015-06-07 14:08:08 +0000 UTC}, ResourceID:5610866, ResourceType:"droplet", Region:godo.Region{Slug:"lon1", Name:"London 1", Sizes:["512mb" "1gb" "2gb" "4gb" "8gb" "32gb" "48gb" "64gb" "16gb"], Available:true, Features:["private_networking" "backups" "ipv6" "metadata"]}, RegionSlug:"lon1"}
	    godo.Action{ID:52715167, Status:"completed", Type:"shutdown", StartedAt:godo.Timestamp{2015-06-07 14:07:37 +0000 UTC}, CompletedAt:godo.Timestamp{2015-06-07 14:07:40 +0000 UTC}, ResourceID:5610866, ResourceType:"droplet", Region:godo.Region{Slug:"lon1", Name:"London 1", Sizes:["512mb" "1gb" "2gb" "4gb" "8gb" "32gb" "48gb" "64gb" "16gb"], Available:true, Features:["private_networking" "backups" "ipv6" "metadata"]}, RegionSlug:"lon1"}
	    godo.Action{ID:52715079, Status:"completed", Type:"power_on", StartedAt:godo.Timestamp{2015-06-07 14:05:52 +0000 UTC}, CompletedAt:godo.Timestamp{2015-06-07 14:06:00 +0000 UTC}, ResourceID:5610866, ResourceType:"droplet", Region:godo.Region{Slug:"lon1", Name:"London 1", Sizes:["512mb" "1gb" "2gb" "4gb" "8gb" "32gb" "48gb" "64gb" "16gb"], Available:true, Features:["private_networking" "backups" "ipv6" "metadata"]}, RegionSlug:"lon1"}
	    godo.Action{ID:52715050, Status:"completed", Type:"power_off", StartedAt:godo.Timestamp{2015-06-07 14:05:23 +0000 UTC}, CompletedAt:godo.Timestamp{2015-06-07 14:05:42 +0000 UTC}, ResourceID:5610866, ResourceType:"droplet", Region:godo.Region{Slug:"lon1", Name:"London 1", Sizes:["512mb" "1gb" "2gb" "4gb" "8gb" "32gb" "48gb" "64gb" "16gb"], Available:true, Features:["private_networking" "backups" "ipv6" "metadata"]}, RegionSlug:"lon1"}
	    godo.Action{ID:52714162, Status:"completed", Type:"create", StartedAt:godo.Timestamp{2015-06-07 13:49:47 +0000 UTC}, CompletedAt:godo.Timestamp{2015-06-07 14:01:50 +0000 UTC}, ResourceID:5610866, ResourceType:"droplet", Region:godo.Region{Slug:"lon1", Name:"London 1", Sizes:["512mb" "1gb" "2gb" "4gb" "8gb" "32gb" "48gb" "64gb" "16gb"], Available:true, Features:["private_networking" "backups" "ipv6" "metadata"]}, RegionSlug:"lon1"}
	*/
	c.Ui.Output(fmt.Sprintf("Event | Initiated             | Execution Time\n"))
	for _, a := range actions {
		// c.Ui.Output(fmt.Sprintf("%+v", a))
		start := a.StartedAt.String()
		finish := a.CompletedAt.String()
		// parse, err := time.Parse(a.StartedAt, a.CompletedAt)
		parse, err := time.Parse(start, finish)

		if err != nil {
			c.Ui.Error(fmt.Sprintf("Failed to parse time %v", err))
			return -1
		}

		delta := time.Since(parse)

		c.Ui.Output(fmt.Sprintf("%s    | %s                    | %s\n", a.Type, a.StartedAt, delta))

	}

	return 0
}

func (c *InfoCommand) Synopsis() string {
	return fmt.Sprintf("Show a droplet's information.")
}
