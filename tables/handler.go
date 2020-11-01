package tables

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetRoutes sets available routes to tables api
func SetRoutes(mongo *mongo.Client, logger ...*log.Logger) http.Handler {
	var c config
	c.mongo = mongo
	if logger != nil {
		c.log = logger[0]
	} else {
		c.log = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
	}

	r := chi.NewRouter()

	r.Get("/", c.getTablesHandler)
	r.Get("/{ID}", getTableByIDHandler)
	r.Post("/", createTableHandler)
	r.Put("/{ID}", editTableByIDHandler)
	r.Delete("/{ID}", deleteTableByIDHandler)

	return r
}

func (c config) getTablesHandler(w http.ResponseWriter, r *http.Request) {
	tables, err := getTables(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(tables)
	if err != nil {
		c.log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		c.log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTableByIDHandler(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func createTableHandler(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func editTableByIDHandler(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func deleteTableByIDHandler(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
