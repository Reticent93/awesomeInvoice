package repository

import (
	"awesomeInvoice/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetPlayers() ([]*models.Player, error)
	OnePlayer(id int) (*models.Player, error)
}
