package port

import "net/http"

type Router interface {
	Get(string, func(http.ResponseWriter, *http.Request))
	Post(string, func(http.ResponseWriter, *http.Request))
	Put(string, func(http.ResponseWriter, *http.Request))
	Patch(string, func(http.ResponseWriter, *http.Request))
	Delete(string, func(http.ResponseWriter, *http.Request))
}
