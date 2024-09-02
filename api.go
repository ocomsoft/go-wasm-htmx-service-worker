package main

import (
	"encoding/json"
	"fmt"
	"github.com/d5/tengo/v2"
	"io"
	"net/http"
	"sync/atomic"
	"wasm-http/data"
	"wasm-http/templates"

	wasmhttp "github.com/nlepage/go-wasm-http-server"
)

func main() {
	var counter int32

	http.HandleFunc("/calc", func(w http.ResponseWriter, req *http.Request) {
		if err := req.ParseForm(); err != nil {
			panic(err)
		}

		expr := req.FormValue("expr")
		values := map[string]interface{}{}

		for key, value := range req.Form {
			if len(value) > 0 {
				values[key] = value[0]
			}
		}

		res, err := tengo.Eval(req.Context(), expr, values)
		if err != nil {
			panic(err)
		}

		resStr := fmt.Sprintf("%v", res)
		templates.RenderExpr(expr, resStr).Render(req.Context(), w)
	})

	http.HandleFunc("/get_list", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
		if err != nil {
			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}

		var posts []data.Post
		if err := json.Unmarshal(body, &posts); err != nil {
			http.Error(w, "Failed to unmarshal JSON", http.StatusInternalServerError)
			return
		}

		templates.Posts(posts).Render(r.Context(), w)
	})
	http.HandleFunc("/hello2", func(res http.ResponseWriter, req *http.Request) {

		if err := req.ParseForm(); err != nil {
			panic(err)
		}
		name := req.FormValue("name")

		templates.Hello(name).Render(req.Context(), res)
	})

	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		params := make(map[string]string)
		if err := json.NewDecoder(req.Body).Decode(&params); err != nil {
			panic(err)
		}

		res.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(res).Encode(map[string]string{
			"message": fmt.Sprintf("Hello %s! (request %d)", params["name"], atomic.AddInt32(&counter, 1)),
		}); err != nil {
			panic(err)
		}
	})

	wasmhttp.Serve(nil)

	select {}
}
