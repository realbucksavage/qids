package pkg

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// NewULID generates a random ulid.
func NewULID() (string, error) {
	t := time.Now()
	es := rand.New(rand.NewSource(t.UnixNano()))
	entropy := ulid.Monotonic(es, 0)

	id, err := ulid.New(ulid.Timestamp(t), entropy)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
