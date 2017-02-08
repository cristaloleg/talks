package code

import (
	"fmt"
	"net/http"
)

const LOCALHOST = "127.0.0.1:2990"

func main() {
	hello := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi, ApplicationUser \\ʕ◔ϖ◔ʔ/")
	}
	http.ListenAndServe(LOCALHOST, http.HandlerFunc(hello))
}
