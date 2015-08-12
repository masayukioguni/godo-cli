package command

import (
	"fmt"
	"github.com/digitalocean/godo"
)

type GodoUtil struct {
	Client *godo.Client
}

func (u *GodoUtil) GetDomains() ([]godo.Domain, error) {

	list := []godo.Domain{}
	opt := &godo.ListOptions{}

	for {
		domains, resp, err := u.Client.Domains.List(opt)

		if err != nil {
			return list, fmt.Errorf("Failed to request %v", err)
		}

		for _, i := range domains {
			list = append(list, i)
		}

		if resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return list, fmt.Errorf("Failed to CurrentPage %v", err)
		}

		opt.Page = page + 1
	}
	return list, nil
}

func (u *GodoUtil) GetDroplets() ([]godo.Droplet, error) {

	list := []godo.Droplet{}
	opt := &godo.ListOptions{}

	for {
		droplets, resp, err := u.Client.Droplets.List(opt)

		if err != nil {
			return list, fmt.Errorf("Failed to request %v", err)
		}

		for _, i := range droplets {
			list = append(list, i)
		}

		if resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return list, fmt.Errorf("Failed to CurrentPage %v", err)
		}

		opt.Page = page + 1
	}
	return list, nil
}

func (u *GodoUtil) GetDropletByName(targetName string) (*godo.Droplet, error) {
	droplets, err := u.GetDroplets()

	if err != nil {
		return nil, err
	}

	for _, droplet := range droplets {
		if droplet.Name == targetName {
			return &droplet, nil
		}
	}
	return nil, fmt.Errorf("Error %v is not found", targetName)
}

func (u *GodoUtil) GetDropletByID(targetID int) (*godo.Droplet, error) {
	droplets, err := u.GetDroplets()

	if err != nil {
		return nil, err
	}

	for _, droplet := range droplets {
		if droplet.ID == targetID {
			return &droplet, nil
		}
	}
	return nil, fmt.Errorf("Error %v is not found", targetID)
}

func (u *GodoUtil) GetDropletActions(targetID int) ([]godo.Action, error) {
	list := []godo.Action{}
	opt := &godo.ListOptions{}

	for {
		actions, resp, err := u.Client.Droplets.Actions(targetID, opt)

		if err != nil {
			return list, fmt.Errorf("Failed to request %v", err)
		}

		for _, i := range actions {
			list = append(list, i)
		}

		if resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return list, fmt.Errorf("Failed to CurrentPage %v", err)
		}

		opt.Page = page + 1
	}
	return list, nil
}
