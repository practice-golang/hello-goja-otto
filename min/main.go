package main // import "hello-goja-min"

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

func main() {
	vm := goja.New()

	registry := new(require.Registry)
	registry.Enable(vm)
	registry.RegisterNativeModule("console", console.Require)
	console.Enable(vm)

	v, err := vm.RunString("2 + 2")
	if err != nil {
		panic(err)
	}

	num := v.Export().(int64)
	if num != 4 {
		panic(num)
	}

	fmt.Println("Hello, Goja!")
	fmt.Println("2 + 2 =", num)

	c := vm.Get("console")
	if c == nil {
		panic("console is nil")
	}

	_, err = vm.RunString("console.log('console.log is alive!!')")
	if err != nil {
		panic(err)
	}

	_, err = vm.RunString(`
	const Class = class extends HTMLElement {
		constructor() {
			super(Component, slots, use_shadow_dom);
			this.$$props_definition = props_definition;
		}
		static get observedAttributes() {
			return Object.keys(props_definition).map((key) =>
				(props_definition[key].attribute || key).toLowerCase()
			);
		}
	};`)
	if err != nil {
		panic(err)
	}
}
