package main

import (
	"database/sql"
	"digimon-world-3ds-evo-req-api/domain"
	"digimon-world-3ds-evo-req-api/postgres"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/go-chi/chi"
)

func main() {

	// Bootstrap environment variables.
	postgresUser := envOrDefault("POSTGRES_USER", "dbuser")
	postgresPassword := envOrDefault("POSTGRES_PASSWORD", "dbpassword")
	postgresHost := envOrDefault("POSTGRES_HOST", "0.0.0.0")
	postgresPort := envOrDefault("POSTGRES_PORT", "5432")
	postgresDB := envOrDefault("POSTGRES_DB", "digimonsql")
	postgresSSL := envOrDefault("POSTGRES_SSL", "disable")
	httpPort := envOrDefault("HTTP_PORT", "8080")

	// Postgres repository
	repo, err := postgres.NewRepository(
		postgresUser, postgresPassword, postgresHost,
		postgresPort, postgresDB, postgresSSL)

	if err != nil {
		fmt.Println("unable to connect to the database")
		os.Exit(1)
	}

	// HTTP server
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", healthHandler())
		r.Get("/digimon", digimonHandler(repo))
		r.Get("/digimon/{name}", possibleEvolutionsHandler(repo))
	})

	err = http.ListenAndServe(fmt.Sprintf(":%s", httpPort), r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func healthHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-type", "application/json")
		bytes, _ := json.Marshal(struct{ OK bool }{OK: true})
		rw.Write(bytes)
	}
}

func digimonHandler(repo postgres.Repository) http.HandlerFunc {

	type response struct {
		Digimons []string `json:"digimons"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {

		digimons, err := repo.GetDigimons(r.Context())
		if err == sql.ErrNoRows {
			rw.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-type", "application/json")
		bytes, _ := json.Marshal(response{digimons})
		rw.Write(bytes)
	}
}

func possibleEvolutionsHandler(repo postgres.Repository) http.HandlerFunc {

	type response struct {
		Digimons []domain.Digimon `json:"digimons"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		input := chi.URLParam(r, "name")

		digimons, err := repo.GetEvolutions(r.Context(), input)
		if err == sql.ErrNoRows {
			rw.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-type", "application/json")
		bytes, _ := json.Marshal(response{digimons})
		rw.Write(bytes)
	}
}

func envOrDefault(key string, fallback string) string {
	value, ok := syscall.Getenv(key)
	if !ok {
		return fallback
	}
	return value
}
