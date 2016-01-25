package logger

import (
	"io"
	"log"
	"net/http"
	"time"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

/*
 * Basic Wrapper for handlers which logs all requests.
 * Handlers get passed to this function which then
 * call the serveHTTP method on the handler and logs
 * the call
 * Source :
 * - http://thenewstack.io/make-a-restful-json-api-go/
 * - https://groups.google.com/forum/#!topic/golang-nuts/s7Xk1q0LSU0
 *
 * Usage Example :
 *
 * func main() {
 *	 router := mux.NewRouter()
 *
 *   router.Handle("/api/someresource", &handlers.SomeHandler{})
 *   http.ListenAndServe(":3000",  logger.HttpLog(router))
 * }
 *
 * Notice the logger.HttpLog(router) in the last line of the code above
 */

func HttpLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		Info.Printf(
			"Method : %s\t URI: %s\t URL: %s\t Time: %s",
			r.Method,
			r.RequestURI,
			r.URL,
			time.Since(start),
		)

		handler.ServeHTTP(w, r)

	})
}

/*
 * Source : http://www.goinggo.net/2013/11/using-log-package-in-go.html
 *
 * Usage Example :
 *
 * func main() {
 *   logger.InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
 *	 logger.Info.Printf("Some Info ...")
 *   logger.Error.Println("Something has failed")
 *   //Warning, Trace etc...
 * }
 *
 */
func InitLogger(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
