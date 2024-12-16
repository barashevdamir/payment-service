package payments

import (
	"context"
	"database/sql"
)

type PaymentMethod string

const (
	Yookassa PaymentMethod = "yookassa"
	Telegram               = "telegram"
)

type Transaction struct {
	Amount        float64           `json:"amount"`
	Description   string            `json:"description"`
	Recipient     map[string]string `json:"recipient"`
	PaymentToken  int64             `json:"payment_token"`
	PaymentMethod PaymentMethod     `json:"payment_method"`
	Metadata      map[string]string `json:"metadata"`
}

type TransactionRepository interface {
	GetTransactionByToken(ctx context.Context, transaction *Transaction) error
}

type transactionRepo struct {
	db *sql.DB
}

func (t transactionRepo) GetTransactionByToken(paymentToken string) (Transaction, error) {
	ctx := context.Background()
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return Transaction{}, err
	}
	var transaction = Transaction{}
	query := "select amount, description, recipient, payment_method, metadata FROM transactions where payment_token=$1"
	err = tx.QueryRowContext(ctx, query, paymentToken).Scan(
		&transaction.Amount, &transaction.Description, &transaction.Recipient, &transaction.PaymentMethod, &transaction.Metadata)
	if err != nil {
		tx.Rollback()
		return Transaction{}, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return Transaction{}, err
	}
	return transaction, err
}
