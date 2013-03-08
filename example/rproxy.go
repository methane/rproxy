package main

import (
    "flag"
    "fmt"
    "github.com/methane/rproxy"
    "net/http"
    "net/url"
)

func main() {
    var port *int = flag.Int("port", 8000, "Listen port")
    var backend *string = flag.String("backend", "http://127.0.0.1:9000", "Backend host")

    flag.Parse()
    backendUrl, _ := url.Parse(*backend)
    proxy := rproxy.NewSingleHostReverseProxy(backendUrl)
    http.Handle("/", proxy)
    listenHost := fmt.Sprintf(":%v", *port)
    fmt.Println("Listen on ", listenHost)
    fmt.Println("Forward to ", *backend)
    http.ListenAndServe(listenHost, nil)
}
