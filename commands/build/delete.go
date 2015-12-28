//
package build

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/nanopack/sherpa/config"
)

//
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "",
	Long:  ``,

	Run: delete,
}

// delete
func delete(ccmd *cobra.Command, args []string) {

	// an HTTP request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/builds/%s", config.Options.URI, args[0]), nil)
	if err != nil {
		fmt.Println("ERR!!", err)
	}

	req.Header.Set("Content-Type", "application/json")

	//
	if _, err := http.DefaultClient.Do(req); err != nil {
		fmt.Println("ERR!!", err)
	}

	// // read the body
	// b, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return err
	// }
	// defer res.Body.Close()
	//
	// // unmarshal response into the provided container. If no container was given
	// // it's mostly likely a raw request where the body isn't needed, and therfore
	// // this step can be skipped.
	// if v != nil {
	// 	if err := json.Unmarshal(b, &v); err != nil {
	// 		return err
	// 	}
	// }
}
