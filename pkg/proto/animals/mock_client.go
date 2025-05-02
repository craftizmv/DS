package animals

import (
	"context"
	"errors"
)

// MockAnimalsClient provides a mock implementation of AnimalsClient for testing
type MockAnimalsClient struct {
	animals map[string]*Animal
}

// NewMockAnimalsClient creates a new mock client
func NewMockAnimalsClient() *MockAnimalsClient {
	return &MockAnimalsClient{
		animals: make(map[string]*Animal),
	}
}

// GetAnimals retrieves all animals
func (m *MockAnimalsClient) GetAnimals(ctx context.Context, req *GetAnimalsRequest) (*GetAnimalsResponse, error) {
	result := make([]*Animal, 0, len(m.animals))
	for _, animal := range m.animals {
		result = append(result, animal)
	}
	return &GetAnimalsResponse{Animals: result}, nil
}

// DeleteAnimals deletes all animals
func (m *MockAnimalsClient) DeleteAnimals(ctx context.Context, req *DeleteAnimalsRequest) (*DeleteAnimalsResponse, error) {
	m.animals = make(map[string]*Animal)
	return &DeleteAnimalsResponse{}, nil
}

// AddAnimal adds a new animal
func (m *MockAnimalsClient) AddAnimal(ctx context.Context, req *AddAnimalRequest) (*AddAnimalResponse, error) {
	if req.Animal == nil {
		return nil, errors.New("animal cannot be nil")
	}
	
	if req.Animal.Name == "" {
		return nil, errors.New("animal name cannot be empty")
	}
	
	animal := &Animal{
		Name:   req.Animal.Name,
		Color:  req.Animal.Color,
		Height: req.Animal.Height,
	}
	
	m.animals[animal.Name] = animal
	return &AddAnimalResponse{Animal: animal}, nil
}

// GetAnimal gets an animal by name
func (m *MockAnimalsClient) GetAnimal(ctx context.Context, req *GetAnimalRequest) (*GetAnimalResponse, error) {
	animal, exists := m.animals[req.Name]
	if !exists {
		return nil, errors.New("animal not found")
	}
	return &GetAnimalResponse{Animal: animal}, nil
}
