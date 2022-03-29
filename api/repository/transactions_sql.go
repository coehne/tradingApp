package repository

import (
	"fmt"

	"github.com/dakicka/tradingApp/api/db"
	"github.com/dakicka/tradingApp/api/entity"
)

// TransactionSQL wraps the SQL DB and implements the required operations.
type TransactionSQL struct {
	db.GormDB
}

// NewTransactionSQLRepo instanciates and returns a new transaction repository.
func NewTransactionSQLRepo(db *db.GormDB) Transactions {
	return TransactionSQL{*db}
}

func (r TransactionSQL) CreateTransaction(tx entity.Transaction) (entity.Transaction, error) {
	// Insert into DB
	result := r.DB.Create(&tx)

	// Check for errors during insertion
	if result.Error != nil {
		fmt.Println(result.Error)
		return entity.Transaction{}, result.Error
	}

	return tx, nil
}

func (r TransactionSQL) GetAllForUserId(userId uint) ([]entity.Transaction, error) {

	transactions := []entity.Transaction{}
	result := r.DB.Find(&transactions, "user_id = ?", userId)

	// Check for errors during query
	if result.Error != nil {
		return []entity.Transaction{}, result.Error
	}

	return transactions, nil
}
