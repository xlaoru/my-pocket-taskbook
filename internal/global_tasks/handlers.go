package global_tasks

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
	globalTasks, err := h.service.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":      "Global tasks fetched successfully",
		"global tasks": globalTasks,
	})
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, "invalid global task id", http.StatusBadRequest)
		return
	}

	globalTask, err := h.service.GetByID(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":     "Global task fetched successfully",
		"global task": globalTask,
	})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var globalTask models.Task

	err := json.NewDecoder(r.Body).Decode(&globalTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newGlobalTask, err := h.service.Create(r.Context(), &globalTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message":     "Global task created successfully",
		"global task": newGlobalTask,
	})
}

func (h *Handler) Edit(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, "invalid global task id", http.StatusBadRequest)
		return
	}

	var globalTask models.Task

	err = json.NewDecoder(r.Body).Decode(&globalTask)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	editedGlobalTask, err := h.service.Edit(r.Context(), &globalTask, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":     "Global task edited successfully",
		"global task": editedGlobalTask,
	})
}

func (h *Handler) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, "invalid global task id", http.StatusBadRequest)
		return
	}

	var status string = parts[3]

	statusChangedGlobalTask, err := h.service.ChangeStatus(r.Context(), id, status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":     "Global task status updated successfully",
		"global task": statusChangedGlobalTask,
	})
}
