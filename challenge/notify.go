package challenge

import (
	"fmt"
	"net/http"
)

func Process(r *http.Request, challenge string, correct bool) {
	fmt.Printf("%s -> %t\n", challenge, correct)
}
