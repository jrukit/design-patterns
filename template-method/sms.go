package main

import "fmt"

type Sms struct{}

func (s *Sms) getRandomOTP(otpLen int) string {
	randOTP := "1234"
	fmt.Printf("Sms: generating random otp %s\n", randOTP)

	return randOTP
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("Sms: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "Sms OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("Sms: sending Sms: %s\n", message)
	return nil
}
