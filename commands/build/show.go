//
package build

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/nanopack/sherpa/config"
)

//
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "",
	Long:  ``,

	Run: show,
}

// show
func show(ccmd *cobra.Command, args []string) {
	if _, err := http.Get(fmt.Sprintf("%s/builds/%s", config.Options.URI, args[0])); err != nil {
		fmt.Println("ERR!!", err)
	}
}
