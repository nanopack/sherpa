//
package api

import (
	"net/http"
)

//
func getBuilds(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("getBuilds"))
}

//
func postBuild(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("postBuilds"))
}

//
func getBuild(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("getBuild"))
}
