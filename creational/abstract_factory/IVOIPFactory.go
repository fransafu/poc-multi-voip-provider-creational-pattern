package main

import "fmt"

type IVOIPFactory interface {
	makeCall() ICall
	makeConference() IConference
}

func getVOIPFactory(provider string) (IVOIPFactory, error) {
	switch provider {
	case "twilio":
		return &twilio{}, nil
	case "plivo":
		return &plivo{}, nil
	default:
		return nil, fmt.Errorf("Wrong provider type passed")
	}
}
