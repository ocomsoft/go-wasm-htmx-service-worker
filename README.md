# Wasm Http HTMLX Server

This is a test of using `github.com/nlepage/go-wasm-http-server` to serve files in the Service Worker using HTMLX generating template using `github.com/a-h/templ`

## To Run

```bash
templ generate
GOOS=js GOARCH=wasm go build -o api.wasm .
```
(these are build.sh if you need)

Then use a websever (ie preview in your IDE) to serve `index.html` and it should load the other files.

If the changes are NOT visible "unregister" the service worker in Chrome Dev tools and Reload.

