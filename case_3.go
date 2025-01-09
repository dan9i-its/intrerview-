package main

import (
    "fmt"
    "path"
    "net/http"
    "io"
    "os"
)

func payment_check(w http.ResponseWriter, r *http.Request) {
    product := r.Form.Get("product")
    if !isEnoughMoneyOnBalance(product) {
	    http.Error(w, "Not enought money.", 403)
    }
    send_order(product)
    
}

func isEnoughMoneyOnBalance(product string) bool {
  ... // secure check
}

func send_order(product string) {
  ...
}

func main() {
    http.HandleFunc("/", payment_check)
    http.ListenAndServe(":8080", nil)
}
