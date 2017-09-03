package main

import (
	// Router
	"github.com/krishamoud/game-server/app/common/router"

	// Common code
	"github.com/krishamoud/game-server/app/common/db"
	_ "github.com/krishamoud/game-server/app/common/docker"

	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	// close the db connection when we're done
	defer db.MongoConn.Close()

	// Handle all requests with gorilla/mux
	http.Handle("/", router.Router())

	// Listen on port 9090
	log.Println("Server listening on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
