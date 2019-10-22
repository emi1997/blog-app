package api

import (
	"fmt"
	"net/http"
)


//EmptyCall handles requests with nothing after the slash
func EmptyCall(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}