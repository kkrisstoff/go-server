package models

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/kkrisstoff/go-server/api/id_generator"
)

//ItemsStoreMapped Store type
type ItemsStoreMappedType struct {
	length int
	Store  map[int]Item
}

//ItemsStoreMapped instance
var ItemsStoreMapped = ItemsStoreMappedType{
	0,
	map[int]Item{}}

//AddItem adds new item to store
func (items ItemsStoreMappedType) AddItem(message string) Item {
	id := getID()
	msg := message
	item := buildItem(id, msg)

	items.Store[id] = item
	items.length++
	return item
}

//GetItems get all items
func (items ItemsStoreMappedType) GetItems() []Item {
	return mapToSlice(items.Store)
}

//GetItemByID get item by id
func (items ItemsStoreMappedType) GetItemByID(idStr string) (Item, error) {
	id := idToInt(idStr)
	item, ok := items.Store[id]
	if ok {
		return item, nil
	}
	return Item{}, errors.New("ID is not exist")
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

func buildItem(id int, msg string) Item {
	item := Item{
		ID:      id,
		Message: msg}

	fmt.Printf("New item id: %d, message: %v\n", id, msg)
	return item
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
