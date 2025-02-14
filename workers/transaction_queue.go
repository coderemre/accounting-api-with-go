package workers

import (
	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/services"
	"sync"
	"sync/atomic"
)

var transactionCounter int64

type TransactionQueue struct {
	queue   chan *models.Transaction
	workers int
}

func NewTransactionQueue(workerCount int) *TransactionQueue {
	return &TransactionQueue{
		queue:   make(chan *models.Transaction, 100),
		workers: workerCount,
	}
}

func (tq *TransactionQueue) StartProcessing(service services.TransactionService) {
	var wg sync.WaitGroup
	for i := 0; i < tq.workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for tx := range tq.queue {
				if err := service.ProcessTransaction(tx); err == nil {
					atomic.AddInt64(&transactionCounter, 1)
				}
			}
		}()
	}
	wg.Wait()
}

func (tq *TransactionQueue) AddTransaction(tx *models.Transaction) {
	tq.queue <- tx
}

func GetProcessedTransactionCount() int64 {
	return atomic.LoadInt64(&transactionCounter)
}