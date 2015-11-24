//
package config

import (
	"github.com/jcelliott/lumber"
)

//
var (
	Log lumber.Logger
)

//
func init() {

	// create a new logger
	Log = lumber.NewConsoleLogger(lumber.INFO)
}
