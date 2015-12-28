//
package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/jcelliott/lumber"
)

const (
	DEFAULT_DIR  = "/var/db/sherpa"
	DEFAULT_HOST = "0.0.0.0"
	DEFAULT_PORT = ":6430"
	VERSION      = "0.0.1"
)

//
var (
	config  string        // /path/to/config (if you want to config more than just the DB location)
	Log     lumber.Logger //
	Options Opts          // configuration options
)

//
type Opts struct {
	BuildGCInterval           int    `json:"build_gc_interval"`             // how often to remove failed builds (default 1hr)
	BuildGCMarkFailedInterval int    `json:"build_gc_mark_failed_interval"` // how long to wait before marking a build as failed (default 60sec)
	DataDir                   string `json:"data_dir"`                      // (/var/db/sherpa by default)
	DBConn                    string `json:"db_conn"`                       // postgres://username:password@ip:port/database (postgres://postgres@localhost/sherpa)
	Host                      string `json:"host"`                          // Binding IP
	Port                      string `json:"port"`                          // Binding Port
	TemplateGCInterval        int    `json:"template_gc_interval"`          // how often to remove unknown templates (default 1hr)
	URI                       string //
}

// NOTICE: set this up to use a map instead of an Opts struct

// create a new logger
func init() {
	Log = lumber.NewConsoleLogger(lumber.INFO)

	// create a new opts with default values
	Options = Opts{
		BuildGCInterval:           3600,
		BuildGCMarkFailedInterval: 60,
		DataDir:                   DEFAULT_DIR,
		DBConn:                    "postgres://postgres@localhost/sherpa", // not specifying a port will default to the default postgres port (:5432)
		Host:                      DEFAULT_HOST,
		Port:                      DEFAULT_PORT,
		TemplateGCInterval:        3600,
		URI:                       "http://" + DEFAULT_HOST + DEFAULT_PORT,
	}

	fmt.Printf("THING! %#v\n", Options)
}

// Parse
func Parse(path string) error {

	// if a config is provided (and found), parse the config file overwriting any
	// options
	if fp, err := filepath.Abs(path); err == nil {

		//
		f, err := ioutil.ReadFile(fp)
		if err != nil {
			return err
		}

		// parse config file
		if err := yaml.Unmarshal(f, &Options); err != nil {
			return err
		}

		// override the uri incase host/port were passed
		Options.URI = "http://" + Options.Host + Options.Port
	}

	return nil
}
