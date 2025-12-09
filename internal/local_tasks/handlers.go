package local_tasks

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
	localTasks, err := h.service.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":     "local tasks fetched successfully",
		"local tasks": localTasks,
	})
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, "invalid local task id", http.StatusBadRequest)
		return
	}

	localTask, err := h.service.GetByID(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":    "local task fetched successfully",
		"local task": localTask,
	})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var localTask models.Task

	err := json.NewDecoder(r.Body).Decode(&localTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newLocalTask, err := h.service.Create(r.Context(), &localTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message":    "local task created successfully",
		"local task": newLocalTask,
	})
}

func (h *Handler) Edit(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, "invalid local task id", http.StatusBadRequest)
		return
	}

	var localTask models.Task

	err = json.NewDecoder(r.Body).Decode(&localTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	editedLocalTask, err := h.service.Edit(r.Context(), &localTask, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":    "local task edited successfully",
		"local task": editedLocalTask,
	})
}

func (h *Handler) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, "invalid local task id", http.StatusBadRequest)
		return
	}

	var status string = parts[3]

	statusChangedLocalTask, err := h.service.ChangeStatus(r.Context(), id, status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":    "Local task status updated successfully",
		"local task": statusChangedLocalTask,
	})
}
