package main

import (
    "fmt"
    "net/http"
    "github.com/a-h/templ"
    pages "photoUploader/Pages"
)

func main() {
    component := pages.Index()
    http.Handle("/", templ.Handler(component))

    localPage := pages.LocalPage()
    http.Handle("/local", templ.Handler(localPage))

    downloadPage := pages.DownloadPage()
    http.Handle("/download", templ.Handler(downloadPage))

    uploadPage := pages.UploadPage()
    http.Handle("/upload", templ.Handler(uploadPage))

    fmt.Println("listening to http://localhost:3000")
    http.ListenAndServe(":3000", nil)

}
