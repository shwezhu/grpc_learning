package server

import (
	"context"
	"fmt"
	"github.com/shwezhu/grpc_learning/pokemon/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Server struct {
	repository *Repository
	pb.UnimplementedPokemonServer
}

func NewServer(repository *Repository) *Server {
	return &Server{
		repository: repository,
	}
}

func (s *Server) Create(ctx context.Context, req *pb.PokemonRequest) (*pb.Pokemon, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "ID cannot be empty")
	}

	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name cannot be empty")
	}

	if len(req.Type) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Type cannot be empty")
	}

	types := make([]string, len(req.Type))
	for i, t := range req.Type {
		types[i] = t.String()
	}

	now := time.Now()

	pokemon := &Pokemon{
		ID:        req.Id,
		Name:      req.Name,
		Types:     types,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repository.CreatePokemon(ctx, pokemon); err != nil {
		return nil, fmt.Errorf("failed to create pokemon: %w", err)
	}

	return pokemonToResponse(pokemon), nil
}

func (s *Server) Read(ctx context.Context, _ *pb.PokemonFilter) (*pb.PokemonListResponse, error) {
	res, err := s.repository.GetAllPokemon(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all pokemon: %w", err)
	}

	pokemons := make([]*pb.Pokemon, len(res))
	for i, t := range res {
		pokemons[i] = pokemonToResponse(&t)
	}

	return &pb.PokemonListResponse{Pokemon: pokemons}, nil
}

func (s *Server) ReadOne(ctx context.Context, req *pb.PokemonID) (*pb.Pokemon, error) {
	pokemon, err := s.repository.GetPokemonByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get pokemon by id: %w", err)
	}

	return pokemonToResponse(&pokemon), nil
}

func (s *Server) Update(ctx context.Context, req *pb.PokemonUpdateRequest) (*pb.Pokemon, error) {
	pokemon, err := s.repository.GetPokemonByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get pokemon by id: %w", err)
	}

	if req.Name != "" {
		pokemon.Name = req.Name
	}

	if len(req.Type) > 0 {
		types := make([]string, len(req.Type))
		for i, t := range req.Type {
			types[i] = t.String()
		}
		pokemon.Types = types
	}

	if err := s.repository.UpdatePokemon(ctx, pokemon); err != nil {
		return nil, fmt.Errorf("failed to update pokemon: %w", err)
	}

	return pokemonToResponse(&pokemon), nil
}
