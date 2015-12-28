//
package git

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

//
func Clone(path string) {
	if !exists() {
		fmt.Printf(`Git is not available on this machine. Use "Download" to fetch template releases\n`)
		return
	}

	os.Exec("git", "clone", path).Run()
}

//
// "https://api.github.com/repos/nanobox-io/pagodabox-cli/contents/version"
func Download(path string) error {

	//
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	// tell github we want raw format
	req.Header.Set("Accept", "application/vnd.github.V3.raw")

	//
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	//
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// exists checks to see if "git" is available. If not then some commands wont work
// and downloading will be the only method.
func exists() (exists bool) {
	if err := exec.Command("which", "git").Run(); err == nil {
		exists = true
	}

	return
}
