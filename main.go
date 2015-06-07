package main

import (
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/masayukioguni/godo-cli/command"
	"github.com/masayukioguni/godo-cli/config"
	"github.com/mitchellh/cli"
	"os"
)

var GitCommit string

const ApplicationName = "godo-cli"
const Version = "0.0.5"

func getClinet(accessToken string) *godo.Client {

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: accessToken},
	}
	return godo.NewClient(t.Client())

}

func main() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	configPath, err := config.GetConfigPath()
	if err != nil {
		fmt.Errorf("Error GetConfigPath %s", err)
		os.Exit(1)
	}

	_, err = os.Stat(configPath)
	if err != nil {
		configDummy := &config.Config{}
		configDirPath, _ := config.GetConfigDirectory()

		if err := os.Mkdir(configDirPath, 0766); err != nil {
			fmt.Errorf("Error LoadConfig %s", err)
			os.Exit(1)
		}

		config.SaveConfig(configPath, configDummy)
	}

	config, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Errorf("Error LoadConfig %s", err)
		return
	}

	godoCli := getClinet(config.Authentication.APIKey)

	c := cli.NewCLI(ApplicationName, Version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"authorize": func() (cli.Command, error) {
			return &command.AuthorizeCommand{
				Ui:     ui,
				Config: config,
				Client: godoCli,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Ui:      ui,
				AppName: ApplicationName,
				Version: Version,
			}, nil
		},
		"sizes": func() (cli.Command, error) {
			return &command.SizesCommand{
				Ui:     ui,
				Client: godoCli,
			}, nil
		},
		"regions": func() (cli.Command, error) {
			return &command.RegionsCommand{
				Ui:     ui,
				Client: godoCli,
			}, nil
		},
		"images": func() (cli.Command, error) {
			return &command.ImagesCommand{
				Ui:     ui,
				Client: godoCli,
			}, nil
		},
		"droplets": func() (cli.Command, error) {
			return &command.DropletsCommand{
				Ui:     ui,
				Client: godoCli,
			}, nil
		},
		"keys": func() (cli.Command, error) {
			return &command.KeysCommand{
				Ui:     ui,
				Client: godoCli,
			}, nil
		},
		"account": func() (cli.Command, error) {
			return &command.AccountCommand{
				Ui:     ui,
				Client: godoCli,
			}, nil
		},
		"create": func() (cli.Command, error) {
			return &command.CreateCommand{
				Ui:     ui,
				Client: godoCli,
				Config: config,
			}, nil
		},
		"info": func() (cli.Command, error) {
			return &command.InfoCommand{
				Ui:     ui,
				Client: godoCli,
			}, nil
		},
		"destroy": func() (cli.Command, error) {
			return &command.DestroyCommand{
				Ui:     ui,
				Client: godoCli,
			}, nil
		},
		"config": func() (cli.Command, error) {
			return &command.ConfigCommand{
				Ui:     ui,
				Config: config,
				Client: godoCli,
			}, nil
		},
		"ssh": func() (cli.Command, error) {
			return &command.SSHCommand{
				Ui:     ui,
				Client: godoCli,
			}, nil
		},
		"power": func() (cli.Command, error) {
			return &command.PowerCommand{
				Ui:     ui,
				Client: godoCli,
			}, nil
		},
		"snapshot": func() (cli.Command, error) {
			return &command.SnapshotCommand{
				Ui:     ui,
				Client: godoCli,
			}, nil
		},
		"shutdown": func() (cli.Command, error) {
			return &command.ShutdownCommand{
				Ui:     ui,
				Client: godoCli,
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(exitStatus, err)
	}
	os.Exit(exitStatus)
}
