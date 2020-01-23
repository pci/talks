package main

import (
	"syscall/js"
	"fmt"
)

func add(this js.Value, i []js.Value) interface{} {
	output := js.ValueOf(i[0].Int() + i[1].Int())
	js.Global().Set("output", output)
	fmt.Println(output.String())
	return output
}

func subtract(this js.Value, i []js.Value) interface{} {
	output := js.ValueOf(i[0].Int() - i[1].Int())
	js.Global().Set("output", output)
	fmt.Println(output.String())
	return output
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	// we listen on a channel that's never used to
	// make the program run forever
	<-c
}
