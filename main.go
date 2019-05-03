package main

import (
	"TDKTServer/driver"
	"TDKTServer/routes"
	"TDKTServer/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func main() {

	port := ":3000"
	driver.StartPosgresDbConnection()

	fmt.Printf("TDKT GoLang Server Api's running on port %s\n", port)

	go utils.Manager.Start()
	r := routes.NewRouter()
	r.HandleFunc("/ws", WsPage)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
	handler := c.Handler(r)
	srv := &http.Server{
		Handler: handler,
		Addr:    port,
	}
	log.Fatal(srv.ListenAndServe())
}
func WsPage(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	client := &utils.Client{Id: uuid.NewV4().String(), Socket: conn, Send: make(chan []byte)}
	utils.Manager.Register <- client
	go client.Read()
	go client.Write()
}
