package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"

	"go-crud/model"
	"go-crud/store"
)

// TaskHandler — warstwa HTTP; zna Echo, nie zna SQL.
// Analogia FastAPI: router + Depends(get_task_store).
type TaskHandler struct {
	store *store.TaskStore
}

func NewTaskHandler(s *store.TaskStore) *TaskHandler {
	return &TaskHandler{store: s}
}

func (h *TaskHandler) RegisterRoutes(app *echo.Echo) {
	app.GET("/tasks", h.ListTasks)
	app.POST("/tasks", h.CreateTask)
	app.GET("/tasks/:id", h.GetTask)
	app.PUT("/tasks/:id", h.UpdateTask)
	app.DELETE("/tasks/:id", h.DeleteTask)
}

// GET /tasks — wzorzec: store robi SQL, handler mapuje err → HTTP.
func (h *TaskHandler) ListTasks(c *echo.Context) error {
	tasks, err := h.store.List(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}
	return c.JSON(http.StatusOK, tasks)
}

// POST /tasks
func (h *TaskHandler) CreateTask(c *echo.Context) error {
	var input model.TaskCreate
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json"})
	}
	if input.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "title can't be empty"})
	}
	task, err := h.store.Create(c.Request().Context(), input.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}
	return c.JSON(http.StatusCreated, task)
}

// GET /tasks/:id
func (h *TaskHandler) GetTask(c *echo.Context) error {
	id, err := parseID(c)
	if err != nil {
		return err
	}
	task, err := h.store.GetByID(c.Request().Context(), id)
	if errors.Is(err, sql.ErrNoRows) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}
	return c.JSON(http.StatusOK, task)
}

// PUT /tasks/:id
func (h *TaskHandler) UpdateTask(c *echo.Context) error {
	id, err := parseID(c)
	if err != nil {
		return err
	}

	var input model.TaskUpdate
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json"})
	}
	if input.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "title can't be empty"})
	}

	task, err := h.store.Update(c.Request().Context(), id, input)
	if errors.Is(err, store.ErrNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{"error":"task not found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}
	return c.JSON(http.StatusOK, task)
}

// DELETE /tasks/:id
func (h *TaskHandler) DeleteTask(c *echo.Context) error {
	id, err := parseID(c)
	if err != nil {
		return err
	}
	err = h.store.Delete(c.Request().Context(), id)
	if errors.Is(err, store.ErrNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{"error":"task not found"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "db error"})
	}
	return c.NoContent(http.StatusNoContent)
}

// parseID helper
func parseID(c *echo.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	return id, nil
}
