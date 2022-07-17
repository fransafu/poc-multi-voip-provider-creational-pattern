package main

type ICall interface {
	call()
	mute()
	unmute()
	hold()
	unhold()
	playMusic()
	stopMusic()
	endCall()
}

type call struct{}

func (c *call) call() {

}

func (c *call) mute() {

}

func (c *call) unmute() {

}

func (c *call) hold() {

}

func (c *call) unhold() {

}

func (c *call) playMusic() {

}

func (c *call) stopMusic() {

}

func (c *call) endCall() {

}
