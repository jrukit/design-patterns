package main

type IOtp interface {
	getRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOTP(otpLen int) error {
	otp := o.iOtp.getRandomOTP(otpLen)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}

	return nil
}
