package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/nbio/httpcontext"
	"github.com/zbindenren/negroni-mongo"
	"gopkg.in/mgo.v2"
)

func main() {
	n := negroni.New()
	// connect to test database on localhost
	m := negronimongo.NewMongoMiddleware("localhost", "test")
	// change session settings if you want
	m.Session.EnsureSafe(&mgo.Safe{FSync: true})
	n.Use(m)

	r := http.NewServeMux()
	r.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		// get the database reference from context
		db := httpcontext.Get(r, negronimongo.ContextKey).(*mgo.Database)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "got db reference with name %s\n", db.Name)
	})

	n.UseHandler(r)

	n.Run(":3000")
}
