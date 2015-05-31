package command

import (
	"bufio"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/masayukioguni/godo-cli/config"
	"github.com/mitchellh/cli"
	"os"
)

type AuthorizeCommand struct {
	Ui     cli.Ui
	Config *config.Config
	Client *godo.Client
}

func (c *AuthorizeCommand) Help() string {
	return ""
}

func (c *AuthorizeCommand) ask(text string, defaultText string) string {
	fmt.Printf("%s", text)
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	if input != "" {
		return input
	}
	return defaultText

}

func (c *AuthorizeCommand) Run(args []string) int {
	apikey := c.ask("Enter your API Token:", "Input api token")
	fmt.Println(`Defaults can be changed at any time in your ~/.godo-cli/config.yaml configuration file.`)

	c.Config.Authentication.APIKey = apikey

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

	fmt.Println("Authentication with DigitalOcean was successful!")
	return 0
}

func (c *AuthorizeCommand) Synopsis() string {
	return fmt.Sprintf("Authorize a DigitalOcean account with godo-cli")
}
