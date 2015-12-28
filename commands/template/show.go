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
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "",
	Long:  ``,

	Run: show,
}

// show
func show(ccmd *cobra.Command, args []string) {

	path := fmt.Sprintf("%s/templates/%s", config.Options.URI, args[0])

	fmt.Println("PATH??", path)

	res, err := http.Get(path)
	if err != nil {
		fmt.Println("ERR!!", err)
	}
	defer res.Body.Close()

	// read the body
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERR!!", err)
	}
	defer res.Body.Close()

	fmt.Println("THING!!??", string(b))

	// if err := json.Unmarshal(b, &v); err != nil {
	// 	return err
	// }
}
