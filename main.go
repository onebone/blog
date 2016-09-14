package main

import (
	"log"
	"net/http"

	pongo "github.com/flosch/pongo2"
	"github.com/gorilla/mux"
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
		log.Println(err.Error())
		return false
	}

	if _, ok := val["title"]; !ok {
		val["title"] = "Blog"
	}

	err = tmpl.ExecuteWriter(val, res)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func main() {
	log.Println("Starting blog...")

	r := mux.NewRouter()

	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		render("index", pongo.Context{}, res)
	})

	r.HandleFunc("/post/{id:[0-9]+}", func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		res.Write([]byte("Matched: " + vars["id"]))
	})

	r.HandleFunc("/admin", func(res http.ResponseWriter, req *http.Request) {
		render("admin/login", pongo.Context{}, res)
	})

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	r.NewRoute().MatcherFunc(func(req *http.Request, rm *mux.RouteMatch) bool {
		return true
	}).HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("Oops! The page cannot be found."))
	})

	http.ListenAndServe(":1234", r)
}
