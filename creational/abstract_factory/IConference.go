package main

type IConference interface {
	call()
	addParticipant()
	removeParticipant()
	mute()
	unmute()
	hold()
	unhold()
	playMusic()
	stopMusic()
	endCall()
}

type conference struct{}

func (c *conference) call() {

}

func (c *conference) addParticipant() {

}

func (c *conference) removeParticipant() {

}

func (c *conference) mute() {

}

func (c *conference) unmute() {

}

func (c *conference) hold() {

}

func (c *conference) unhold() {

}

func (c *conference) playMusic() {

}

func (c *conference) stopMusic() {

}

func (c *conference) endCall() {

}
