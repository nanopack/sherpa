// package build ...
package build

import "github.com/spf13/cobra"

//
var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "",
	Long:  ``,
}

//
func init() {
	BuildCmd.AddCommand(createCmd)
	BuildCmd.AddCommand(deleteCmd)
	BuildCmd.AddCommand(listCmd)
	BuildCmd.AddCommand(showCmd)
}
