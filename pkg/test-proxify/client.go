package test_proxify

import (
	"context"

	pb "github.com/craftizmv/DS/pkg/proto/animals"
)

type filterFunc func(*pb.Animal) bool

// GetAnimals returns all animals returned by the server
func GetAnimals(client pb.AnimalsClient) ([]*pb.Animal, error) {
	resp, err := client.GetAnimals(context.Background(), &pb.GetAnimalsRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Animals, nil
}

// GetAnimalsMatching returns animals matching condition specified in filter function
func GetAnimalsMatching(client pb.AnimalsClient, filter filterFunc) ([]*pb.Animal, error) {
	resp, err := client.GetAnimals(context.Background(), &pb.GetAnimalsRequest{})
	if err != nil {
		return nil, err
	}

	var result []*pb.Animal
	for _, animal := range resp.Animals {
		if filter(animal) {
			result = append(result, animal)
		}
	}
	return result, nil
}

// DeleteAnimals deletes all animals on the server
func DeleteAnimals(client pb.AnimalsClient) error {
	_, err := client.DeleteAnimals(context.Background(), &pb.DeleteAnimalsRequest{})
	return err
}

// DeleteAnimalsMatching deletes all animals on the server matching condition specified in filter function
func DeleteAnimalsMatching(client pb.AnimalsClient, filter filterFunc) error {
	// Get all animals
	resp, err := client.GetAnimals(context.Background(), &pb.GetAnimalsRequest{})
	if err != nil {
		return err
	}

	// Delete all animals
	_, err = client.DeleteAnimals(context.Background(), &pb.DeleteAnimalsRequest{})
	if err != nil {
		return err
	}

	// Re-add only the animals that don't match the filter
	for _, animal := range resp.Animals {
		if !filter(animal) {
			_, err := client.AddAnimal(context.Background(), &pb.AddAnimalRequest{
				Animal: animal,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// AddBrownCat adds animal with color brown and height of 30
func AddBrownCat(client pb.AnimalsClient, catName string) (*pb.Animal, error) {
	resp, err := client.AddAnimal(context.Background(), &pb.AddAnimalRequest{
		Animal: &pb.Animal{
			Name:   catName,
			Color:  "brown",
			Height: 30,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp.Animal, nil
}

// GetAnimalColor returns color of animal with given name
func GetAnimalColor(client pb.AnimalsClient, animalName string) (string, error) {
	resp, err := client.GetAnimal(context.Background(), &pb.GetAnimalRequest{
		Name: animalName,
	})
	if err != nil {
		return "", err
	}
	return resp.Animal.Color, nil
}
