package main
import (
	"fmt"
	"net/http"
)
func main() {
	handler := http.NewServeMux()
	fmt.Println("App is running...")
	handler.HandleFunc("/api/hello", SayHello)
	http.ListenAndServe("0.0.0.0:8080", handler)
}
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `Hello world `)
	fmt.Println("Hit")
}
