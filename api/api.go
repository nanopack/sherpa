//
package api

import (
	"net/http"

	"github.com/gorilla/pat"

	"github.com/nanopack/sherpa/config"
)

// Start creates a new http server listner
func Start() error {

	//
	routes, err := registerRoutes()
	if err != nil {
		return err
	}

	//
	config.Log.Info("Starting sherpa server (listening on port %v)...\n", config.Options.Port)

	// blocking...
	if err := http.ListenAndServe(config.Options.Port, routes); err != nil {
		return err
	}

	return nil
}

// registerRoutes registers all api routes with the router
func registerRoutes() (*pat.Router, error) {
	config.Log.Debug("[sherpa/api] Registering routes...\n")

	//
	router := pat.New()

	//
	router.Get("/ping", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("pong"))
	})

	// templates
	router.Delete("/templates/{id}", handleRequest(deleteTemplate))
	router.Get("/templates/{id}", handleRequest(getTemplate))
	router.Post("/templates", handleRequest(postTemplate))
	router.Get("/templates", handleRequest(getTemplates))

	// builds
	router.Delete("/builds/{id}", handleRequest(deleteBuild))
	router.Get("/builds/{id}", handleRequest(getBuild))
	router.Post("/builds", handleRequest(postBuild))
	router.Get("/builds", handleRequest(getBuilds))

	return router, nil
}

// handleRequest is a wrapper for the actual route handler, simply to provide some
// debug output
func handleRequest(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		config.Log.Debug(`
Request:
--------------------------------------------------------------------------------
%+v

`, req)

		//
		fn(rw, req)

		config.Log.Debug(`
Response:
--------------------------------------------------------------------------------
%+v

`, rw)
	}
}

// parseBody
// func parseBody(req *http.Request, v interface{}) error {
//
// 	//
// 	b, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		return err
// 	}
//
// 	defer req.Body.Close()
//
// 	//
// 	if err := json.Unmarshal(b, v); err != nil {
// 		return err
// 	}
//
// 	return nil
// }

// writeBody
// func writeBody(v interface{}, rw http.ResponseWriter, status int) error {
// 	b, err := json.Marshal(v)
// 	if err != nil {
// 		return err
// 	}
//
// 	rw.Header().Set("Content-Type", "application/json")
// 	rw.WriteHeader(status)
// 	rw.Write(b)
//
// 	return nil
// }
