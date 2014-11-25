/**
* This plugin is an example plugin that allows a user to call a cli-command
* by typing `cf cli-command name-of-command args.....`. This plugin also prints
* the output returned by the CLI when a cli-command is invoked.
 */
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
	cliConnection.CliCommand("target", "-o", orgName)
	cliConnection.CliCommand("create-org", orgName)
	cliConnection.CliCommand("target", "-o", orgName)
	cliConnection.CliCommand("target", "-o", orgName, "-s", spaceName)
	cliConnection.CliCommand("create-space", spaceName, "-o", orgName)
	cliConnection.CliCommand("target", "-o", orgName, "-s", spaceName)
	fmt.Print("Created org")
	return
	// Invoke the cf command passed as the set of arguments
	// after the first argument.
	//
	// Calls to plugin.CliCommand([]string) must be done after the invocation
	// of plugin.Start() to ensure the environment is bootstrapped.

	// The call to plugin.CliCommand() returns an error if the cli command
	// returns a non-zero return code or panics. The output written by the CLI
	// is returned in any case.
	// if err != nil {
	// 	fmt.Println("PLUGIN ERROR: Error from CliCommand: ", err)
	// }

	// // Print the output returned from the CLI command.
	// fmt.Println("")
	// fmt.Println("---------- Command output from the plugin ----------")
	// for index, val := range output {
	// 	fmt.Println("#", index, " value: ", val)
	// }
	// fmt.Println("----------              FIN               -----------")
}
