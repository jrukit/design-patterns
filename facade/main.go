package main

import (
	"fmt"
	"log"
)

// Conceptual Example
// It’s easy to underestimate the complexities that happen behind the scenes when you order a pizza using your credit card. There are dozens of subsystems that are acting in this process. Here’s just a shortlist of them:

// Check account
// Check security PIN
// Credit/debit balance
// Make ledger entry
// Send notification
// In a complex system like this, it’s easy to get lost and easy to break stuff if you’re doing something wrong. That’s why there’s a concept of the Facade pattern: a thing that lets the client work with dozens of components using a simple interface. The client only needs to enter the card details, the security pin, the amount to pay, and the operation type. The Facade directs further communications with various components without exposing the client to internal complexities.

// without Facade

func withoutFacade() {
	account := newAccount("abc")
	securityCode := newSecurityCode(1234)
	wallet := newWallet()
	notification := Notification{}
	ledger := Ledger{}

	// add money to wallet
	err := account.checkAccount("abc")
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	err = securityCode.checkCode(1234)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	wallet.creditBalance(10)
	notification.sendWalletCreditNotification()
	ledger.makeEntry("abc", "credit", 10)

	fmt.Println()

	// deduct money from wallet
	err = account.checkAccount("abc")
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	err = securityCode.checkCode(1234)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	wallet.debitBalance(5)
	notification.sendWalletDebitNotification()
	ledger.makeEntry("abc", "debit", 5)
}

func withFacade() {
	walletFacade := newWalletFacade("abcd", 1111)

	err := walletFacade.addMoneyToWallet("abcd", 1111, 100)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	err = walletFacade.deductMoneyFromWallet("abcd", 1111, 20)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}

func main() {
	// withoutFacade()

	withFacade()
}
