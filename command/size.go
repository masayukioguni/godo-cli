package command

import (
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/mitchellh/cli"
	"strings"
)

type SizeCommand struct {
	Ui  cli.Ui
	Cli *godo.Client
}

func (c *SizeCommand) Help() string {
	helpText := `
Usage: godo-cli size [options]
Options:
  Todo
`
	return strings.TrimSpace(helpText)
}

func (c *SizeCommand) Run(args []string) int {
	opt := &godo.ListOptions{}
	sizes, _, err := c.Cli.Sizes.List(opt)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
		return -1
	}

	for _, size := range sizes {
		c.Ui.Output(fmt.Sprintf("slug:%5s memory:%6dmb vcpus:%2d disk:%3dgb",
			size.Slug, size.Memory, size.Vcpus, size.Disk))
	}

	return 0
}

func (c *SizeCommand) Synopsis() string {
	return fmt.Sprintf("Show available droplet sizes")
}
