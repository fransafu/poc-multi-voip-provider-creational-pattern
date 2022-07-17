package main

type twilio struct {
}

func (t *twilio) makeCall() ICall {
	return &twilioCall{
		call: call{},
	}
}

func (t *twilio) makeConference() IConference {
	return &twilioConference{
		conference: conference{},
	}
}
