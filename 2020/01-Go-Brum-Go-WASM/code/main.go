package main

import (
	"syscall/js"
	"fmt"
	"github.com/endeveit/guesslanguage"
)

func greet(this js.Value, i []js.Value) interface{} {
	fmt.Println(i)
	output := "Hello, " + i[0].String()
	return output
}

func whatLang(this js.Value, i []js.Value) interface{} {
	fmt.Println(i[0].String())
	lang := guesslanguage.GuessName(i[0].String())
	return lang
}

func registerCallbacks() {
	js.Global().Set("greet", js.FuncOf(greet))
	js.Global().Set("whatlang", js.FuncOf(whatLang))
}

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("WASM hi Go Initialized")
	// register functions
	registerCallbacks()
	// we listen on a channel that's never used to
	// make the program run forever
	<-c
}
