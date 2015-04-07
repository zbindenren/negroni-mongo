package negronimongo

import (
	"net/http"

	"github.com/nbio/httpcontext"
	"gopkg.in/mgo.v2"
)

const (
	ContextKey = "mongosession" // ContextKey defines the key under wich the the database reference is stored in the request context.
)

// MongoMiddleware is a middleware handler that stores a mongo database reference in the request context.
type MongoMiddleware struct {
	Session *mgo.Session
}

// NewMongoMiddleware retunrs a new MongoMiddleware instance.
func NewMongoMiddleware(uri string) *MongoMiddleware {
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}
	return &MongoMiddleware{Session: session}
}

func (m *MongoMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	reqSession := m.Session.Clone()
	defer reqSession.Close()
	httpcontext.Set(r, ContextKey, m.Session)
	next(rw, r)
}
