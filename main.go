package main // import hello-otto-goja-sobek

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

	fmt.Println("## Goja / linkedom mjs w/ Babel ----")
	gojaRun("./run_linkedom.mjs")

	fmt.Println("## Sobek / jsdom mjs ----")
	sobekRun("./run_jsdom.mjs")

	fmt.Println("## Sobek / linkedom mjs ----")
	sobekRun("./run_linkedom.mjs")

	fmt.Println("## Sobek / linkedom mjs w/ Babel ----")
	sobekRun("./run_linkedom.mjs")
}
