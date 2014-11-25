package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type OrgSpace struct{}

func (c *OrgSpace) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "OrgSpace",
		Commands: []plugin.Command{
			{
				Name:     "org-space",
				HelpText: "Command to call cli command. It passes all arguments through to the command",
			},
		},
	}
}

func main() {
	plugin.Start(new(OrgSpace))
}

func (c *OrgSpace) Run(cliConnection plugin.CliConnection, args []string) {

	if len(args) != 3 {
		fmt.Print("You must provide an org and space i.e. org-space org space")
		return
	}

	orgName := args[1]
	spaceName := args[2]
	cliConnection.CliCommandWithoutTerminalOutput("create-org", orgName)
	cliConnection.CliCommandWithoutTerminalOutput("create-space", spaceName, "-o", orgName)
	cliConnection.CliCommandWithoutTerminalOutput("target", "-o", orgName, "-s", spaceName)

	fmt.Printf("Org %s and Space %s is now available and targeted", orgName, spaceName)
	return
}
