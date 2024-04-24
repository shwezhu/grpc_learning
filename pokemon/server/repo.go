package server

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var ErrNotFound = errors.New("pokemon not found")

type Repository struct {
	db *mongo.Client
	//Collection is safe for concurrent use by multiple goroutines.
	PokemonColl *mongo.Collection
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		db:          db,
		PokemonColl: db.Database("pokemon").Collection("pokemons"),
	}
}

func (r *Repository) CreatePokemon(ctx context.Context, pokemon *Pokemon) error {
	_, err := r.PokemonColl.InsertOne(ctx, pokemon)
	if err != nil {
		return fmt.Errorf("failed to insert pokemon: %w", err)
	}
	return nil
}

func (r *Repository) GetAllPokemon(ctx context.Context) ([]Pokemon, error) {
	cursor, err := r.PokemonColl.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all pokemon: %w", err)
	}

	var results []Pokemon
	// This method will close the cursor after retrieving all documents.
	// If the number and size of documents is very large, your program will crash.
	// Learn more: https://arc.net/l/quote/eqpmlspn
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, fmt.Errorf("failed to decode pokemon: %w", err)
	}

	return results, nil
}

func (r *Repository) GetPokemonByID(ctx context.Context, id string) (Pokemon, error) {
	var pokemon Pokemon
	err := r.PokemonColl.FindOne(ctx, bson.D{{"_id", id}}).Decode(&pokemon)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Pokemon{}, ErrNotFound
		}
		return Pokemon{}, fmt.Errorf("failed to get pokemon by id: %w", err)
	}
	return pokemon, nil
}

func (r *Repository) UpdatePokemon(ctx context.Context, pokemon Pokemon) error {
	filter := bson.D{{"_id", pokemon.ID}}

	pokemon.UpdatedAt = time.Now()
	update := bson.D{
		{"$set", bson.D{
			{"name", pokemon.Name},
			{"types", pokemon.Types},
			{"updated_at", pokemon.UpdatedAt},
		}},
	}

	result, err := r.PokemonColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("no document found with the given ID to update")
	}

	return nil
}
