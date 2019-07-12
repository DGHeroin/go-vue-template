package main

import (
    "github.com/gobuffalo/packr"
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    distBox := packr.NewBox("frontend/dist")
    mux.Handle("/", http.FileServer(distBox))

    staticBox := packr.NewBox("frontend/static")
    mux.Handle("/static", http.FileServer(staticBox))

    srv := &http.Server{
        Addr:":80",
        Handler: mux,
    }

    if err := srv.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
