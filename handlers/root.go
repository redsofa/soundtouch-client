package handlers

import (
	"fmt"
	"github.com/redsofa/soundtouch-client/logger"
	"net/http"
)

type Root struct {
}

func (h *Root) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("In Root Handler ... ")

	//Create response with our results...
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "<html><body>TODO...</body></html>")
}
