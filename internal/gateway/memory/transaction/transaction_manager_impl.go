package transaction

import (
	"context"
	"log/slog"

	"github.com/Haya372/go-samples/internal/common/transaction"
)

type transactionManagerImpl struct {
}

func (txm *transactionManagerImpl) Do(ctx context.Context, f func(ctx context.Context) error) error {
	slog.Info("start transaction")
	if err := f(ctx); err != nil {
		slog.Info("rollback transaction")
		return err
	}
	slog.Info("commit transaction")
	return nil
}

func NewInMemoryTransactionManager() transaction.TransactionManager {
	return &transactionManagerImpl{}
}
