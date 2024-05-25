// import { DOMParser, parseHTML } from 'linkedom';
const ld = require("linkedom");
const { DOMParser, parseHTML } = ld

const {
    window, document, customElements,
    HTMLElement,
    Event, CustomEvent
} = parseHTML(`
<!DOCTYPE html>
<html>
<body>
    <div id="main">
        <p>Hello</p>
        <p>World</p>
    </div>
</body>
</html>
`);

console.log(document.querySelector("#main").textContent);
var doc = document.toString();
