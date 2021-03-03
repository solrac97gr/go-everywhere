package main

import (
	"strconv"
	"syscall/js"
)

func add(this js.Value,i []js.Value) interface{}{
	value1:=js.Global().Get("document").Call("getElementById",i[0].String()).Get("value").String()
	value2:=js.Global().Get("document").Call("getElementById",i[1].String()).Get("value").String()

	int1,_:=strconv.Atoi(value1)
	int2,_:=strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById",i[2].String()).Set("value",int1+int2)
	return nil
}

func subtract(this js.Value,i []js.Value) interface{}{
	value1:=js.Global().Get("document").Call("getElementById",i[0].String()).Get("value").String()
	value2:=js.Global().Get("document").Call("getElementById",i[1].String()).Get("value").String()

	int1,_:=strconv.Atoi(value1)
	int2,_:=strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById",i[2].String()).Set("value",int1-int2)
	return nil
}

func multiplies(this js.Value,i []js.Value) interface{}{
	value1:=js.Global().Get("document").Call("getElementById",i[0].String()).Get("value").String()
	value2:=js.Global().Get("document").Call("getElementById",i[1].String()).Get("value").String()

	int1,_:=strconv.Atoi(value1)
	int2,_:=strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById",i[2].String()).Set("value",int1*int2)
	return nil
}

func divide(this js.Value,i []js.Value) interface{}{
	value1:=js.Global().Get("document").Call("getElementById",i[0].String()).Get("value").String()
	value2:=js.Global().Get("document").Call("getElementById",i[1].String()).Get("value").String()

	int1,_:=strconv.Atoi(value1)
	int2,_:=strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById",i[2].String()).Set("value",float64(int1)/float64(int2))
	return nil
}

func registerCallBacks(){
	js.Global().Set("add",js.FuncOf(add))
	js.Global().Set("subtract",js.FuncOf(subtract))
	js.Global().Set("multiplies",js.FuncOf(multiplies))
	js.Global().Set("divide",js.FuncOf(divide))
}

func main()  {
	c:=make(chan struct{},0)
	println("Go WebAssembly Initialized")
	registerCallBacks()
	<-c
}