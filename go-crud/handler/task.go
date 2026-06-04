package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"

	"go-crud/model"
)

// TaskHandler holds connection with db which every handler uses (Depends() in FastAPI)
type TaskHandler struct {
	db *sql.DB
}

// NewTaskHandler builds the handler with its db injected
func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{db: db}
}

// RegisterRoutes attaches endpoints to app (.include_router in FastAPI)
func (h *TaskHandler) RegisterRoutes(app *echo.Echo) {
	app.GET("/tasks", h.ListTasks)
	app.POST("/tasks", h.CreateTask)
	app.GET("/tasks/:id", h.GetTask)
	app.PUT("/tasks/:id", h.UpdateTask)
	app.DELETE("/tasks/:id", h.DeleteTask)
}

// TODO: GET /tasks
func (h *TaskHandler) ListTasks(c *echo.Context) error {
	rows, err := h.db.QueryContext(c.Request().Context(), "SELECT id, title, done, created_at FROM tasks ORDER BY id")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}
	defer rows.Close()

	tasks := []model.TaskPublic{}
	for rows.Next() {
		var t model.TaskPublic
		if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
		}
		tasks = append(tasks, t)
	}

	return c.JSON(http.StatusOK, tasks)
}

// TODO: POST /tasks
func (h *TaskHandler) CreateTask(c *echo.Context) error {
	var input model.TaskCreate
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json"})	
	}
	if input.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "title can't be empty"})
	}

	ctx := c.Request().Context()
	result, err := h.db.ExecContext(ctx, "INSERT INTO tasks (title) VALUES (?)", input.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}

	var task model.TaskPublic
	err = h.db.QueryRowContext(ctx,
		"SELECT id, title, done, created_at FROM tasks WHERE id = ?",
		id).Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}

	return c.JSON(http.StatusCreated, task)
}

// TODO: GET /tasks/:id
func (h *TaskHandler) GetTask(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var task model.TaskPublic
	err = h.db.QueryRowContext(c.Request().Context(),
		"SELECT id, title, done, created_at FROM tasks WHERE id = ?",
		id).Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}

	return c.JSON(http.StatusOK, task)
}

// TODO: PUT /tasks/:id
func (h *TaskHandler) UpdateTask(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "not found"})
	}

	var input model.TaskUpdate
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json"})
	}
	if input.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "title can't be empty"})
	}

	ctx := c.Request().Context()
	result, err := h.db.ExecContext(ctx,
	"UPDATE tasks SET title = ?, done = ? WHERE ID = ?",
	input.Title, input.Done, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}

	n, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}
	if n == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}

	var task model.TaskPublic
	err = h.db.QueryRowContext(ctx,
	"SELECT task FROM tasks WHERE id = ?",
	id).Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}

	return c.JSON(http.StatusOK, task)
}

// TODO: DELETE /tasks/:id
func (h *TaskHandler) DeleteTask(c *echo.Context) error {
	idStr := c.Param(("id"))
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	result, err := h.db.ExecContext(c.Request().Context(),
	"DELETE FROM tasks WHERE id = ?",
	id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}
	n, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}
	if n == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}

	return c.NoContent(http.StatusNoContent)
}
