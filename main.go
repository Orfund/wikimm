package main
import (
"fmt"
"net/http"
"net"
)
func main() {

    c, _ := net.Dial("unix", "/Users/jernicozz/Documents/conn")
   	defer c.Close()
   	msg := "hi"
   	_, _ = c.Write([]byte(msg))
    println("Client sent:", msg)
    buf := make([]byte, 1024)
    for {
  		n, err := c.Read(buf[:])
   		if err != nil {
   			return
   	    }
   	    println("Client got:", string(buf[0:n]))
    }
	fmt.Println("Server is listening...")
	http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("static/"))))
	http.Handle("/data/", http.StripPrefix("/data/",http.FileServer(http.Dir("data/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r* http.Request){
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/person", func(w http.ResponseWriter, r* http.Request){

	})
	http.ListenAndServe(":80", nil)

}
