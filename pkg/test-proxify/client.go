package test_proxify

import (
	"errors"

	pb "github.com/codility/protobuf_client/proto/animals"
)

type filterFunc func(pb.Animal) bool

// GetAnimals returns all animals returned by the server
func GetAnimals(client pb.AnimalsClient) ([]pb.Animal, error) {
	return nil, errors.New("Not implemented")
	result := make([]int, 0)
	result = append(result, 1)
	for i := range collection {

	}
}

// GetAnimalsMatching returns animals matching condition specified in filter function
func GetAnimalsMatching(client pb.AnimalsClient, filter filterFunc) ([]pb.Animal, error) {
	return []pb.Animal{}, errors.New("Not implemented")
}

// DeleteAnimals deletes all animals on the server
func DeleteAnimals(client pb.AnimalsClient) error {
	return errors.New("Not implemented")
}

// DeleteAnimalsMatching deletes all animals on the server matching condition specified in filter function
func DeleteAnimalsMatching(client pb.AnimalsClient, filter filterFunc) error {
	return errors.New("Not implemented")
}

// AddBrownCat adds animal with color brown and height of 30
func AddBrownCat(client pb.AnimalsClient, catName string) (*pb.Animal, error) {
	return nil, errors.New("Not implemented")
}

// GetAnimalColor returns color of animal with given name
func GetAnimalColor(client pb.AnimalsClient, animalName string) (string, error) {
	return "", errors.New("Not implemented")
}
