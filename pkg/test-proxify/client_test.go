package test_proxify

import (
	"context"
	"errors"
	"testing"

	pb "github.com/craftizmv/DS/pkg/proto/animals"
	"google.golang.org/grpc"
)

// MockAnimalsClient is a test implementation of the AnimalsClient interface
type mockAnimalsClient struct {
	animals   []*pb.Animal
	getError  error
	addError  error
	delError  error
	findError error
}

func (m *mockAnimalsClient) GetAnimals(ctx context.Context, req *pb.GetAnimalsRequest, opts ...grpc.CallOption) (*pb.GetAnimalsResponse, error) {
	if m.getError != nil {
		return nil, m.getError
	}
	return &pb.GetAnimalsResponse{Animals: m.animals}, nil
}

func (m *mockAnimalsClient) DeleteAnimals(ctx context.Context, req *pb.DeleteAnimalsRequest, opts ...grpc.CallOption) (*pb.DeleteAnimalsResponse, error) {
	if m.delError != nil {
		return nil, m.delError
	}
	m.animals = []*pb.Animal{}
	return &pb.DeleteAnimalsResponse{}, nil
}

func (m *mockAnimalsClient) AddAnimal(ctx context.Context, req *pb.AddAnimalRequest, opts ...grpc.CallOption) (*pb.AddAnimalResponse, error) {
	if m.addError != nil {
		return nil, m.addError
	}
	m.animals = append(m.animals, req.Animal)
	return &pb.AddAnimalResponse{Animal: req.Animal}, nil
}

func (m *mockAnimalsClient) GetAnimal(ctx context.Context, req *pb.GetAnimalRequest, opts ...grpc.CallOption) (*pb.GetAnimalResponse, error) {
	if m.findError != nil {
		return nil, m.findError
	}
	for _, animal := range m.animals {
		if animal.Name == req.Name {
			return &pb.GetAnimalResponse{Animal: animal}, nil
		}
	}
	return nil, errors.New("animal not found")
}

func newMockClient() *mockAnimalsClient {
	return &mockAnimalsClient{
		animals: []*pb.Animal{},
	}
}

func TestGetAnimals(t *testing.T) {
	// Test with empty animal list
	client := newMockClient()
	animals, err := GetAnimals(client)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(animals) != 0 {
		t.Fatalf("Expected empty list, got %d animals", len(animals))
	}

	// Test with animals
	client.animals = []*pb.Animal{
		{Name: "Dog", Color: "brown", Height: 40},
		{Name: "Cat", Color: "black", Height: 25},
	}

	animals, err = GetAnimals(client)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(animals) != 2 {
		t.Fatalf("Expected 2 animals, got %d", len(animals))
	}

	// Test error handling
	client.getError = errors.New("server error")
	_, err = GetAnimals(client)
	if err == nil || err.Error() != "server error" {
		t.Fatalf("Expected server error, got %v", err)
	}
}

func TestGetAnimalsMatching(t *testing.T) {
	client := newMockClient()
	client.animals = []*pb.Animal{
		{Name: "Dog", Color: "brown", Height: 40},
		{Name: "Cat", Color: "black", Height: 25},
		{Name: "Horse", Color: "brown", Height: 150},
	}

	// Test filter by color
	brownFilter := func(animal *pb.Animal) bool {
		return animal.Color == "brown"
	}
	brownAnimals, err := GetAnimalsMatching(client, brownFilter)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(brownAnimals) != 2 {
		t.Fatalf("Expected 2 brown animals, got %d", len(brownAnimals))
	}

	// Test filter by height
	tallFilter := func(animal *pb.Animal) bool {
		return animal.Height > 30
	}
	tallAnimals, err := GetAnimalsMatching(client, tallFilter)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(tallAnimals) != 2 {
		t.Fatalf("Expected 2 tall animals, got %d", len(tallAnimals))
	}

	// Test with no matches
	blueFilter := func(animal *pb.Animal) bool {
		return animal.Color == "blue"
	}
	blueAnimals, err := GetAnimalsMatching(client, blueFilter)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(blueAnimals) != 0 {
		t.Fatalf("Expected 0 blue animals, got %d", len(blueAnimals))
	}

	// Test error handling
	client.getError = errors.New("server error")
	_, err = GetAnimalsMatching(client, brownFilter)
	if err == nil || err.Error() != "server error" {
		t.Fatalf("Expected server error, got %v", err)
	}
}

