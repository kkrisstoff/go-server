package models

import (
	"fmt"
	"strconv"

	"github.com/kkrisstoff/go-server/utils/id_generator"
)

//Item struct type
type Item struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

//ItemsStoreMapped Store type
type ItemsStoreMappedType struct {
	length int
	Store  map[int]Item
}

//ItemsStoreMapped instance
var ItemsStoreMapped = ItemsStoreMappedType{
	0,
	map[int]Item{},
}

//AddItem adds new item to store
func (items ItemsStoreMappedType) AddItem(name string, message string) Item {
	id := getID()
	item := Item{
		ID:      id,
		Name:    name,
		Message: message,
	}
	items.Store[id] = item
	items.length++

	//fmt.Printf("New item id: %d, message: %v\n", id, message)
	return item
}

//GetItems get all items
func (items ItemsStoreMappedType) GetItems() []Item {
	return mapToSlice(items.Store)
}

//GetItemByID get item by id
func (items ItemsStoreMappedType) GetItemByID(idStr string) Item {
	id := idToInt(idStr)
	item := items.Store[id]
	return item
}

//DeleteItemByID delete item by id
func (items ItemsStoreMappedType) DeleteItemByID(idStr string) {
	id := idToInt(idStr)
	delete(items.Store, id)
	items.length--
}

func getID() int {
	id := id_generator.Generator.Generate()

	return id
}

func idToInt(id string) int {
	intID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return intID
}

func mapToSlice(m map[int]Item) []Item {
	var items []Item
	for _, value := range m {
		items = append(items, value)
	}

	return items
}
