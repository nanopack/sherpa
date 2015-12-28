// package template ...
package template

import "github.com/spf13/cobra"

//
var TemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "",
	Long:  ``,
}

//
func init() {
	TemplateCmd.AddCommand(createCmd)
	TemplateCmd.AddCommand(deleteCmd)
	TemplateCmd.AddCommand(listCmd)
	TemplateCmd.AddCommand(showCmd)
}
