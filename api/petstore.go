//go:generate go tool oapi-codegen --config=types.cfg.yaml ../schema/openapi.yaml
//go:generate go tool oapi-codegen --config=server.cfg.yaml ../schema/openapi.yaml

package api

import (
	"context"
	"fmt"
)

type PetStore struct {
}

// Make sure we conform to StrictServerInterface
var _ StrictServerInterface = (*PetStore)(nil)

// NewPetStore creates a new PetStore repository.
func NewPetStore() *PetStore {
	return &PetStore{}
}

// FindPets implements all the handlers in the ServerInterface.
func (p *PetStore) FindPets(ctx context.Context, request FindPetsRequestObject) (FindPetsResponseObject, error) {
	return nil, fmt.Errorf("not implemented")
}

func (p *PetStore) FindPetByID(ctx context.Context, request FindPetByIDRequestObject) (FindPetByIDResponseObject, error) {
	return nil, fmt.Errorf("not implemented")
}

func (p *PetStore) AddPet(ctx context.Context, request AddPetRequestObject) (AddPetResponseObject, error) {
	return nil, fmt.Errorf("not implemented")
}

func (p *PetStore) DeletePet(ctx context.Context, request DeletePetRequestObject) (DeletePetResponseObject, error) {
	return nil, fmt.Errorf("not implemented")
}
