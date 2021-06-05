package cart

import (
	"fmt"
	item2 "likeMindProj/item"
	"likeMindProj/user"
)

type Cart struct {
	User *user.User
	Item []item2.Item
	Quantity []int
}

func AddToCart(cart []Cart, u *user.User, item item2.Item, quantity int) []Cart {
	flag:=false
	for i:=range cart{
		if cart[i].User==u{
			f2:=false
			for j,v:=range cart[i].Item{
				if v==item{
					cart[i].Quantity[j]+=quantity
					f2=true
				}
			}
			if !f2{
				cart[i].Item=append(cart[i].Item, item)
				cart[i].Quantity=append(cart[i].Quantity, quantity)
			}
			flag=true
		}
	}
	if !flag{
		q:=[]int{quantity}
		it:=[]item2.Item{item}

		c:= Cart{
			User:     u,
			Item:     it,
			Quantity: q,
		}
		cart=append(cart, c)
	}

	return cart
}

func removeInt(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}


func removeItem(slice []item2.Item, s int) []item2.Item {
	return append(slice[:s], slice[s+1:]...)
}

func RemoveCart(cart []Cart, u *user.User, item item2.Item) []Cart {
	for i:=range cart{
		if cart[i].User==u{
			for j:=range cart[i].Item{
				if cart[i].Item[j]==item{
					cart[i].Item = removeItem(cart[i].Item, j)
					cart[i].Quantity= removeInt(cart[i].Quantity, j)
				}
			}
		}
	}
	return cart
}

func GetCart(c []Cart, userName string){
	for i:=range c{
		if c[i].User.Name==userName{
			fmt.Println(c[i].User)
			fmt.Println(c[i].Item)
			fmt.Println(c[i].Quantity)
		}
	}
}

func CartAmount(c []Cart, username string) float64{
	var sum float64
	for i:=range c{
		if c[i].User.Name==username{
			for j:=range c[i].Item{
				sum=sum+float64(c[i].Quantity[j])*c[i].Item[j].Price
			}
		}
	}
	return sum
}
