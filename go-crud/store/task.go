package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go-crud/model"
)

// UPDATE/DELETE matched 0 rows.
var ErrNotFound = errors.New("task not found")

type TaskStore struct {
	db *sql.DB
}

func NewTaskStore(db *sql.DB) *TaskStore {
	return &TaskStore{db: db}
}

// List
func (s *TaskStore) List(ctx context.Context) ([]model.TaskPublic, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, title, done, created_at FROM tasks ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("list tasks: %w", err)
	}
	defer rows.Close()

	tasks := []model.TaskPublic{}
	for rows.Next() {
		var t model.TaskPublic
		if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan task: %w", err)
		}
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tasks: %w", err)
	}

	return tasks, nil
}

// GetByID
func (s *TaskStore) GetByID(ctx context.Context, id int) (model.TaskPublic, error) {
	var task model.TaskPublic
	err := s.db.QueryRowContext(ctx, 
		"SELECT id, title, done, created_at FROM tasks WHERE id = ?", 
		id).Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt)
	if err != nil {
		return model.TaskPublic{}, err
	}
	return task, nil
}

// Create
func (s *TaskStore) Create(ctx context.Context, title string) (model.TaskPublic, error) {
	result, err := s.db.ExecContext(ctx, "INSERT INTO tasks (title) VALUES (?)", title)
	if err != nil {
		return model.TaskPublic{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return model.TaskPublic{}, err
	}
	return s.GetByID(ctx, int(id))
}

// Update
func (s *TaskStore) Update(ctx context.Context, id int, input model.TaskUpdate) (model.TaskPublic, error) {
	result, err := s.db.ExecContext(ctx,
		"UPDATE tasks SET title = ?, done = ? WHERE id = ?",
		input.Title, input.Done, id)
	if err != nil {
		return model.TaskPublic{}, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return model.TaskPublic{}, err
	}
	if n == 0 {
		return model.TaskPublic{}, ErrNotFound
	}

	return s.GetByID(ctx, id)
}

// Delete
func (s *TaskStore) Delete(ctx context.Context, id int) error {
	result, err := s.db.ExecContext(ctx, "DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}
