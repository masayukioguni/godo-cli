package command

import (
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/masayukioguni/godo-cli/config"
	"github.com/mitchellh/cli"
	"strconv"
	"strings"
)

type ConfigCommand struct {
	Ui     cli.Ui
	Config *config.Config
	Client *godo.Client
}

func (c *ConfigCommand) Help() string {
	helpText := `
Usage: godo-cli config subcommand [options]

Subcommand:
  get
  set 

Options:
  -region=string region slug ex:nyc1
  -image=int image id
  -size=string size slug ex:512mb
  -key=int SSHKey id 
  -backups=bool 
  -ipv6=bool 
  -private_networking=bool 


e.g.
  configuration
    godo-cli config get 

  set region to config
    godo-cli config set -region=nyc3


`
	return strings.TrimSpace(helpText)
}

func (c *ConfigCommand) Set(args []string) int {
	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	var region string
	var image string
	var size string
	var key string
	var backups string
	var ipv6 string
	var private_networking string

	cmdFlags.StringVar(&region, "region", "", "")
	cmdFlags.StringVar(&image, "image", "", "")
	cmdFlags.StringVar(&size, "size", "", "")
	cmdFlags.StringVar(&key, "key", "", "")
	cmdFlags.StringVar(&backups, "backups", "", "")
	cmdFlags.StringVar(&ipv6, "ipv6", "", "")
	cmdFlags.StringVar(&private_networking, "private_networking", "", "")

	err := cmdFlags.Parse(args)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to parse %v", err))
		return -1
	}

	if region != "" {
		c.Config.Defaults.Region = region
	}

	if image != "" {
		c.Config.Defaults.Image = image
	}

	if size != "" {
		c.Config.Defaults.Size = size
	}

	if key != "" {
		c.Config.Defaults.Key = key
	}

	if backups != "" {
		c.Config.Defaults.Backups, _ = strconv.ParseBool(backups)
	}

	if ipv6 != "" {
		c.Config.Defaults.IPv6, _ = strconv.ParseBool(ipv6)
	}

	if private_networking != "" {
		c.Config.Defaults.PrivateNetworking, _ = strconv.ParseBool(private_networking)
	}

	savePath, err := config.GetConfigPath()

	if err != nil {
		fmt.Errorf("Error GetConfigPath %s", err)
		return 1
	}

	err = config.SaveConfig(savePath, c.Config)

	if err != nil {
		fmt.Errorf("Error SaveConfig %s", err)
		return 1
	}

	fmt.Println("successful!")

	return 0

}

func (c *ConfigCommand) Get(args []string) int {
	c.Ui.Output(fmt.Sprintf("Defaults"))
	c.Ui.Output(fmt.Sprintf("image: %v", c.Config.Defaults.Image))
	c.Ui.Output(fmt.Sprintf("Size: %v", c.Config.Defaults.Size))
	c.Ui.Output(fmt.Sprintf("Region: %v", c.Config.Defaults.Region))
	c.Ui.Output(fmt.Sprintf("Key: %v", c.Config.Defaults.Key))
	c.Ui.Output(fmt.Sprintf("Backups: %v", c.Config.Defaults.Backups))
	c.Ui.Output(fmt.Sprintf("IPv6: %v", c.Config.Defaults.IPv6))
	c.Ui.Output(fmt.Sprintf("PrivateNetworking: %v", c.Config.Defaults.PrivateNetworking))

	return 0
}

func (c *ConfigCommand) Run(args []string) int {

	if len(args) < 1 {
		c.Ui.Output(c.Help())
		return 0
	}

	newArgs := args[1:]
	switch args[0] {
	case "set":
		return c.Set(newArgs)
	case "get":
		return c.Get(newArgs)
	}

	return -1
}

func (c *ConfigCommand) Synopsis() string {
	return fmt.Sprintf("configuration of the default setting.")
}
