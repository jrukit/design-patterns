package main

/*
 Let’s consider the example of One Time Password (Otp) functionality. There are different ways that the Otp can be delivered to a user (SMS, email, etc.). But irrespective whether it’s an SMS or email Otp, the entire Otp process is the same:

1. Generate a random n digit number.
2. Save this number in the cache for later verification.
3. Prepare the content.
4. Send the notification.

Any new Otp types that will be introduced in the future most likely still go through the above steps.

So, we have a scenario where the steps of a particular operation are the same, but these steps’ implementation may differ. This is an appropriate situation to consider using the Template Method pattern.

First, we define a base template algorithm that consists of a fixed number of methods. That’ll be our template method. We will then implement each of the step methods, but leave the template method unchanged.

*/

func main() {
	/*
		Otp := Otp{}

		smsOtp := &sms{
			Otp: Otp,
		}

		smsOtp.genAndSentOtp(smsOtp, 4)

		emailOtp := &emailOtp{
			Otp: Otp,
		}
		emailOtp.genAndSendOtp(emailOtp, 4)
		fmt.Scanln()
	*/

	smsOtp := &Sms{}
	o := Otp{
		iOtp: smsOtp,
	}
	o.genAndSendOTP(4)

	emailOtp := &Email{}
	o = Otp{
		iOtp: emailOtp,
	}
	o.genAndSendOTP(4)
}
