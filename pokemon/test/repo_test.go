package test

import (
	"context"
	"github.com/shwezhu/grpc_learning/pokemon/database"
	"github.com/shwezhu/grpc_learning/pokemon/server"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

func teardown(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		return
	}
}

func TestCreatePokemon(t *testing.T) {
	client, err := database.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer teardown(client)

	repo := server.NewRepository(client)
	testPokemon := server.Pokemon{
		Name:      "Pikachu",
		Types:     []string{"Electric"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = repo.CreatePokemon(testPokemon)
	assert.NoError(t, err)

	var result server.Pokemon
	err = repo.PokemonColl.FindOne(context.TODO(), bson.M{"name": "Pikachu"}).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "Pikachu", result.Name)
	assert.Equal(t, []string{"Electric"}, result.Types)
}

func TestUpdatePokemon(t *testing.T) {
	client, err := database.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer teardown(client)

	repo := server.NewRepository(client)
	testPokemon := server.Pokemon{
		ID:        "003",
		Name:      "Bulbasaur",
		Types:     []string{"Grass", "Poison"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = repo.PokemonColl.InsertOne(context.TODO(), testPokemon)
	assert.NoError(t, err)

	testPokemon.Name = "Ivysaur"
	testPokemon.Types = []string{"Grass"}
	err = repo.UpdatePokemon(context.TODO(), testPokemon)
	assert.NoError(t, err)

	var result server.Pokemon
	err = repo.PokemonColl.FindOne(context.TODO(), bson.M{"_id": "003"}).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "Ivysaur", result.Name)
	assert.Equal(t, []string{"Grass"}, result.Types)
}
