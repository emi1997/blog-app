package api

import (
	"net/http"
)

//EmptyCall handles requests with nothing after the slash and serves the home site
func EmptyCall(w http.ResponseWriter, r *http.Request) {
	//indicates where the files are located
	http.FileServer(http.Dir("../html"))
}
