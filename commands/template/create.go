//
package template

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

// create
func create(ccmd *cobra.Command, args []string) {

	path := fmt.Sprintf("%s/templates", config.Options.URI)
	body := bytes.NewBufferString(args[0])

	if _, err := http.Post(path, "application/json", body); err != nil {
		fmt.Println("ERR!!", err)
	}
}
