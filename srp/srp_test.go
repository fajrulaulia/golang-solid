package srp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlowTransactionAndUserSi(t *testing.T) {
	var UserSvc UserService
	var TrxSvc TransactionService

	// Proses tambakan user
	UserSvc.AddUser(&User{ID: 1, Name: "Fajrul", Balance: 150})
	UserSvc.AddUser(&User{ID: 2, Name: "Widia", Balance: 100})

	// Widia Check Data Sendiri
	userWidia := UserSvc.GetUserByID(2)
	assert.Equal(t, userWidia.Balance, float64(100), "they should be equal", "balance widia di awal hanya 100")
	assert.Equal(t, userWidia.Name, "Widia", "they should be equal", "nama widia")
	assert.Equal(t, userWidia.ID, 2, "they should be equal", "id nama widia")

	// Widia melakukan Deposit
	TrxSvc.Deposit(userWidia, 50)

	// Widia check balance
	userWidia = UserSvc.GetUserByID(2)
	assert.Equal(t, userWidia.Balance, float64(150), "they should be equal", "Balance Widia telah mencapai 150")

	// Widia Witdrawal
	TrxSvc.Withdraw(userWidia, 120)
	userWidia = UserSvc.GetUserByID(2)
	assert.Equal(t, userWidia.Balance, float64(30), "they should be equal", "Balance Widia telah setelah di WD menjadi 30")

}
