package repository

import (
	"gb-ui-core/internal/ui/model"
	"gb-ui-core/pkg/terrors"
	"github.com/jmoiron/sqlx"
	"github.com/sarulabs/di"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func BuildPostgresRepository(ctn di.Container) (interface{}, error) {
	return &PostgresRepository{
		db: ctn.Get("postgres").(*sqlx.DB),
	}, nil
}

func (p *PostgresRepository) GetUiTypes() ([]*model.UiTypeDAO, error) {
	query := `
			 SELECT * FROM ui.types;
			 `
	data := make([]*model.UiTypeDAO, 0, 16)
	err := p.db.Select(&data, query)
	if err != nil {
		return nil, terrors.Raise(err, 300007)
	}

	return data, nil
}
