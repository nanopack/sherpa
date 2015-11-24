//
package main

import (
	"fmt"

	"github.com/nanopack/sherpa/api"
)

//
func main() {

	//
	if err := api.Start(""); err != nil {
		fmt.Println("FAILED TO START!", err.Error())
	}
}
