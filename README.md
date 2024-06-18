Taste Go js runtimes.


## Taste working

|     |[Otto](https://github.com/robertkrimen/otto)|[Goja](https://github.com/dop251/goja)|[Sobek](https://github.com/grafana/sobek)|
|-----|----|----|----|
|[jsDOM](https://github.com/jsdom/jsdom)|No|No|No|
|[linkeDOM](https://github.com/WebReflection/linkedom)|No|Yes|Yes|

* Sobek is a fork of Goja for js module


## Run
```sh
yarn
go run .


## Otto / jsdom ----
(anonymous): Line 1:1 Unexpected reserved word (and 2 more errors)
## Otto / linkedom ----
(anonymous): Line 2:1 Unexpected reserved word (and 3 more errors)
## Goja / jsdom ----
GoError: Invalid module at github.com/dop251/goja_nodejs/require.(*RequireModule).require-fm (native)
## Goja / linkedom ----
2024/06/18 23:57:48 
        Hello
        World
    
#### Transfered html from js to go:
 <!DOCTYPE >

<html>
<body>
    <div id="main">
        <p>Hello</p>
        <p>World</p>
    </div>
</body>
</html>

## Goja / linkedom mjs w/ Babel ----
2024/06/18 23:57:49 "use strict";

var _linkedom = require("linkedom");
var _parseHTML = (0, _linkedom.parseHTML)("\n<!DOCTYPE html>\n<html>\n<body>\n    <div id=\"main\">\n        <p>Hello</p>\n        <p>World</p>\n    </div>\n</body>\n</html>\n"),
  window = _parseHTML.window,
  document = _parseHTML.document,
  customElements = _parseHTML.customElements,
  HTMLElement = _parseHTML.HTMLElement,
  Event = _parseHTML.Event,
  CustomEvent = _parseHTML.CustomEvent;
console.log(document.querySelector("#main").textContent);
var doc = document.toString();
2024/06/18 23:57:49 
        Hello
        World
    
#### Transfered html from js to go:
 <!DOCTYPE >

<html>
<body>
    <div id="main">
        <p>Hello</p>
        <p>World</p>
    </div>
</body>
</html>

## Sobek / jsdom mjs ----
2024/06/18 23:57:49 "use strict";

var _jsdom = require("jsdom");
var dom = new _jsdom.JSDOM("\n<!DOCTYPE html>\n<html>\n<body>\n    <div id=\"main\">\n        <p>Hello</p>\n        <p>World</p>\n    </div>\n</body>\n</html>\n");
console.log(dom.window.document.querySelector("#main").textContent);
GoError: Invalid module at github.com/practice-golang/goja_nodejs/require.(*RequireModule).require-fm (native)
## Sobek / linkedom mjs ----
2024/06/18 23:57:49 "use strict";

var _linkedom = require("linkedom");
var _parseHTML = (0, _linkedom.parseHTML)("\n<!DOCTYPE html>\n<html>\n<body>\n    <div id=\"main\">\n        <p>Hello</p>\n        <p>World</p>\n    </div>\n</body>\n</html>\n"),
  window = _parseHTML.window,
  document = _parseHTML.document,
  customElements = _parseHTML.customElements,
  HTMLElement = _parseHTML.HTMLElement,
  Event = _parseHTML.Event,
  CustomEvent = _parseHTML.CustomEvent;
console.log(document.querySelector("#main").textContent);
var doc = document.toString();
2024/06/18 23:57:49 
        Hello
        World

#### Transfered html from js to go:
 <!DOCTYPE >

<html>
<body>
    <div id="main">
        <p>Hello</p>
        <p>World</p>
    </div>
</body>
</html>

## Sobek / linkedom mjs w/ Babel ----
2024/06/18 23:57:50 "use strict";

var _linkedom = require("linkedom");
var _parseHTML = (0, _linkedom.parseHTML)("\n<!DOCTYPE html>\n<html>\n<body>\n    <div id=\"main\">\n        <p>Hello</p>\n        <p>World</p>\n    </div>\n</body>\n</html>\n"),
  window = _parseHTML.window,
  document = _parseHTML.document,
  customElements = _parseHTML.customElements,
  HTMLElement = _parseHTML.HTMLElement,
  Event = _parseHTML.Event,
  CustomEvent = _parseHTML.CustomEvent;
console.log(document.querySelector("#main").textContent);
var doc = document.toString();
2024/06/18 23:57:50 
        Hello
        World

#### Transfered html from js to go:
 <!DOCTYPE >

<html>
<body>
    <div id="main">
        <p>Hello</p>
        <p>World</p>
    </div>
</body>
</html>

```
