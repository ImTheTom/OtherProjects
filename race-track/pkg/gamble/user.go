package gamble

type User struct {
	ID      int
	Name    string
	Balance float64
}

const StartingBalance = 100.0

var users []*User

var currentUserID int

func CreateNewUser(name string) *User {
	usr := &User{
		ID:      currentUserID,
		Name:    name,
		Balance: StartingBalance,
	}
	users = append(users, usr)

	return usr
}

func AdjustUserBalance(id int, amount float64) {
	for _, v := range users {
		if v.ID == id {
			v.Balance += amount
		}
	}
}
