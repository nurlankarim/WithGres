package services

import (
	"WithGres/internal/configs"
	"WithGres/internal/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func FindAllItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	rows, err := configs.DB.Query("select * from items order by id")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var item models.Item
		err = rows.Scan(&item.Id, &item.Name, &item.Count, &item.Price)

		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		log.Fatal(err)
	}

}

func FindItemById(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	err = configs.DB.QueryRow("select * from items where id = $1", id).Scan(&item.Id, &item.Name, &item.Count, &item.Price)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var itemCreate models.ItemEdit
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &itemCreate)
	if err != nil {
		log.Fatal(err)
	}

	_, err = configs.DB.Exec("insert into items (name, count, price) values ($1, $2, $3)",
		itemCreate.Name, itemCreate.Count, itemCreate.Price)
	if err != nil {
		log.Println("DataBase error:", err)
		//log.Fatal(err)
		http.Error(w, "Failed to insert item", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("OK")
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var itemEdit models.ItemEdit
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &itemEdit)
	if err != nil {
		log.Fatal(err)
	}

	_, err = configs.DB.Exec("update items set name = $1, count = $2, price = $3 where id = $4",
		itemEdit.Name, itemEdit.Count, itemEdit.Price, id)
	if err != nil {
		//log.Fatal(err)
		log.Println("DataBase error:", err)

	}

	json.NewEncoder(w).Encode("OK")
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = configs.DB.Exec("delete from market_items where item_id = $1", id)
	if err != nil {
		log.Println("DataBase error:", err)

		//		log.Fatal(err)
	}

	_, err = configs.DB.Query("delete from items where id = $1", id)
	if err != nil {
		//log.Fatal(err)
		log.Println("DataBase error:", err)

	}

	json.NewEncoder(w).Encode("OK")
}

// package services

// import (
// 	"WithGres/internal/configs"
// 	"WithGres/internal/models"
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strconv"
// )

// func FindAllItems(w http.ResponseWriter, r *http.Request) {
// 	var items []models.Item
// 	rows, err := configs.DB.Query("select id, name, count, price from items order by id")
// 	if err != nil {
// 		log.Println("Error querying items:", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var item models.Item
// 		err = rows.Scan(&item.Id, &item.Name, &item.Count, &item.Price)
// 		if err != nil {
// 			log.Println("Error scanning item:", err)
// 			continue
// 		}
// 		items = append(items, item)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(items)
// }

// func CreateItem(w http.ResponseWriter, r *http.Request) {
// 	var itemCreate models.ItemEdit
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Println("Error reading body:", err)
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	err = json.Unmarshal(body, &itemCreate)
// 	if err != nil {
// 		log.Println("Error unmarshaling JSON:", err)
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}

// 	// ВАЖНО: Используем Exec для INSERT
// 	_, err = configs.DB.Exec("insert into items (name, count, price) values ($1, $2, $3)",
// 		itemCreate.Name, itemCreate.Count, itemCreate.Price)

// 	if err != nil {
// 		log.Println("Database INSERT error:", err)
// 		http.Error(w, "Database error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode("Created Successfully")
// }

// func UpdateItem(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	var itemEdit models.ItemEdit
// 	body, _ := ioutil.ReadAll(r.Body)
// 	json.Unmarshal(body, &itemEdit)

// 	// ВАЖНО: Используем Exec для UPDATE
// 	_, err = configs.DB.Exec("update items set name = $1, count = $2, price = $3 where id = $4",
// 		itemEdit.Name, itemEdit.Count, itemEdit.Price, id)

// 	if err != nil {
// 		log.Println("Database UPDATE error:", err)
// 		http.Error(w, "Database error", http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode("Updated OK")
// }

// func DeleteItem(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Сначала удаляем зависимости, если они есть, потом сам товар
// 	configs.DB.Exec("delete from market_items where item_id = $1", id)
// 	_, err = configs.DB.Exec("delete from items where id = $1", id)

// 	if err != nil {
// 		log.Println("Database DELETE error:", err)
// 		http.Error(w, "Database error", http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode("Deleted OK")
// }
