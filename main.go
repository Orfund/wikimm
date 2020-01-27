package main
import (
"fmt"
"net/http"
)
func main() {
	fmt.Println("Server is listening...")
	http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r* http.Request){
		http.ServeFile(w, r, "index.html")
	})
	http.ListenAndServe(":80", nil)

}
