package helper

import (
	"math/rand"
	"sync"
	"time"
)

var (
	seededIDGen = rand.New(rand.NewSource(time.Now().UnixNano()))
	// The golang rand generators are *not* intrinsically thread-safe.
	seededIDLock sync.Mutex
)

// RandomID ...
func RandomID() uint64 {
	seededIDLock.Lock()
	defer seededIDLock.Unlock()
	return uint64(seededIDGen.Int63())
}

