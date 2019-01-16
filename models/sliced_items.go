package models

import (
	"errors"
	"fmt"
	"strconv"
)

//ItemsStoreSliceType type
type ItemsStoreSliceType []Item

//ItemsStoreSlice instance
var ItemsStoreSlice ItemsStoreSliceType

//AddItem add item
func (i ItemsStoreSliceType) AddItem(item Item) {
	ItemsStoreSlice = append(ItemsStoreSlice, item)
	//fmt.Println("Store: ", ItemsStore)
}

//GetItems get all items
func (i ItemsStoreSliceType) GetItems() ItemsStoreSliceType {
	return ItemsStoreSlice
}

//GetItemById get item by id
func (i ItemsStoreSliceType) GetItemById(id string) (Item, error) {

	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("ERROR", err)
		return Item{}, err
	}

	fmt.Println(intId)
	fmt.Println(len(ItemsStoreSlice))
	if len(ItemsStoreSlice) >= intId {
		return ItemsStoreSlice[intId-1], nil
	} else {
		return Item{}, errors.New("id is incorrect")
	}
}
