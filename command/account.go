package command

import (
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type AccountCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *AccountCommand) Help() string {
	helpText := `
Usage: godo-cli account [options]

accont command that are provided in the digitalocean.

Options:
  Todo
`
	return strings.TrimSpace(helpText)
}

func (c *AccountCommand) Run(args []string) int {
	account, _, err := c.Client.Account.Get()

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	c.Ui.Output(fmt.Sprintf("%s(%s) verified:%t limit:%d",
		account.Email, account.UUID, account.EmailVerified, account.DropletLimit))

	return 0
}

func (c *AccountCommand) Synopsis() string {
	return fmt.Sprintf("Show a account information.")
}
