package main // import "hello-goja-moment"

import (
	"fmt"
	"os"

	"github.com/dop251/goja"
)

func evalString(vm *goja.Runtime, cmd string) goja.Value {
	v, err := vm.RunString(cmd)
	if err != nil {
		panic(err)
	}
	return v
}

func loadJS(vm *goja.Runtime, file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	_, err = vm.RunString(string(data))
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Hello, Goja!")

	vm := goja.New()
	// evalString(vm, "var global = (function(){ return this; }).call(null);")

	loadJS(vm, "moment.min.js")

	v := evalString(vm, "moment().format('YYYY-MM-DD HH:mm:ss')")
	s := v.String()
	fmt.Println(s)
}
