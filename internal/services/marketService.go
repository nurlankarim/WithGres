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

func FindAllMarkets(w http.ResponseWriter, r *http.Request) {
	var marketsView []models.MarketView

	rows, err := configs.DB.Query("select * from markets order by id")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var market models.Market
		err = rows.Scan(&market.Id, &market.Name, &market.Address, &market.PhoneNumber)
		if err != nil {
			log.Fatal(err)
		}

		marketsView = append(marketsView, market.MapToView(getItemsByMarketId(market.Id)))
	}

	err = json.NewEncoder(w).Encode(marketsView)
	if err != nil {
		log.Fatal(err)
	}
}

func getItemsByMarketId(marketId int) []models.ItemEdit {
	var items []models.ItemEdit
	itemRows, err := configs.DB.Query("select distinct i.name, i.count, i.price from items i "+
		"inner join market_items mi on mi.item_id = i.id where mi.market_id = $1", marketId)
	if err != nil {
		log.Fatal(err)
	}

	for itemRows.Next() {
		var item models.ItemEdit
		err = itemRows.Scan(&item.Name, &item.Count, &item.Price)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}

	return items
}

func FindMarketById(w http.ResponseWriter, r *http.Request) {
	var market models.MarketView

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	err = configs.DB.QueryRow("select name, address, phone_number from markets where id = $1", id).Scan(&market.Name, &market.Address, &market.PhoneNumber)
	if err != nil {
		log.Fatal(err)
	}

	market.Items = getItemsByMarketId(id)

	err = json.NewEncoder(w).Encode(market)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateMarket(w http.ResponseWriter, r *http.Request) {
	var marketCreate models.MarketEdit
	var newMarket models.Market
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &marketCreate)
	if err != nil {
		log.Fatal(err)
	}

	err = configs.DB.QueryRow("insert into markets (name, address, phone_number) values ($1, $2, $3) returning *", marketCreate.Name, marketCreate.Address, marketCreate.PhoneNumber).Scan(&newMarket.Id, &newMarket.Name, &newMarket.Address, &newMarket.PhoneNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, itemId := range marketCreate.ItemIds {
		configs.DB.QueryRow("insert into market_items (market_id, item_id) values ($1, $2)", newMarket.Id, itemId)
	}

	err = json.NewEncoder(w).Encode(newMarket)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateMarket(w http.ResponseWriter, r *http.Request) {
	var editMarket models.MarketEdit

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &editMarket)
	if err != nil {
		log.Fatal(err)
	}

	configs.DB.QueryRow("delete from market_items where market_id = $1", id)
	for _, itemId := range editMarket.ItemIds {
		configs.DB.QueryRow("insert into market_items (market_id, item_id) "+
			"values ($1, $2)", id, itemId)
	}

	_, err = configs.DB.Query("update markets set name = $1, address = $2, phone_number = $3 where id = $4",
		editMarket.Name, editMarket.Address, editMarket.PhoneNumber, id)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode("OK")
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteMarket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	configs.DB.QueryRow("delete from market_items where market_id = $1", id)

	configs.DB.QueryRow("delete from markets where id = $1", id)

	err = json.NewEncoder(w).Encode("OK")
	if err != nil {
		log.Fatal(err)
	}
}
