package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	//Add a title
	js.Global.Get("document").Set("title", "GopherJS Canvas Test")
	//See if it works
	js.Global.Call("alert", "Success...")
}