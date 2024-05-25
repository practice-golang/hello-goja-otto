package main

import (
	"fmt"
	"os"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

func loadToGoja(vm *goja.Runtime, file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	_, err = vm.RunString(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("## Transfered html from js to go:\n", vm.Get("doc").ToString())
}

func gojaRun(fname string) {
	runtime := goja.New()
	registry := new(require.Registry)
	registry.Enable(runtime)
	registry.RegisterNativeModule("console", console.Require)
	console.Enable(runtime)

	c := runtime.Get("console")
	if c == nil {
		fmt.Println("console is nil")
	}

	loadToGoja(runtime, fname)
}
