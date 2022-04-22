package db

import (
	"fmt"
	"strconv"
	"sync"
)

// record of problems
var storage = map[string]*problem{}
var storageMtx = &sync.RWMutex{}

// register a new problem
// create a record for it
func Register(name string) {
	storageMtx.Lock()
	defer storageMtx.Unlock()

	storage[name] = &problem{
		kv:  map[uint64]interface{}{},
		mtx: sync.RWMutex{},
	}
}

// problem manages the entries and results

type problem struct {
	kv   map[uint64]interface{} // maps ids to solution structs
	mtx  sync.RWMutex
	next uint64
}

func GetEntry(name string, idHex string) (interface{}, error) {

	// retrieve problem
	storageMtx.RLock()
	problem, ok := storage[name]
	storageMtx.RUnlock()
	if !ok {
		return nil, fmt.Errorf("invalid problem name: %s", name)
	}

	// parse hex id
	id, err := strconv.ParseUint(idHex, 16, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid hex id: %s", idHex)
	}

	// retrieve struct from entry
	problem.mtx.RLock()
	struc, ok := problem.kv[id]
	problem.mtx.RUnlock()
	if !ok {
		return nil, fmt.Errorf("invalid entry id: %d", id)
	}

	return struc, nil
}

func AddEntry(name string, struc interface{}) (string, error) {

	// retrieve problem
	storageMtx.RLock()
	problem, ok := storage[name]
	storageMtx.RUnlock()
	if !ok {
		return "", fmt.Errorf("invalid problem name: %s", name)
	}

	problem.mtx.Lock()

	// get next id
	id := fmt.Sprintf("%x", problem.next)
	problem.next += 1 // incr

	// set struct
	problem.kv[problem.next-1] = struc

	problem.mtx.Unlock()

	return id, nil
}
