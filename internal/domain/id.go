package domain

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/oklog/ulid/v2"
)

const idChars = "0123456789abcdefghijklmnopqrstuvwxyz"

func GenerateId() (string, error) {
	return ulid.Make().String(), nil
}

func GenerateShareId() (string, error) {
	return gonanoid.Generate(idChars, 10)
}
