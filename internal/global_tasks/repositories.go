package global_tasks

import (
	"context"
	"my_pocket_taskbook/internal/db"
	"my_pocket_taskbook/internal/models"
)

type PostgresRepository struct {
	Storage *db.PostgresStorage
}

func NewRepo(s *db.PostgresStorage) *PostgresRepository {
	return &PostgresRepository{s}
}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]models.Task, error) {
	rows, err := r.Storage.Pool.Query(
		ctx,
		`SELECT id, title, body, status, type, created_at, updated_at
		FROM tasks 
		WHERE type=$1`,
		"global",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := []models.Task{}

	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Body, &task.Status, &task.Type, &task.CreatedAt, &task.UpdatedAt)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id int) (*models.Task, error) {
	var task models.Task

	err := r.Storage.Pool.QueryRow(
		ctx,
		`SELECT id, title, body, status, type, created_at, updated_at
		FROM tasks 
		WHERE type=$1 AND id=$2`,
		"global", id,
	).Scan(&task.ID, &task.Title, &task.Body, &task.Status, &task.Type, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *PostgresRepository) Create(ctx context.Context, t *models.Task) (*models.Task, error) {
	err := r.Storage.Pool.QueryRow(
		ctx,
		`INSERT INTO tasks(title, body, status, type)
		VALUES ($1, $2, $3, $4)
		RETURNING id, title, body, status, type, created_at, updated_at`,
		t.Title, t.Body, "active", "global",
	).Scan(&t.ID, &t.Title, &t.Body, &t.Status, &t.Type, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *PostgresRepository) Edit(ctx context.Context, t *models.Task, id int) (*models.Task, error) {
	err := r.Storage.Pool.QueryRow(
		ctx,
		`UPDATE tasks
		SET title = $1, body = $2
		WHERE id = $3 AND type=$4
		RETURNING id, title, body, status, type, created_at, updated_at`,
		t.Title, t.Body, id, "global",
	).Scan(&t.ID, &t.Title, &t.Body, &t.Status, &t.Type, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *PostgresRepository) ChangeStatus(ctx context.Context, id int, status string) (*models.Task, error) {
	var task models.Task

	err := r.Storage.Pool.QueryRow(
		ctx,
		`UPDATE tasks
		SET status = $1, updated_at = NOW()
		WHERE id = $2 AND type=$3
		RETURNING id, title, body, status, type, created_at, updated_at`,
		status, id, "global",
	).Scan(&task.ID, &task.Title, &task.Body, &task.Status, &task.Type, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &task, nil
}
