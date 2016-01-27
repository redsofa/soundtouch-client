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

const (
	VERSION = "0.0.1"
)

func main() {
	//Setup the logger
	logger.InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	//Read configuration information from config.json
	config.ReadServerConf("./")

	//The port our server listens on
	listenPort := config.ServerConf.Port

	logger.Info.Printf("Sever Starting - Listing on port %d - (Version - %s)", listenPort, VERSION)

	//Create router
	router := mux.NewRouter()

	//Routes
	router.Handle("/", http.FileServer(http.Dir("../static")))
	router.Handle("/rest/presets", &handlers.Presets{}).Methods("GET")
	router.Handle("/rest/select/{source}/{location}", &handlers.Select{}).Methods("GET")
	router.Handle("/rest/volume", &handlers.GetVolume{}).Methods("GET")
	router.Handle("/rest/volume/{target}", &handlers.SetVolume{}).Methods("GET")

	//Listen for connections and serve content
	logger.Info.Println(http.ListenAndServe(":"+strconv.Itoa(listenPort), logger.HttpLog(router)))
}
