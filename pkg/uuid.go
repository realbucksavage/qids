package pkg

import "github.com/google/uuid"

// NewV1UUID generates a random UUIDv1
func NewV1UUID() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}

// NewV1UUID generates a random UUIDv4
func NewV4UUID() (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}
