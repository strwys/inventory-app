package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateTransactionNumber(trxType string) string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000000)
	timestamp := time.Now().Format("20060102150405")
	transactionNumber := fmt.Sprintf("TX/%s/%s%d", trxType, timestamp, randomNumber)
	return transactionNumber
}
