package routine_tasks

import (
	"encoding/json"
	"my_pocket_taskbook/internal/models"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	routineTasks, err := h.service.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":       "Routine tasks fetched successfully",
		"routine tasks": routineTasks,
	})
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, "invalid routine task id", http.StatusBadRequest)
		return
	}

	routineTask, err := h.service.GetByID(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":      "Routine task fetched successfully",
		"routine task": routineTask,
	})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var routineTask models.Task

	err := json.NewDecoder(r.Body).Decode(&routineTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newRoutineTask, err := h.service.Create(r.Context(), &routineTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message":      "Routine task created successfully",
		"routine task": newRoutineTask,
	})
}

func (h *Handler) Edit(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, "invalid routine task id", http.StatusBadRequest)
		return
	}

	var routineTask models.Task

	err = json.NewDecoder(r.Body).Decode(&routineTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	editedRoutineTask, err := h.service.Edit(r.Context(), &routineTask, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":      "Routine task edited successfully",
		"routine task": editedRoutineTask,
	})
}

func (h *Handler) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, "invalid routine task id", http.StatusBadRequest)
		return
	}

	var status string = parts[3]

	statusChangedRoutineTask, err := h.service.ChangeStatus(r.Context(), id, status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":      "Routine task status updated successfully",
		"routine task": statusChangedRoutineTask,
	})
}
