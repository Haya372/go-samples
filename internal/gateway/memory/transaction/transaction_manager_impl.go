package transaction

import (
	"context"
	"log/slog"

	"github.com/Haya372/go-samples/internal/common/transaction"
)

type transactionManagerImpl struct {
}

func (txm *transactionManagerImpl) Do(ctx context.Context, f func(ctx context.Context) error) error {
	slog.Debug("start transaction")
	if err := f(ctx); err == nil {
		slog.Debug("commit transaction")
		return err
	} else {
		slog.Debug("rollback transaction")
	}
	return nil
}

func NewInMemoryTransactionManager() transaction.TransactionManager {
	return &transactionManagerImpl{}
}
