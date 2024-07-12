package main


import (
    "fmt"
    "net/http"
)


func routeHandler(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
        case "/":
            fmt.Fprint(w, "Hej, Verden!")
            break
        case "/henk":
            fmt.Fprint(w, "Yo mama, my name is Henk de Henker");
            break
        default:
            fmt.Fprint(w, "Page not found!");
    }
}


func main() {
    http.HandleFunc("/", routeHandler)
    http.ListenAndServe(":8080", nil)
}