func TestDeleteAnimals(t *testing.T) {
	client := newMockClient()
	client.animals = []*pb.Animal{
		{Name: "Dog", Color: "brown", Height: 40},
		{Name: "Cat", Color: "black", Height: 25},
	}

	// Test deleting all animals
	err := DeleteAnimals(client)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(client.animals) != 0 {
		t.Fatalf("Expected 0 animals after deletion, got %d", len(client.animals))
	}

	// Test error handling
	client.delError = errors.New("server error")
	err = DeleteAnimals(client)
	if err == nil || err.Error() != "server error" {
		t.Fatalf("Expected server error, got %v", err)
	}
}

func TestDeleteAnimalsMatching(t *testing.T) {
	client := newMockClient()
	client.animals = []*pb.Animal{
		{Name: "Dog", Color: "brown", Height: 40},
		{Name: "Cat", Color: "black", Height: 25},
		{Name: "Horse", Color: "brown", Height: 150},
	}

	// Test deleting animals by color
	brownFilter := func(animal *pb.Animal) bool {
		return animal.Color == "brown"
	}
	err := DeleteAnimalsMatching(client, brownFilter)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	
	// After deletion, only non-brown animals should remain
	animals, _ := GetAnimals(client)
	if len(animals) != 1 {
		t.Fatalf("Expected 1 animal after deletion, got %d", len(animals))
	}
	if animals[0].Color != "black" {
		t.Fatalf("Expected black animal, got %s", animals[0].Color)
	}

	// Test error during get
	client = newMockClient()
	client.animals = []*pb.Animal{
		{Name: "Dog", Color: "brown", Height: 40},
		{Name: "Cat", Color: "black", Height: 25},
	}
	client.getError = errors.New("get error")
	err = DeleteAnimalsMatching(client, brownFilter)
	if err == nil || err.Error() != "get error" {
		t.Fatalf("Expected get error, got %v", err)
	}

	// Test error during delete
	client = newMockClient()
	client.animals = []*pb.Animal{
		{Name: "Dog", Color: "brown", Height: 40},
		{Name: "Cat", Color: "black", Height: 25},
	}
	client.delError = errors.New("delete error")
	err = DeleteAnimalsMatching(client, brownFilter)
	if err == nil || err.Error() != "delete error" {
		t.Fatalf("Expected delete error, got %v", err)
	}

	// Test error during add
	client = newMockClient()
	client.animals = []*pb.Animal{
		{Name: "Dog", Color: "brown", Height: 40},
		{Name: "Cat", Color: "black", Height: 25},
	}
	client.addError = errors.New("add error")
	err = DeleteAnimalsMatching(client, brownFilter)
	if err == nil || err.Error() != "add error" {
		t.Fatalf("Expected add error, got %v", err)
	}
}

func TestAddBrownCat(t *testing.T) {
	client := newMockClient()

	// Test adding a brown cat
	cat, err := AddBrownCat(client, "Whiskers")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if cat.Name != "Whiskers" {
		t.Fatalf("Expected cat named Whiskers, got %s", cat.Name)
	}
	if cat.Color != "brown" {
		t.Fatalf("Expected brown cat, got %s", cat.Color)
	}
	if cat.Height != 30 {
		t.Fatalf("Expected height 30, got %d", cat.Height)
	}

	// Verify cat was added to client
	animals, _ := GetAnimals(client)
	if len(animals) != 1 {
		t.Fatalf("Expected 1 animal, got %d", len(animals))
	}

	// Test error handling
	client.addError = errors.New("add error")
	_, err = AddBrownCat(client, "Felix")
	if err == nil || err.Error() != "add error" {
		t.Fatalf("Expected add error, got %v", err)
	}
}

func TestGetAnimalColor(t *testing.T) {
	client := newMockClient()
	client.animals = []*pb.Animal{
		{Name: "Dog", Color: "brown", Height: 40},
		{Name: "Cat", Color: "black", Height: 25},
	}

	// Test getting color of existing animal
	color, err := GetAnimalColor(client, "Dog")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if color != "brown" {
		t.Fatalf("Expected brown color, got %s", color)
	}

	// Test getting color of non-existent animal
	_, err = GetAnimalColor(client, "Horse")
	if err == nil {
		t.Fatalf("Expected error for non-existent animal")
	}

	// Test error handling
	client.findError = errors.New("find error")
	_, err = GetAnimalColor(client, "Dog")
	if err == nil || err.Error() != "find error" {
		t.Fatalf("Expected find error, got %v", err)
	}
}

