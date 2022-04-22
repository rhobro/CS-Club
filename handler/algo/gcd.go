package algo

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/rhobro/csclub/db"
)

// generate gcd problem
func giveGCD(w http.ResponseWriter, r *http.Request) {
	// gen
	n1 := rand.Uint64()
	n2 := rand.Uint64()
	problem := gcdProblem{
		N1: n1,
		N2: n2,
	}

	// store
	id, err := db.AddEntry(algoGcd, problem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// marshal
	body, err := json.MarshalIndent(problem, "", "    ") // we like it tidy :)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send
	w.Header().Set("entry", id)
	_, _ = w.Write(body)
}

// check answer
func receiveGCD(w http.ResponseWriter, r *http.Request) {

}

// PROBLEM MANAGEMENT

type gcdProblem struct {
	N1 uint64 `json:"n1"`
	N2 uint64 `json:"n2"`
}

const algoGcd = "algo-gcd"

func init() {
	db.Register(algoGcd)
}
