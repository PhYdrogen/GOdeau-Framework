package Controller

import (
	"GOdeau/Module/Controllers"
	"fmt"
	"net/http"
)

func Init() {

	Controllers.InitController(
		"HelloWorld",
		Controllers.Handler(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World")
		}))

}
