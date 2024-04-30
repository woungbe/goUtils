package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items []Item

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", GetItem).Methods("GET")
	router.HandleFunc("/items", CreateItem).Methods("POST")
	router.HandleFunc("/items", UpdateItem).Methods("PUT")
	router.HandleFunc("/items", DeleteItem).Methods("DELETE")

	// 라우터를 설정하고 웹 서버 시작
	http.Handle("/", router)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("이렇게라도 보내겠습니다.")
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println("id : ", id)
	// for _, item := range items {
	// 	if item.ID == id {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	json.NewEncoder(w).Encode(id)
	// json.NewEncoder(w).Encode(&Item{})
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range items {
		if item.ID == id {
			items = append(items[:index], items[index+1:]...)
			var newItem Item
			_ = json.NewDecoder(r.Body).Decode(&newItem)
			newItem.ID = id
			items = append(items, newItem)
			json.NewEncoder(w).Encode(newItem)
			return
		}
	}

	json.NewEncoder(w).Encode(items)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range items {
		if item.ID == id {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(items)
}
