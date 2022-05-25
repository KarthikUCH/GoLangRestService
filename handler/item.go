package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KarthikUCH/buildapi/db"
)

func GetItems(dbInstance db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := dbInstance.GetAllItems()

		if err != nil {
			fmt.Println("Error accecessing all items")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(items)
		//fmt.Fprint(w, items)
	}
}

func GetItem(dbInstance db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func SaveItem(dbInstance db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func UpdateItem(dbInstance db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func DeleteItem(dbInstance db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
