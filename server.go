package boxcars

import (
	"fmt"
	"net/http"
)

func Listen(port int) {
	debug("Starting at %d", port)
	http.HandleFunc("/", OnRequest)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
