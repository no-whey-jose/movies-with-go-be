package repository

import (
	"context"
	"database/sql"
	"movies-be/internal/models"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 5

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			id, title, release_date, runtime, mpaa_rating, description, coalesce(image, ''), created_at, updated_at
		from
			movies
		order by
			title
	`

	data, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var movies []*models.Movie
	for data.Next() {
		var movie models.Movie
		err := data.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.Rated,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}
	return movies, nil
}
