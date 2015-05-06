package command

import (
	"bufio"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/masayukioguni/godo-cli/config"
	"github.com/mitchellh/cli"

	"os"
	"path/filepath"
)

type AuthorizeCommand struct {
	Ui     cli.Ui
	Client *godo.Client
}

func (c *AuthorizeCommand) Help() string {
	return ""
}

func (c *AuthorizeCommand) ask(text string, defaultText string) string {
	fmt.Printf(text)
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	if input != "" {
		return input
	}
	return defaultText

}

func (c *AuthorizeCommand) Run(args []string) int {
	apikey := c.ask("Entser your API Token:", "Input api ooken")
	fmt.Println(`Defaults can be changed at any time in your ~/.godo-cli/config.yaml configuration file.`)

	env := &config.Config{}
	env.Authentication.APIKey = apikey
	home := os.Getenv("HOME")
	if home == "" {
		fmt.Errorf("Error Getenv $HOME not found")
		return 1
	}

	saveDirectory := filepath.Join(home, config.GetDefaultDirectory())
	_, err := os.Stat(saveDirectory)
	if err != nil {
		if err = os.Mkdir(saveDirectory, 0755); err != nil {
			fmt.Errorf("Error mkdir %s", saveDirectory)
			return 0
		}
	}
	savePath := filepath.Join(saveDirectory, config.GetDefaultConfigName())

	err = config.SaveConfig(savePath, env)

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
