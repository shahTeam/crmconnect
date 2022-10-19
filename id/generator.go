package id

import (
	"github.com/google/uuid"
)

type IGenerator interface {
	GenerateUUID() uuid.UUID
}

type Generater struct{}

func (g Generater) GenerateUUID() uuid.UUID {
	return uuid.New()
}