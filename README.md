# Capital-Investment-NPV-IRR-Calculator

[demo](https://codymccarty.github.io/Capital-Investment-NPV-IRR-Calculator/)

1. Make sure the runtime shim is present (you already copied this)  
$`cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" docs/`

1. Build WASM (this must succeed; no errors about strconv)  
$`GOOS=js GOARCH=wasm go build -o docs/app.wasm ./cmd/wasm`

1. Serve + open  
$`python3 -m http.server 8080`   
Visit http://localhost:8080/docs/ and do a hard reload (Ctrl+F5 / Shift+Reload)






