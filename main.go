package main
import (
"fmt"
"net/http"
"net"
)
func main() {

    c, _ := net.Dial("unix", "conn")
   	//defer c.Close()


	fmt.Println("Server is listening...")
	http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("static/"))))
	http.Handle("/data/", http.StripPrefix("/data/",http.FileServer(http.Dir("data/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r* http.Request){
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/person", func(w http.ResponseWriter, r* http.Request){
        msg := r.URL.Query()["uid"][0]
   	    _, _ = c.Write([]byte(msg))
        println("Client sent:", msg)
		buf := make([]byte, 1000)
		n, err := c.Read(buf[:])
		if err != nil {
			return
		}
		ans := string(buf[0:n])
		println("client received: ", ans)
		_, _ = w.Write([]byte(ans))
	})
	http.ListenAndServe(":80", nil)

}
