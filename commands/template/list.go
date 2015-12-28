//
package template

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/nanopack/sherpa/config"
)

//
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  ``,

	Run: list,
}

// list
func list(ccmd *cobra.Command, args []string) {

	path := fmt.Sprintf("%s/templates", config.Options.URI)

	res, err := http.Get(path)
	if err != nil {
		fmt.Println("ERR!!", err)
	}

	// read the body
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERR!!", err)
	}
	defer res.Body.Close()

	fmt.Println("THING!!??", string(b))
}
