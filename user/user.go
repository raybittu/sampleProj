package user

type User struct {
	Name string
	WalletAmount float64
}

func AddUser(name string, walletAmount float64) *User {
	return &User{Name: name, WalletAmount: walletAmount}
}