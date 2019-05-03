package hello

import (
	"fmt"
	"net/http"
)

func HelloGopher(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
