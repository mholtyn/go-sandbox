package handler

import (
	"database/sql"
	"net/http"

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
	// TODO: register the rest of handlers
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
	// TODO: bind body into model.TaskCreate
	// TODO: INSERT the task and read back its generated id
	// TODO: respond 201 with the created task
	return nil
}

// TODO: GET /tasks/:id
func (h *TaskHandler) GetTask(c *echo.Context) error {
	// TODO: parse :id from the URL
	// TODO: query a single task by id
	// TODO: respond 404 if not found, otherwise 200
	return nil
}

// TODO: PUT /tasks/:id
func (h *TaskHandler) UpdateTask(c *echo.Context) error {
	// TODO: parse :id and bind the body
	// TODO: update the task by id
	// TODO: respond 404 if not found, otherwise 200
	return nil
}

// TODO: DELETE /tasks/:id
func (h *TaskHandler) DeleteTask(c *echo.Context) error {
	// TODO: parse :id from the URL
	// TODO: delete the task by id
	// TODO: respond 204 (No Content)
	return nil
}
