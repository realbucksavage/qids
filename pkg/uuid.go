package pkg

import "github.com/google/uuid"

func NewV1UUID() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}

func NewV4UUID() (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}
