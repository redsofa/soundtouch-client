package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/redsofa/soundtouch-client/config"
	"github.com/redsofa/soundtouch-client/domain"
	"github.com/redsofa/soundtouch-client/logger"
	"io/ioutil"
	"net/http"
)

type GetVolume struct {
}

func (h *GetVolume) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("In GetVolume Handler ... ")

	data, err := h.fetchVolume()

	if err != nil {
		msg := "Internal server error fetching Volume. Check Logs."
		logger.Error.Println(err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	//Create response with our results...
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, *data)
}

func (h *GetVolume) fetchVolume() (*string, error) {
	url := config.ServerConf.SoundTouchUrl + "/volume"

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	x := domain.Volume{}

	err = xml.Unmarshal([]byte(body), &x)
	if err != nil {
		return nil, err
	}

	j, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	retVal := string(j)

	return &retVal, nil
}
