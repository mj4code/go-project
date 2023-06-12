package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"net/http"
	"log"
)

func main() {
	godotenv.Load(".env")
	port:=os.Getenv("PORT")

	if port==""{
		fmt.Println("Port is not defined");
	}else{
		fmt.Println("Port :",port)
	}

	router:=chi.NewRouter()

	server:=&http.Server{
		Handler :router,
		Addr : ":"+port,
	}

	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}


}
