package quack

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/jessemillar/jsonresp"
)

func Quack(writer http.ResponseWriter, request *http.Request) {
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		jsonresp.New(writer, http.StatusOK, "QUACK! There was an error during request parsing.")
		return
	}

	jsonresp.New(writer, http.StatusOK, "QUACK QUACK QUACK QUACK QUACK")

	log.SetOutput(os.Stdout)
	log.Printf("%s\n\n", dump)
}
