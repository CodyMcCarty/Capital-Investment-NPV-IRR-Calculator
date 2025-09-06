//go:build js && wasm

package main

import (
	"strconv"
	"syscall/js"
)

func double(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return js.ValueOf(0)
	}
	return js.ValueOf(args[0].Float() * 2)
}

var onClick js.Func // keep a reference so GC won't collect it

func main() {
	doc := js.Global().Get("document")
	input := doc.Call("getElementById", "n")
	btn := doc.Call("getElementById", "btn")
	out := doc.Call("getElementById", "out")

	// Optional: show we booted
	js.Global().Get("console").Call("log", "WASM booted; wiring click handler")

	onClick = js.FuncOf(func(this js.Value, args []js.Value) any {
		val := input.Get("value").String()
		x, err := strconv.ParseFloat(val, 64)
		if err != nil {
			out.Set("textContent", "Please enter a valid number.")
			return nil
		}
		y := x * 2
		out.Set("textContent", "Result: "+strconv.FormatFloat(y, 'f', -1, 64))
		return nil
	})
	btn.Call("addEventListener", "click", onClick)

	// Clean up on page unload (optional)
	js.Global().Call("addEventListener", "beforeunload",
		js.FuncOf(func(this js.Value, args []js.Value) any { onClick.Release(); return nil }))

	select {} // keep the Go program alive
}
