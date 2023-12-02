package repository

import (
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/domain/ports"
	"github.com/jackc/pgx/v4/pgxpool"
)

// postgresqlRepo Struct
type postgresqlRepo struct {
	db *pgxpool.Pool
}

// NewPostgresqlRepository Auth Domain postgresql repository constructor
func NewPostgresqlRepository(db *pgxpool.Pool) ports.IPostgresqlRepository {
	return &postgresqlRepo{db: db}
}
