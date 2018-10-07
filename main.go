
package main

import (
    "fmt"
    "net/http"
)

func root(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w,"Hello World")

}

func main(){
    fmt.Printf("hello world\n");
    http.HandleFunc("/",root)
    http.ListenAndServe(":8000",nil)
}

