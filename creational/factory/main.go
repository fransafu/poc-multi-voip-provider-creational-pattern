package main

func main() {
	twilioFactory, _ := getVOIPFactory("twilio")
	plivoFactory, _ := getVOIPFactory("plivo")

	twilioCall := makeCall()
	plivoCall := makeCall()

	twilioConference := makeConference()
	plivoConference := makeConference()

	twilioCall.call()
	plivoCall.call()

	twilioConference.call()
	plivoConference.call()
}
