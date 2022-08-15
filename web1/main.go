package main
import (
	"fmt"
	"net/http"
)
func handlerFunc(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"<h1>Welcome to my website</h1>")
}
func main(){
	http.HandleFunc("/",handlerFunc)
	fmt.Println("Starting server at port :3000...")
	http.ListenAndServe(":3000",nil)
}
