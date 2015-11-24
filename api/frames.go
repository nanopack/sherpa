//
package api

import (
	"net/http"
)

//
func getFrames(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("getFrames"))
}

//
func postFrame(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("postFrames"))
}

//
func getFrame(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("getFrame"))
}
