//
package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/nanopack/sherpa/config"
)

type Template struct {
	CreatedAt       time.Time `json:"created_at"`       //
	Download        string    `json:"download"`         // where this template can be downloaded from
	ID              string    `json:"template_id"`      // NOTE: when pulling from the database, this will be organized alphabetically as ID
	TransformScript string    `json:"transform_script"` // the transform script to run when converting this into a build
}

//
func (t *Template) Fetch() error {
	fmt.Println("FETCH!", t.Download)

	// if the path is a git path then use git commands to download it

	// if the path is not a git path just download it
	res, err := http.Get(t.Download)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	templateDir := filepath.Join(config.Options.DataDir, "templates", filepath.Base(t.Download))

	// create the template file
	f, err := os.Create(templateDir)
	if err != nil {
		return err
	}
	defer f.Close()

	// save the downloaded template
	f.Write(b)

	return nil
}

// Register
func (t *Template) Register() (id int, err error) {

	db, err := sql.Open("postgres", config.Options.DBConn)
	if err != nil {
		fmt.Println("BONK!", err)
	}

	return id, db.QueryRow(`INSERT INTO templates(created_at, download, transform_script) VALUES($1, $2, $3) RETURNING template_id;`, time.Now(), t.Download, t.TransformScript).Scan(&id)
}

// Unregister
func (t *Template) Unregister() (err error) {

	// connect to the DB
	db, err := sql.Open("postgres", config.Options.DBConn)
	if err != nil {
		return err
	}

	// remove target template from database (by id)
	return db.QueryRow("DELETE FROM templates WHERE template_id=$1", t.ID).Scan()
}

// Transform
func (t *Template) Transform() {

}
