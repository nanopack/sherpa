//
package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/nanopack/sherpa/config"
	"github.com/nanopack/sherpa/models"
)

//
func getBuilds(rw http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", config.Options.DBConn)
	if err != nil {
		fmt.Println("BONK!", err)
	}

	rows, err := db.Query(`SELECT * FROM builds`)
	if err != nil {
		fmt.Println("ERR!", err.Error())
	}
	defer rows.Close()

	builds := []models.Build{}

	//
	for rows.Next() {
		t := models.Build{}
		err = rows.Scan()
		builds = append(builds, t)
	}

	err = rows.Err() // get any error encountered during iteration

	fmt.Printf("BUILDS? %#v\n", builds)
}

//
func postBuild(rw http.ResponseWriter, req *http.Request) {

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

	// create a new build
	build := models.Build{}
	if err := json.Unmarshal(b, &build); err != nil {
		fmt.Println("MARSHALL ERR!", err.Error())
		return
	}

	fmt.Printf("BUILD! %#v\n", build)

	// Create a build in the database
	db, err := sql.Open("postgres", config.Options.DBConn)
	if err != nil {
		fmt.Println("BONK!", err)
	}

	// 1. Create entry in DB that represents the build
	if err := db.QueryRow(`INSERT INTO builds(created_at, meta_data, template_id, transform_payload, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING build_id;`, time.Now(), build.MetaData, build.TemplateID, build.TransformPayload, time.Now()).Scan(&build.ID); err != nil {
		fmt.Println("BINK!", err)
	}

	// once the build is done send on this channel to close anything
	done := make(chan bool)

	// 1.5 At a short interval, update last_updated field (goroutine) because packer run might take time
	go func() {
		for {
			select {

			//
			case <-time.After(time.Second * 1):
				if err := db.QueryRow(`UPDATE builds SET updated_at=$1 WHERE build_id=$2;`, time.Now(), build.ID); err != nil {
					fmt.Println("BINK! %#v\n", err)
				}

			//
			case <-done:
				fmt.Println("DONE!!")
				return
			}
		}
	}()

	// fetch the template; we fetch the template because we MIGHT need its `download`, but we WILL needs its `transform_script`
	template := models.Template{}
	if err := db.QueryRow(`SELECT * FROM templates WHERE template_id=$1`, build.TemplateID).Scan(&template.CreatedAt, &template.Download, &template.ID, &template.TransformScript); err != nil {
		fmt.Println("ERR!", err.Error())
	}

	fmt.Printf("TEMPLATE! %#v\n", template)

	// 2. Install template on local node if not already available ([/var/db/sherpa]/templates/:id)
	templateDir := filepath.Join(config.Options.DataDir, "templates", template.ID)
	if _, err := os.Stat(templateDir); err != nil {

		//
		template.Fetch()
	}

	// 3. Copy template into a tmp build folder for duration of build ([/var/db/sherpa]/builds/:build_id)
	buildDir := filepath.Join(config.Options.DataDir, "builds", build.ID)
	if err := os.Mkdir(buildDir, 0755); err != nil {
		fmt.Println("BOONK!", err.Error())
	}

	fmt.Println("THINGS", templateDir, buildDir)

	//
	done <- true

	return

	// 4. Run transform script (builds/:build_id/{{transform_script}}) if present, passing transform_payload
	if template.TransformScript != "" {
		out, err := exec.Command(template.TransformScript, build.TransformPayload).CombinedOutput()
		if err != nil {
			fmt.Println("BUNK!", string(out))
		}
	}

	// 5. Packer run (all output (\n) is a new build_log for that build); update last_updated/state of build throughout process

	// BuildLog struct {
	// 	BuildID   string    `json:"build_id"`     //
	// 	CreatedAt time.Time `json:"created_at"`   //
	// 	ID        string    `json:"build_log_id"` //
	// 	Message   string    `json:"message"`      //
	// }

	// 6. Record - take packer info and dump it into meta_data

	// 7. Cleanup - delete build folder
	if err := os.Remove(buildDir); err != nil {
		fmt.Println("BOONK!", err.Error())
	}
}

//
func getBuild(rw http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("postgres", config.Options.DBConn)
	if err != nil {
		fmt.Println("BONK!", err)
	}

	// create a build and read the build from the database
	b := models.Build{}
	if err := db.QueryRow(`SELECT * FROM builds WHERE build_id=$1`, req.URL.Query().Get(":id")).Scan(); err != nil {
		fmt.Println("ERR!", err.Error())
	}

	fmt.Printf("BUILD! %#v", b)
}

//
func deleteBuild(rw http.ResponseWriter, req *http.Request) {

	// connect to the DB
	db, err := sql.Open("postgres", config.Options.DBConn)
	if err != nil {
		fmt.Println("BONK!", err)
	}

	// remove target build from database (by id)
	if err := db.QueryRow("DELETE FROM builds WHERE build_id=$1", req.URL.Query().Get(":id")).Scan(); err != nil {
		fmt.Println("BINK!", err)
	}
}
