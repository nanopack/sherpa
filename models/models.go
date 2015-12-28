package models

import (
	"os"
	"path/filepath"

	"github.com/nanopack/sherpa/config"
)

// Init creates the directory where builds and templates are going to live
func Init() (err error) {

	// create a /var/db/sherpa/builds (if it doesn't already exist)
	if err = os.MkdirAll(filepath.Join(config.Options.DataDir, "builds"), 0755); err != nil {
		return
	}

	// create a [/var/db/sherpa]/templates (if it doesn't already exist)
	return os.MkdirAll(filepath.Join(config.Options.DataDir, "templates"), 0755)
}
