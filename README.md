# Wasm Http HTMLX Server

This is a quick Proof of Concept test of using `github.com/nlepage/go-wasm-http-server` to serve files in the Service Worker using HTMLX generating template using `github.com/a-h/templ`

## Links
- https://templ.guide/
- https://htmx.org/docs/#introduction
- Post data from : https://jsonplaceholder.typicode.com/posts

## To Run

```bash
templ generate
GOOS=js GOARCH=wasm go build -o api.wasm .
```
(these are in build.sh if you need)

Then use a webserver (ie preview in your IDE) to serve `index.html` and it should load the other files.


If the changes are NOT visible, "unregister" the service worker in Chrome Dev tools and Reload.

