package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getDoc)
	http.ListenAndServe(":8081", mux)
}



func getDoc(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "https://go.dev/doc/", http.StatusSeeOther)
}