package main

import (
	"fmt"
	"log"
	"my_pocket_taskbook/internal/db"
	"my_pocket_taskbook/internal/global_tasks"
	"my_pocket_taskbook/internal/local_tasks"
	"my_pocket_taskbook/internal/routine_tasks"
	"net/http"
	"strings"
)

const PORT = ":8080"

func main() {
	storage, err := db.New()

	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	defer storage.Pool.Close()

	fmt.Println("Connected to PostgreSQL!")

	if err := storage.Migrate(); err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}

	mux := http.NewServeMux()

	globalTasksRepo := global_tasks.NewRepo(storage)
	globalTasksService := global_tasks.NewService(globalTasksRepo)
	globalTasksHandler := global_tasks.NewHandler(globalTasksService)

	localTasksRepo := local_tasks.NewRepo(storage)
	localTasksService := local_tasks.NewService(localTasksRepo)
	localTasksHandler := local_tasks.NewHandler(localTasksService)

	routineTasksRepo := routine_tasks.NewRepo(storage)
	routineTasksService := routine_tasks.NewService(routineTasksRepo)
	routineTasksHandler := routine_tasks.NewHandler(routineTasksService)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		if len(parts) == 1 && parts[0] == "" {
			fmt.Fprintf(w, "Server is working!")
		} else {
			if parts[1] == "global" {
				if len(parts) == 2 {
					if r.Method == http.MethodGet {
						globalTasksHandler.GetAll(w, r)
					}

					if r.Method == http.MethodPost {
						globalTasksHandler.Create(w, r)
					}
				}

				if len(parts) == 3 {
					if r.Method == http.MethodGet {
						globalTasksHandler.GetByID(w, r)
					}

					if r.Method == http.MethodPut {
						globalTasksHandler.Edit(w, r)
					}
				}

				if len(parts) == 4 {
					if r.Method == http.MethodPatch {
						if parts[3] == "active" || parts[3] == "completed" || parts[3] == "canceled" {
							globalTasksHandler.ChangeStatus(w, r)
						}
					}
				}
			}

			if parts[1] == "local" {
				if len(parts) == 2 {
					if r.Method == http.MethodGet {
						localTasksHandler.GetAll(w, r)
					}

					if r.Method == http.MethodPost {
						localTasksHandler.Create(w, r)
					}
				}

				if len(parts) == 3 {
					if r.Method == http.MethodGet {
						localTasksHandler.GetByID(w, r)
					}

					if r.Method == http.MethodPut {
						localTasksHandler.Edit(w, r)
					}
				}

				if len(parts) == 4 {
					if r.Method == http.MethodPatch {
						if parts[3] == "active" || parts[3] == "completed" || parts[3] == "canceled" {
							localTasksHandler.ChangeStatus(w, r)
						}
					}
				}
			}

			if parts[1] == "routine" {
				if len(parts) == 2 {
					if r.Method == http.MethodGet {
						routineTasksHandler.GetAll(w, r)
					}

					if r.Method == http.MethodPost {
						routineTasksHandler.Create(w, r)
					}
				}

				if len(parts) == 3 {
					if r.Method == http.MethodGet {
						routineTasksHandler.GetByID(w, r)
					}

					if r.Method == http.MethodPut {
						routineTasksHandler.Edit(w, r)
					}
				}

				if len(parts) == 4 {
					if r.Method == http.MethodPatch {
						if parts[3] == "active" || parts[3] == "completed" || parts[3] == "canceled" || parts[3] == "retired" {
							routineTasksHandler.ChangeStatus(w, r)
						}
					}
				}
			}

			if parts[1] == "current" {
				if len(parts) == 2 {
					if r.Method == http.MethodGet {
						localTasksHandler.GetAllCurrent(w, r)
					}
				}
			}
		}
	})

	fmt.Println("Server listening to", PORT)
	http.ListenAndServe(PORT, mux)
}
