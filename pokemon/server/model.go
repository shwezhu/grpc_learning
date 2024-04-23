package server

import (
	"github.com/shwezhu/grpc_learning/pokemon/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Pokemon struct {
	ID        string    `bson:"_id,omitempty"`
	Name      string    `bson:"name,omitempty"`
	Types     []string  `bson:"types,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}

func pokemonToResponse(p Pokemon) pb.Pokemon {
	types := make([]pb.Type, len(p.Types))
	for i, t := range p.Types {
		types[i] = pb.Type(pb.Type_value[t])
	}

	return pb.Pokemon{
		Id:        p.ID,
		Name:      p.Name,
		Type:      types,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
	}
}
