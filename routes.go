package main

import (
	"github.com/gorilla/mux"

	"test.com/server/apis"
)

func routes(router *mux.Router) {

	router.HandleFunc("/", apis.ApiDefault)
	router.HandleFunc("/test", apis.ApiTest)
	router.HandleFunc("/name/{name}/age/{age}", apis.ApiRoute)
	router.HandleFunc("/query", apis.ApiQueryParam)
	router.HandleFunc("/morequery", apis.ApiQueryParams)

	router.HandleFunc("/post", apis.ApiPost)

}
