package main

import (
	"fmt"
	"log"

	"./modules/myblog"
	"./modules/mycontext"
	"./modules/utils"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//var blog = &mycontext.Blog1

var ctx *mycontext.Context

//Run ...
func Run(FileServerDir string, context *mycontext.Context) {

	ctx = context
	r := mux.NewRouter()

	//r.PathPrefix("/static/").
	//	Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(FileServerDir)))).
	//	Methods("GET")

	r.Path("/api/").
		HandlerFunc(apiHandler).
		Methods("GET")

	r.Path("/api/blog/articles").
		HandlerFunc(articlesGetHandler).
		Queries("offset", "{offset}", "count", "{count}"). //"{[0-9]*?}"
		Methods("GET")

	r.Path("/api/blog/articles").
		HandlerFunc(articlesPostHandler).
		Methods("POST")

	r.Path("/api/blog").
		HandlerFunc(blogHandler).
		Methods("GET")

	http.Handle("/api/", r)

	http.Handle("/", http.FileServer(http.Dir(FileServerDir)))

	//srv = &http.Server{
	//	Handler: r,
	//	Addr:    "127.0.0.1:8000",
	//	// Good practice: enforce timeouts for servers you create!
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}

	log.Fatal(http.ListenAndServe(":8000", nil))

}

//Abort ...
func Abort() {

	//ctx, cancel := context.WithTimeout(context.Background(), wait)
	//defer cancel()

	//srv.Shatdown()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.

}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	lg(r.URL.String())

	fmt.Fprint(w, r.URL)
}

// /api/blog
// return {Autor: string, Count: int}
func blogHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	lg(r.URL.String())

	respondWithJSON(w, http.StatusOK, ctx.Blogs[0].About())
}

// /api/blog/articles?offset={}&count={}
// return []Article
func articlesGetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	lg(r.URL.String())

	fmt.Println(mux.Vars(r))

	_offset := mux.Vars(r)["offset"]

	_count := mux.Vars(r)["count"]

	offset := utils.Str2UInt(_offset, 0)
	if offset > ctx.Blogs[0].Count()-1 {
		offset = ctx.Blogs[0].Count() - 1
	}

	if offset < 0 {
		offset = 0
	}

	count := utils.Str2UInt(_count, 10)

	if offset+count > ctx.Blogs[0].Count() {
		count = ctx.Blogs[0].Count() - offset
	}

	if count < 0 {
		count = 0
	}

	articles := ctx.Blogs[0].Articles[offset : offset+count]

	respondWithJSON(w, http.StatusOK, articles)

}

func articlesPostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	lg(r.URL.String())

	_, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	var a myblog.Article

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&a); err != nil {
		ctx.Blogs[0].Articles = append(ctx.Blogs[0].Articles, a)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func lg(msg string) {
	fmt.Println(msg)
}
