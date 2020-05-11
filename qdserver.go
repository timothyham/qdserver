package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var VersionString = ""

var help = flag.Bool("h", false, "Show help page")
var version = flag.Bool("v", false, "Show Version")
var dir = flag.String("d", ".", "directory to server")
var port = flag.Int("p", 8080, "listening port")

func main() {
	flag.Parse()
	if *help {
		printHelp()
		return
	}
	if *version {
		printVersion()
		return
	}

	portStr := ":" + strconv.Itoa(*port)
	fmt.Printf("QDServer version %s\nServing %s on %s\n", VersionString, *dir, portStr)

	http.Handle("/", NewLoggingHandler(*dir))
	err := http.ListenAndServe(portStr, nil)
	if err != nil {
		fmt.Printf("Error '%v'\n", err)
	}
}

type loggingHandler struct {
	dir     string
	handler http.Handler
	logger  *log.Logger
}

func NewLoggingHandler(dir string) *loggingHandler {
	h := &loggingHandler{
		dir:     dir,
		handler: http.FileServer(http.Dir(dir)),
		logger:  log.New(os.Stdout, "QDSERVER:", log.Ldate|log.Ltime)}
	return h
}

func (l *loggingHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	l.logger.Printf("%s:%s", r.RemoteAddr, r.RequestURI)
	l.handler.ServeHTTP(rw, r)
}

func printHelp() {
	helptext := `
qdserver - quick and dirty static file server

qdserver [options] 

options:
  -d Directory to serve. Default is "."
  -p Port to listen on
  -v Print version string
`
	fmt.Printf(helptext)
}

func printVersion() {
	fmt.Printf(VersionString + "\n")
}
