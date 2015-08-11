package command

import (
    "flag"
    "fmt"
    "github.com/digitalocean/godo"
    "github.com/mitchellh/cli"
    "strings"
)

type DomainsCommand struct {
    Ui     cli.Ui
    Client *godo.Client
}

func (c *DomainsCommand) Help() string {
    helpText := `
Usage: godo-cli domains
`
    return strings.TrimSpace(helpText)
}

func (c *DomainsCommand) parse(args []string) (*DomainsFlags, error) {
    flags := new(DomainsFlags)

    cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)

    cmdFlags.BoolVar(&flags.Quiet, "q", false, "")

    err := cmdFlags.Parse(args)

    if err != nil {
        return nil, err
    }

    return flags, nil
}

func (c *DomainsCommand) Run(args []string) int {
    flags, err := c.parse(args)

    if err != nil {
        c.Ui.Error(fmt.Sprintf("Failed to parse %v", err))
        return -1
    }

    util := GodoUtil{Client: c.Client}
    domains, err := util.GetDomains()
    if err != nil {
        c.Ui.Error(fmt.Sprintf("Failed to request %v", err))
        return -1
    }

    for _, domain := range domains {
        fmt.Printf("%s (Name: %s, TTL: %s, ZoneFile :%s)\n",
            domain.Name, 
            domain.TTL, 
            domain.ZoneFile 
        )
    }
    return 0
}

func (c *DomainsCommand) Synopsis() string {
    return fmt.Sprintf("Retrieve the list of domains.")
}

