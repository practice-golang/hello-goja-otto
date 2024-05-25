package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

func GetExtensions(filename string) (nameWithoutExt, fullExt string) {
	ext := filepath.Ext(filename)
	if ext == "" {
		return filename, ""
	}

	nameWithoutExt = strings.TrimSuffix(filename, ext)

	secondExt := filepath.Ext(nameWithoutExt)

	if secondExt != "" {
		fullExt = secondExt + ext
		nameWithoutExt = strings.TrimSuffix(nameWithoutExt, secondExt)
	} else {
		fullExt = ext
	}

	if fullExt != "" {
		fullExt = fullExt[1:]
	}

	return nameWithoutExt, fullExt
}

func loadToGoja(vm *goja.Runtime, file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	tcodeBytes, _ := os.ReadFile("./transpiler.js")
	_, fext := GetExtensions(file)
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
