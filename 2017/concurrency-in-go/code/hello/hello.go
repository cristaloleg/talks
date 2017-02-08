package code

import (
	"fmt"
	"net/http"
)

const localhost = "127.0.0.1:2990"

func main() {
	hello := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi, ApplicationUser \\ʕ◔ϖ◔ʔ/")
	}
	http.ListenAndServe(localhost, http.HandlerFunc(hello))
}
