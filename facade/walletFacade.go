package main

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func newWalletFacade(accountName string, code int) *WalletFacade {
	return &WalletFacade{
		account:      newAccount(accountName),
		wallet:       newWallet(),
		securityCode: newSecurityCode(code),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
}

func (w *WalletFacade) addMoneyToWallet(accountId string, securityCode, amount int) error {
	err := w.account.checkAccount(accountId)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountId, "credit", amount)
	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountId string, securityCode, amount int) error {
	err := w.account.checkAccount(accountId)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	w.wallet.debitBalance(amount)
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountId, "debit", amount)
	return nil
}
