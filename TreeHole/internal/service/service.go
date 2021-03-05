package service

import (
	"database/sql"
	"github.com/hako/branca"
)


// Service contains the core logic, you can back a REST, GraphQL or Grpc api
type Service struct {
	db *sql.DB
	codec *branca.Branca
}

// New Service implement
func New(db *sql.DB, codec *branca.Branca) *Service {
	return &Service{
		db: db,
		codec: codec,
	}
}