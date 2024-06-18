package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/practice-golang/goja_nodejs/console"
	"github.com/practice-golang/goja_nodejs/require"

	"github.com/grafana/sobek"
)

func loadToSobek(vm *sobek.Runtime, file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	tcodeBytes, _ := os.ReadFile("./transpiler.js")
	_, fext := GetFileExtensions(file)
	if fext == "mjs" {
		data = []byte(strings.ReplaceAll(string(data), "`", "\\`"))
		tp := strings.Replace(string(tcodeBytes), "$$_ESM_SCRIPT_$$", string(data), 1)
		vm.RunString(tp)

		data = []byte(vm.Get("transpiled").String())
	}

	_, err = vm.RunString(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("#### Transfered html from js to go:\n", vm.Get("doc").ToString())
}

func sobekRun(fname string) {
	runtime := sobek.New()
	registry := new(require.Registry)
	registry.Enable(runtime)
	registry.RegisterNativeModule("console", console.Require)
	console.Enable(runtime)

	c := runtime.Get("console")
	if c == nil {
		fmt.Println("console is nil")
	}

	loadToSobek(runtime, fname)
}
