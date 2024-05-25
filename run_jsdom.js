const jsdom = require("jsdom");
const { JSDOM } = jsdom;

const dom = new JSDOM(`
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

console.log(dom.window.document.querySelector("#main").textContent);
