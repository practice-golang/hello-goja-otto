const babel = require('@babel/standalone');

// run_linkedom.mjs
const code = `$$_ESM_SCRIPT_$$`;

const out = babel.transform(code, { presets: ["es2015"] });
console.log(out.code)

var transpiled = out.code