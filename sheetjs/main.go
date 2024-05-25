/* SheetGoja (C) SheetJS LLC -- https://sheetjs.com */
// https://docs.sheetjs.com/docs/demos/engines/goja/
package main // import "hello-goja-sheetjs"

import (
	b64 "encoding/base64"
	"fmt"
	"os"

	"io/fs"

	"github.com/dop251/goja"
)

func safe_run_file(vm *goja.Runtime, file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	_, err = vm.RunString(string(data))
	if err != nil {
		panic(err)
	}
}

func eval_string(vm *goja.Runtime, cmd string) goja.Value {
	v, err := vm.RunString(cmd)
	if err != nil {
		panic(err)
	}
	return v
}

func write_type(vm *goja.Runtime, t string) {
	b64str := eval_string(vm, "XLSX.write(wb, {type:'base64', bookType:'"+t+"'})")
	buf, err := b64.StdEncoding.DecodeString(b64str.String())
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("sheetjsw."+t, buf, fs.FileMode(0644))
	if err != nil {
		panic(err)
	}
}

func main() {
	/* initialize */
	vm := goja.New()
	eval_string(vm, "var global = (function(){ return this; }).call(null);")

	/* load library */
	safe_run_file(vm, "shim.min.js")
	safe_run_file(vm, "xlsx.full.min.js")

	/* get version string */
	v := eval_string(vm, "XLSX.version")
	fmt.Printf("SheetJS library version %s\n", v)

	/* read file */
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	vm.Set("buf", vm.ToValue(vm.NewArrayBuffer(data)))

	/* parse workbook */
	eval_string(vm, "wb = XLSX.read(buf, {type:'buffer'});")
	eval_string(vm, "ws = wb.Sheets[wb.SheetNames[0]]")

	/* print CSV */
	csv := eval_string(vm, "XLSX.utils.sheet_to_csv(ws)")
	fmt.Printf("%s\n", csv)

	/* write file */
	write_type(vm, "xlsb")
}
