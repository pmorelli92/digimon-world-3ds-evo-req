package postgres

import (
	"context"
	"digimon-world-3ds-evo-req-api/domain"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Repository is a contract that allows API to interact with database
type Repository interface {
	GetDigimons(ctx context.Context) ([]string, error)
	GetEvolutions(ctx context.Context, digimon string) ([]domain.Digimon, error)
}

type repository struct {
	dbHandler *pgxpool.Pool
}

// NewRepository returns a Repository contract for database interaction
func NewRepository(user string, password string, host string, port string, database string, ssl string) (Repository, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, database, ssl)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &repository{
		dbHandler: db,
	}, nil
}

func (r *repository) GetDigimons(ctx context.Context) ([]string, error) {
	rows, err := r.dbHandler.Query(ctx, "SELECT name FROM app.digimon")
	if err != nil {
		return nil, err
	}

	digimons := make([]string, 0)

	for rows.Next() {
		var i string
		rows.Scan(&i)
		digimons = append(digimons, i)
	}

	return digimons, nil
}

func (r *repository) GetEvolutions(ctx context.Context, digimon string) ([]domain.Digimon, error) {
	rows, err := r.dbHandler.Query(ctx,
		`SELECT
			name, hp, mp, atk, def, spd, int, weight, mistake,
			happiness, discipline, battles, techs, decode, quota
		 FROM app.evo_requirements WHERE evolves_from = $1`, digimon)

	if err != nil {
		return nil, err
	}

	digimons := make([]domain.Digimon, 0)

	for rows.Next() {
		var i domain.Digimon
		rows.Scan(
			&i.Name, &i.HP, &i.MP, &i.Atk, &i.Def, &i.Spd,
			&i.Int, &i.Weight, &i.Mistake, &i.Happiness,
			&i.Discipline, &i.Battles, &i.Techs, &i.Decode, &i.Quota)

		digimons = append(digimons, i)
	}

	return digimons, nil
}
