package temp

const NETHTTP = `package main

import "net/http"

func initRouter() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(conf.Port, nil)
}`
