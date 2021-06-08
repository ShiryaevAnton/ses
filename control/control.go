package control

import (
	"strings"
	"time"

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
		robotgo.KeyTap(key)
		time.Sleep(300 * time.Millisecond)
	}
}

func (c *Control) TypeString(text string) {

	robotgo.TypeStrDelay(text, 1)
}
