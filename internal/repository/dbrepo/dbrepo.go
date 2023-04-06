package dbrepo

import (
	"awesomeInvoice/internal/models"
	"context"
	"database/sql"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

func (p *PostgresDBRepo) Connection() *sql.DB {
	return p.DB
}

const dbTimeout = time.Second * 3

func (p *PostgresDBRepo) GetPlayers() ([]*models.Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name, position, team, age, jersey, retired, created_at, updated_at FROM players ORDER BY name`

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var players []*models.Player

	for rows.Next() {
		var player models.Player
		err := rows.Scan(&player.ID, &player.Name, &player.Position, &player.Team, &player.Age, &player.Jersey, &player.Retired, &player.CreatedAt, &player.UpdatedAt)
		if err != nil {
			return nil, err
		}
		players = append(players, &player)
	}

	return players, nil
}

func (p *PostgresDBRepo) OnePlayer(id int) (*models.Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name, position, team, age, jersey, retired, created_at, updated_at FROM players WHERE id = $1`

	row := p.DB.QueryRowContext(ctx, query, id)

	var player models.Player
	err := row.Scan(&player.ID, &player.Name, &player.Position, &player.Team, &player.Age, &player.Jersey, &player.Retired, &player.CreatedAt, &player.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &player, nil
}
