package main

import (
	"fmt"

	"code.cloudfoundry.org/cli/plugin"
)

type AppCloudPlugin struct{}

func (p *AppCloudPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Swisscom Application Cloud",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "view-tree",
				HelpText: "View organization tree",
				UsageDetails: plugin.Usage{
					Usage: "view-tree",
				},
			},
		},
	}
}

// Run initiates the plugin
func (p *AppCloudPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	var err error

	switch args[0] {
	case "view-tree":

		err = p.ViewTree(cliConnection)

		if err != nil {
			fmt.Printf("\n%s\n", redBold(err.Error()))
		}
	}
}

func main() {
	plugin.Start(new(AppCloudPlugin))
}
