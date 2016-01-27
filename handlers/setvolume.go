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

type SetVolume struct {
}

func (h *SetVolume) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("In SetVolume Handler ... ")

	vars := mux.Vars(r)

	target, ok := vars["target"]
	if !ok {
		msg := "SetVolume variable not set"
		logger.Error.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	logger.Info.Println("Target : ", target)

	_, err := h.setVolume(target)

	if err != nil {
		msg := "Internal server error changing volume. Check Logs."
		logger.Error.Println(err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	//Create response with our results...
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"Status\":\"OK\", \"TargetVolume\":%s}", target)
}

func (h *SetVolume) setVolume(target string) (*string, error) {
	url := config.ServerConf.SoundTouchUrl + "/volume"

	xml := fmt.Sprintf("<volume>%s</volume>", target)

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
