package algo

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"

	"github.com/rhobro/csclub/challenge"
	"github.com/rhobro/csclub/db"
)

// generate gcd problem
func giveGCD(w http.ResponseWriter, r *http.Request) {
	// gen
	n1 := rand.Int63()
	n2 := rand.Int63()
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
	// values
	id := r.Header.Get("entry")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no entry header provided"))
		return
	}

	// read body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("no body provided"))
		return
	}
	type ans struct {
		Ans int64 `json:"ans"`
	}
	attempt := ans{}
	err = json.Unmarshal(body, &attempt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request body"))
		return
	}

	// retrive case
	q, err := db.GetEntry(algoGcd, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid entry id"))
		return
	}
	question := q.(gcdProblem)

	// check
	var correct bool
	if gcd(question.N1, question.N2) == attempt.Ans {
		correct = true
	} else {
		correct = false
	}

	// delete entry
	if correct {
		db.RemoveEntry(algoGcd, id)
	}

	// if correct send 200
	if correct {
		w.Write([]byte("correct"))
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("incorrect"))
	}

	challenge.Process(r, algoGcd, correct)
}

func gcd(a, b int64) int64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

// PROBLEM MANAGEMENT

type gcdProblem struct {
	N1 int64 `json:"n1"`
	N2 int64 `json:"n2"`
}

const algoGcd = "algo-gcd"

func init() {
	db.Register(algoGcd)
}
