package main

import (
	"github.com/ma-portal/core/controller"
	"net/http"
)

func main() {
	_ = http.ListenAndServe(":9080", controller.Router)
}
