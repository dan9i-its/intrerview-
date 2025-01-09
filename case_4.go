func isCsrfTokenValid(token string) bool {
  ... //some secure checks
}

func next_function(password string) {
  ... // set password
}

func change_password(w http.ResponseWriter, r *http.Request) {
  password := r.Form.Get("password")
  token := r.Form.Get("token")

  if token !=""  && !isCsrfTokenValid(token) {
    http.Error(w, "Invalid Token.", 403)
    return
  }
  next_function(password)
}

func main() {
    http.HandleFunc("/", change_password)
    http.ListenAndServe(":8080", nil)
}
