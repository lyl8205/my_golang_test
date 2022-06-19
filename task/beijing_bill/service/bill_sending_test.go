package service

import "testing"

func TestNewBillSending_Send(t *testing.T) {
	limit := 10000
	NewBillSending().Send(limit)
}
func TestNewBillSending_SendTest(t *testing.T) {
	NewBillSending().SendTest()
}
