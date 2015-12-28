//
package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/nanopack/sherpa/config"
	"github.com/nanopack/sherpa/models"
)

// getTemplates
func getTemplates(rw http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", config.Options.DBConn)
	if err != nil {
		fmt.Println("BONK!", err)
	}

	rows, err := db.Query(`SELECT * FROM templates`)
	if err != nil {
		fmt.Println("ERR!", err.Error())
	}
	defer rows.Close()

	templates := []models.Template{}

	//
	for rows.Next() {
		t := models.Template{}
		err = rows.Scan(&t.ID, &t.Download, &t.TransformScript)
		templates = append(templates, t)
	}

	err = rows.Err() // get any error encountered during iteration

	fmt.Printf("TEMPLATES? %#v\n", templates)

	b, err := json.Marshal(templates)
	if err != nil {
		fmt.Println("ERR!", err.Error())
	}

	rw.Write(b)
}

// postTemplate
func postTemplate(rw http.ResponseWriter, req *http.Request) {

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		fmt.Println("BOINK!", err.Error())
	}

	fmt.Println(`
Request:
--------------------------------------------------------------------------------
` + string(dump))

	// read the request body
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("READ ALL ERR!", err.Error())
		return
	}
	defer req.Body.Close()

	fmt.Println("BODY?", string(b))

	// create a new template
	t := models.Template{}
	if err := json.Unmarshal(b, &t); err != nil {
		fmt.Println("MARSHALL ERR!", err.Error())
		return
	}

	fmt.Printf("TEMPLATE?? %#v\n", t)

	// register the template (write it to the database)
	id, err := t.Register()
	if err != nil {
		fmt.Println("BONK!", err.Error())
		return
	}

	fmt.Println("Registered template", id)
}

// getTemplate
func getTemplate(rw http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("postgres", config.Options.DBConn)
	if err != nil {
		fmt.Println("BONK!", err)
	}

	fmt.Println("ID??", req.URL.Query().Get(":id"))

	// create a template and read the template from the database
	t := models.Template{}
	if err := db.QueryRow(`SELECT * FROM templates WHERE template_id=$1`, req.URL.Query().Get(":id")).Scan(&t.CreatedAt, &t.Download, &t.ID, &t.TransformScript); err != nil {
		fmt.Println("ERR!", err.Error())
	}

	fmt.Printf("TEMPLATE! %#v\n", t)

	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println("ERR!", err.Error())
	}

	//
	rw.Write(b)
}

// deleteTemplate
func deleteTemplate(rw http.ResponseWriter, req *http.Request) {

	// create a template and unregister it (remove it from the database)
	t := models.Template{ID: req.URL.Query().Get(":id")}
	if err := t.Unregister(); err != nil {
		fmt.Println("BONK!!!", err.Error())
	}

	fmt.Println("Unregistered template", t.ID)
}
