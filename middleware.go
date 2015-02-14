package negronimongo

import (
	"net/http"

	"github.com/nbio/httpcontext"
	"gopkg.in/mgo.v2"
)

const (
	ContextKey = "database" // ContextKey defines the key under wich the the database reference is stored in the request context.
)

// MongoMiddleware is a middleware handler that stores a mongo database reference in the request context.
type MongoMiddleware struct {
	database string
	Session  *mgo.Session
}

// NewMongoMiddleware retunrs a new MongoMiddleware instance.
func NewMongoMiddleware(url string, database string) *MongoMiddleware {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	session.DB(database)
	return &MongoMiddleware{Session: session, database: database}
}

func (m *MongoMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	reqSession := m.Session.Clone()
	defer reqSession.Close()
	db := reqSession.DB(m.database)
	httpcontext.Set(r, ContextKey, db)
	next(rw, r)
}
