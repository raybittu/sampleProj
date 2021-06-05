package inventory

import "likeMindProj/item"

type Inventory struct {
	Item *item.Item
	Quantity int
}

func RemoveInv(inv []Inventory, ind int) []Inventory {
		return append(inv[:ind], inv[ind+1:]...)
}

func AddItemToInventory( inv []Inventory, item *item.Item, quantity int) []Inventory {
	flag:=false
	for i:=range inv{
		if inv[i].Item.Brand==item.Brand && inv[i].Item.Category==item.Category{
			inv[i].Quantity+=quantity
			flag=true
		}
	}

	if !flag{
		invent:= Inventory{Item: item, Quantity: quantity}
		inv = append(inv, invent)
	}

	return inv
}