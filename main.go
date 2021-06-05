package main

import (
	"fmt"
	"likeMindProj/cart"
	inventory2 "likeMindProj/inventory"
	"likeMindProj/item"
	"likeMindProj/user"
)

func AddInv(inv []inventory2.Inventory, itm []item.Item, ic string,b string,q int) []inventory2.Inventory {
	for i:=range itm{
		if itm[i].Category==ic && itm[i].Brand==b{
			inv = inventory2.AddItemToInventory(inv, &itm[i], q)
		}
	}
	return inv
}

func ADC(c []cart.Cart, us []user.User, it []item.Item, un string, ic string, b  string, q  int) []cart.Cart {
	var im item.Item
	var u user.User

	for i:=range us{
		if us[i].Name==un{
			u=us[i]
		}
	}

	for i:=range it{
		if it[i].Brand==b && it[i].Category==ic{
			im = it[i]
		}
	}

	return cart.AddToCart(c, &u, im, q)
}

func Remove(c []cart.Cart, us []user.User, it []item.Item, un string, ic string, b  string) []cart.Cart {
	var im item.Item
	var u user.User

	for i:=range us{
		if us[i].Name==un{
			u=us[i]
		}
	}

	for i:=range it{
		if it[i].Brand==b && it[i].Category==ic{
			im = it[i]
		}
	}

	return cart.RemoveCart(c, &u, im)
}

func Checkout(username string, c []cart.Cart, u []user.User, inv []inventory2.Inventory) ([]cart.Cart, []inventory2.Inventory, []user.User){
	totalAmount:=0.0
	for i:=range u{
		if u[i].Name==username{
			totalAmount=u[i].WalletAmount
			break
		}
	}

	cartAmount:= cart.CartAmount(c, username)

	if cartAmount>totalAmount{
		fmt.Println("you cant checkout your balance is low")
		return c,inv,u
	}

	val:=[]int{}

	for i:=range c{
		if c[i].User.Name==username{
			for j:=range c[i].Item{
				for k:=range inv{
					if inv[k].Item.Category==c[i].Item[j].Category && inv[k].Item.Brand==c[i].Item[j].Brand{
						val=append(val, k)
					}
				}
			}
			c[i].Item=[]item.Item{}
			c[i].Quantity=[]int{}
		}
	}

	for i:=0;i<len(val);i++{
		inv= inventory2.RemoveInv(inv, val[i])
	}

	return c,inv,u

}





func main()  {
	var items []item.Item
	var users []user.User
	var inventory []inventory2.Inventory
	var carts []cart.Cart

	for{
		var inp string
		fmt.Scan(&inp)
		switch inp {
		case "CIT":
			var (
				c string
				b string
				p float64
			)
			fmt.Scan(&c, &b, &p)
			i := item.Create(c, b, p)
			items = append(items, *i)
		case "AINV":
			var (
				ic string
				b  string
				q  int
			)

			fmt.Scan(&ic, &b, &q)
			inventory = AddInv(inventory, items, ic, b, q)

		case "AUS":
			var (
				name string
				w    float64
			)
			fmt.Scan(&name, &w)
			users = append(users, *user.AddUser(name, w))
		case "ATC":
			var (
				un string
				ic string
				b  string
				q  int
			)
			fmt.Scan(&un, &ic, &b, &q)

		case "UC":
			var (
				un string
				ic string
				b  string
				q  int
			)
			fmt.Scan(&un, &ic, &b, &q)
			carts = ADC(carts, users, items, un, ic, b, q)
		case "RE":
			var (
				un string
				ic string
				b  string
			)
			fmt.Scan(&un, &ic, &b)
			carts = Remove(carts, users, items, un, ic, b)
		case "GC":
			var name string
			fmt.Scan(&name)
			cart.GetCart(carts, name)
		case "CO":
			var un string
			fmt.Scan(&un)
			carts, inventory, users=Checkout(un, carts, users, inventory)



		default:
			break
		}

	}

}