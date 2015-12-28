// package commands ...
package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nanopack/sherpa/api"
	"github.com/nanopack/sherpa/commands/build"
	"github.com/nanopack/sherpa/commands/template"
	"github.com/nanopack/sherpa/config"
	"github.com/nanopack/sherpa/models"
)

//
var (

	//
	SherpaCmd = &cobra.Command{
		Use:   "sherpa",
		Short: "",
		Long:  ``,

		Run: func(ccmd *cobra.Command, args []string) {

			// if there is a config file provided, proceed to boot sherpa
			if conf != "" {

				// parse the config
				if err := config.Parse(conf); err != nil {
					fmt.Println("FAILED TO PARSE!", err.Error())
				}

				// initialize models
				if err := models.Init(); err != nil {
					fmt.Println("FAILED TO INITIALIZE!", err.Error())
				}

				// start up the API
				if err := api.Start(); err != nil {
					fmt.Println("FAILED TO START!", err.Error())
				}

				return
			}

			// fall back on default help if no args/flags are passed
			ccmd.HelpFunc()(ccmd, args)
		},
	}

	//
	version bool   // display the version of sherpa
	conf    string // path to config options
)

// init creates the list of available sherpa commands and sub commands
func init() {

	// local flags
	SherpaCmd.Flags().BoolVarP(&version, "version", "v", false, "Display the current version of this CLI")
	SherpaCmd.Flags().StringVarP(&conf, "config", "", "", "Path to config options")

	// 'sherpa' commands

	// subcommands
	SherpaCmd.AddCommand(build.BuildCmd)
	SherpaCmd.AddCommand(template.TemplateCmd)
}
