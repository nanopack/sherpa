//
package build

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/nanopack/sherpa/config"
)

//
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Long:  ``,

	Run: create,
}

// `{"transform_script":"/some/path","download":"/another/path"}`

// create
func create(ccmd *cobra.Command, args []string) {
	if _, err := http.Post(fmt.Sprintf("%s/builds", config.Options.URI), "application/json", bytes.NewBufferString(args[0])); err != nil {
		fmt.Println("ERR!!", err)
	}
}
