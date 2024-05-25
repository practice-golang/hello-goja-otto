package main // import hello-otto-goja

import "fmt"

func main() {
	fmt.Println("## Otto / jsdom ----")
	ottoRun("./run_jsdom.js")

	fmt.Println("## Otto / linkedom ----")
	ottoRun("./run_linkedom.js")

	fmt.Println("## Goja / jsdom ----")
	gojaRun("./run_jsdom.js")

	fmt.Println("## Goja / linkedom ----")
	gojaRun("./run_linkedom.js")
}
