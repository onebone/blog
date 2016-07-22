package main

import (
	"github.com/gorilla/mux"
	pongo "github.com/flosch/pongo2"
	"net/http"
)

/*var tmpls map[string]*pongo.Template

func init() {
	tmpls = map[string]*pongo.Template {
	 	"index": pongo.Must(pongo.FromFile("views/index.pongo")),
 	}
}*/

func render(name string, val pongo.Context, res http.ResponseWriter) bool {
	tmpl, err := pongo.FromFile("views/" + name + ".pongo")
	if err != nil {
		return false
	}

	tmpl.ExecuteWriter(val, res)
	return true
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		render("index", pongo.Context{}, res)
	})

	r.HandleFunc("/post/{id:[0-9]+}", func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		res.Write([]byte("Matched: " + vars["id"]))
	})

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	http.ListenAndServe(":1234", r)
}
