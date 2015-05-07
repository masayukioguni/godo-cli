package command

import (
	"flag"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/masayukioguni/godo-cli/config"
	"github.com/mitchellh/cli"
	"os"
	"path/filepath"

	"strings"
)

type ConfigCommand struct {
	Ui     cli.Ui
	Config *config.Config
	Client *godo.Client
}

func (c *ConfigCommand) Help() string {
	helpText := `
Usage: godo-cli config [options]

Options:
  -id=int The id of the droplet (required)

  :ex 
  godo-cli info -id=droplet id
`
	return strings.TrimSpace(helpText)
}

func (c *ConfigCommand) Set(args []string) int {
	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

	var region string
	var image string
	var size string
	var key string

	cmdFlags.StringVar(&region, "region", "", "")
	cmdFlags.StringVar(&image, "image", "", "")
	cmdFlags.StringVar(&size, "size", "", "")
	cmdFlags.StringVar(&key, "key", "", "")

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

	home := os.Getenv("HOME")
	if home == "" {
		fmt.Errorf("Error Getenv $HOME not found")
		return 1
	}

	saveDirectory := filepath.Join(home, config.GetDefaultDirectory())
	_, err = os.Stat(saveDirectory)
	if err != nil {
		if err = os.Mkdir(saveDirectory, 0755); err != nil {
			fmt.Errorf("Error mkdir %s", saveDirectory)
			return 0
		}
	}
	savePath := filepath.Join(saveDirectory, config.GetDefaultConfigName())

	err = config.SaveConfig(savePath, c.Config)

	if err != nil {
		fmt.Errorf("Error SaveConfig %s", err)
		return 1
	}

	fmt.Println("Authentication with DigitalOcean was successful!")

	return 0

}

func (c *ConfigCommand) Get(args []string) int {
	c.Ui.Output(fmt.Sprintf("image: %v", c.Config.Defaults.Image))
	c.Ui.Output(fmt.Sprintf("Size: %v", c.Config.Defaults.Size))
	c.Ui.Output(fmt.Sprintf("Region: %v", c.Config.Defaults.Region))
	c.Ui.Output(fmt.Sprintf("Keys: %v", c.Config.Defaults.Key))
	return 0
}

func (c *ConfigCommand) Run(args []string) int {

	if len(args) < 1 {
		c.Ui.Output(c.Help())
		return -1
	}

	subcommand := args[0]
	newArgs := args[1:]
	var res int

	switch subcommand {
	case "set":
		res = c.Set(newArgs)
	case "get":
		res = c.Get(newArgs)
	default:

	}

	return res
}

func (c *ConfigCommand) Synopsis() string {
	return fmt.Sprintf("Show a droplet's information.")
}
