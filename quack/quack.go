package quack

import (
	"log"
	"net/http"

	"github.com/jessemillar/jsonresp"
)

func Quack(writer http.ResponseWriter, request *http.Request) {
	jsonresp.New(writer, http.StatusOK, "QUACK QUACK QUACK QUACK QUACK")
	log.Printf("%s\n", request)
}
