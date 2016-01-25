package main

import (
	"github.com/gorilla/mux"
	"github.com/redsofa/soundtouch-client/config"
	"github.com/redsofa/soundtouch-client/handlers"
	"github.com/redsofa/soundtouch-client/logger"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//Setup the logger
	logger.InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	//Read configuration information from server_config.json
	config.ReadServerConf("./")

	//The port our server listens on
	listenPort := config.ServerConf.Port

	logger.Info.Printf("Sever Starting - Listing on port %d - (Version - %s)", listenPort, "0.0.1")

	//Create router
	router := mux.NewRouter()

	//Setup our routes
	router.Handle("/presets", &handlers.Presets{}).Methods("GET")
	router.Handle("/select/{source}/{location}", &handlers.Select{}).Methods("GET")

	//Listen for connections and serve content
	logger.Info.Println(http.ListenAndServe(":"+strconv.Itoa(listenPort), logger.HttpLog(router)))
}
