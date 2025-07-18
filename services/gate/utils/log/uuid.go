package log

import "github.com/gofrs/uuid"

func NewUuid() UUIDGenerator {
	return &UUID{}
}

type UUIDGenerator interface {
	Generate() (string, error)
}

type UUID struct{}

func (UUID) Generate() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
