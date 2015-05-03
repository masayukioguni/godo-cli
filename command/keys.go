package command

import (
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type KeysCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *KeysCommand) Help() string {
	helpText := `
Usage: godo-cli keys [options]
Options:
  Todo
`
	return strings.TrimSpace(helpText)
}

func (c *KeysCommand) Run(args []string) int {
	opt := &godo.ListOptions{}
	keys, _, err := c.Client.Keys.List(opt)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	for _, key := range keys {
		c.Ui.Output(fmt.Sprintf("id:%d name:%s", key.ID, key.Name))
	}

	return 0
}

func (c *KeysCommand) Synopsis() string {
	return fmt.Sprintf("Show available SSH keys")
}
