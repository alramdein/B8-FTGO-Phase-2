package main

// Single Responsibility
// -- Bad
type OrderTransactionService struct{}

func (o *OrderTransactionService) CreateOrder() {}

func (o *OrderTransactionService) RecordTrasaction() {}

// -- Good
type OrderService struct{}

func (o *OrderService) CreateOrder() {}

type TransactionService struct{}

func (t *TransactionService) RecordTrasaction() {}

func main1() {
	order := OrderService{}
	order.CreateOrder()

	transaction := TransactionService{}
	transaction.RecordTrasaction()
}
