package main

import (
	"fmt"
	"os"

	"github.com/robertkrimen/otto"
)

func loadToOtto(vm *otto.Otto, file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	_, err = vm.Run(string(data))
	if err != nil {
		fmt.Println(err)
	}
}

func ottoRun(fname string) {
	vm := otto.New()
	loadToOtto(vm, fname)
}
