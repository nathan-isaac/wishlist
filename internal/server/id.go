package server

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const idChars = "0123456789abcdefghijklmnopqrstuvwxyz"

func GenerateId() (string, error) {
	return gonanoid.Generate(idChars, 10)
}
