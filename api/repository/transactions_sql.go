package repository

import (
	"fmt"

	"github.com/dakicka/tradingApp/api/entity"
	"github.com/dakicka/tradingApp/api/package/db"
)

// TransactionSQL wraps the SQL DB and implements the required operations.
type TransactionSQL struct {
	db.GormDB
}

// NewTransactionSQLRepo instanciates and returns a new transaction repository.
func NewTransactionSQLRepo(db *db.GormDB) Transactions {
	return TransactionSQL{*db}
}

// Create inserts the transaction into the PSQL db and returns the transaction
// including generated values like id and timestamps
func (r TransactionSQL) Create(tx entity.Transaction) (entity.Transaction, error) {

	result := r.DB.Create(&tx)

	if result.Error != nil {
		fmt.Println(result.Error)
		return entity.Transaction{}, result.Error
	}

	return tx, nil
}

func (r TransactionSQL) GetAllByUserId(userId uint) ([]entity.Transaction, error) {

	transactions := []entity.Transaction{}
	result := r.DB.Find(&transactions, "user_id = ?", userId)

	if result.Error != nil {
		return []entity.Transaction{}, result.Error
	}

	return transactions, nil
}
