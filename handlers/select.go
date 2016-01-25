package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/redsofa/soundtouch-client/config"
	"github.com/redsofa/soundtouch-client/logger"
	"io/ioutil"
	"net/http"
	"strings"
)

type Select struct {
}

func (h *Select) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("In Select Handler ... ")

	vars := mux.Vars(r)

	source, ok := vars["source"]
	if !ok {
		msg := "source variable not set"
		logger.Error.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	location, ok := vars["location"]
	if !ok {
		msg := "location variable not set"
		logger.Error.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	logger.Info.Println("Source : ", source)
	logger.Info.Println("Location : ", location)

	data, err := h.makeSelection(source, location)

	if err != nil {
		msg := "Internal server error making select. Check Logs."
		logger.Error.Println(err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	//Create response with our results...
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, *data)
}

func (h *Select) makeSelection(source string, location string) (*string, error) {
	url := config.ServerConf.SoundTouchUrl + "/select"

	xml := fmt.Sprintf("<ContentItem source=\"%s\" sourceAccount=\"\" location=\"%s\"><itemName></itemName></ContentItem>", source, location)

	req, err := http.NewRequest("POST", url, strings.NewReader(xml))
	req.Header.Set("Content-Type", "text/xml")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	retVal := string(body)

	return &retVal, nil

}
