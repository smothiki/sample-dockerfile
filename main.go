package main

import (
	"fmt"
	"net/http"
	"os"
)

var src = `<a href="https://imgflip.com/i/7zmo16"><img src="https://i.imgflip.com/7zmo16.jpg" title="made at imgflip.com"/></a><div><a href="https://imgflip.com/memegenerator">from Imgflip Meme Generator</a></div>`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
		fmt.Fprintf(w, "<title>Go</title>")
		fmt.Fprintf(w, src)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "7860"
	}

	fmt.Printf("Server listening at :%s 🚀\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}
