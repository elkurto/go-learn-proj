# golang static content service 
- there's no caching (out of the box)
  - so use nginx instead 
  - or create a caching mechanism

### This project exemplifies usage of golang net/http to serve static content

### usage:
cd http_aaa
go run .

visit http://localhost:8080/assets/index.html

[screenshot](./doc/20250829.page.png)


### project's file structure
```
 % tree http_aaa
    http_aaa
    ├── go.mod
    ├── http_serve_static.go
    ├── public
    │   ├── index.css
    │   └── index.html
    └── README.md

```