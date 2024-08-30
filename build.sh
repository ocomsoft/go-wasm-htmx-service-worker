#!/bin/bash

templ generate
GOOS=js GOARCH=wasm go build -o api.wasm .