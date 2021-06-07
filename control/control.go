package control

import (
	"strings"

	"github.com/go-vgo/robotgo"
)

type Control struct{}

func NewControl() *Control {
	c := Control{}
	return &c
}

func (c *Control) Press(keySequence string) {

	keySequenceArray := strings.Split(keySequence, " ")

	for _, key := range keySequenceArray {

		if strings.Contains(key, "+") {
			subKeys := strings.Split(key, "+")
			robotgo.KeyTap(subKeys[0], subKeys[1])
		} else {
			robotgo.KeyTap(key)
		}
		robotgo.Sleep(1)
	}
}

func (c *Control) TypeString(text string) {

	robotgo.TypeStrDelay(text, 1)
}
