package main

import(
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){ //r is pointing to the request
if r.URL.Path != "/hello"{
	http.Error(w,"404 not found",http.StatusNotFound)
	return
}
if r.Method != "GET"{
	http.Error(w,"method not supported",http.StatusNotFound)
	return
}
fmt.Fprintf(w,"hello! from the server")
}


func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w, "POST request successful\n")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
}


func main(){
	// used shorthand operator for declaring and defining a variable
	fileServer :=http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)


	fmt.Printf("Starting server at port 8080")
	if err := http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal(err)
	}
}