# go + webassembly basic example

#### References:
- derived from https://go.dev/wiki/WebAssembly#javascript-goosjs-port
- build a calc in go tutorial : https://youtu.be/4kBvvk2Bzis?si=CxT1GoSyITkSReUt&t=130

```bash
- prerequistie: go compiler (1.24.2) installed locally in dir, ${GOROOT}
- mkdir 00-00-go-wasm
- cd 00-00-go-wasm/
- cp ${GOROOT}/misc/wasm/wasm_exec.html .
- edit ./wasm_exec.html  to enable the "Run" button (or just use wasm_exec.html from this repo)
- cp ${GOROOT}/lib/wasm/wasm_exec.js .
- go mod init 00-00-go-wasm
- touch main.go
- add the following context to main.go
    package main
    
    import (
      "fmt"
    )
    
    func main() {
      fmt.Println("Hello Wasm")
    }
- create main.wasm via the following command;
  GOOS=js GOARCH=wasm go build -o main.wasm
- start simple [python] server to show the content
  python3 -m http.server 8888
- in a browser open web-page http://localhost:8888/wasm_exec.html
  - open the debug console (e.g. key F12 in chrome)
  - click the button, Run. (see image below for anticapted result)
```
[./support/wasm_exec.html.png](./support/wasm_exec.html.png)


```bash
tree 00-00-go-wasm/
00-00-go-wasm/
├── go.mod
├── main.go
├── main.wasm
├── README.md
├── wasm_exec.html
└── wasm_exec.js



```