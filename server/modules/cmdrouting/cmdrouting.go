package cmdrouting

import (
	"bufio"
	"fmt"
	"io"
)

//Listener Hanle commands from Reader
type Listener map[string]func()

//AddHandler is routing function
func (l Listener) AddHandler(cmd string, handleFunc func()) {
	l[cmd] = handleFunc
}

//Start CmdListener
func (l Listener) Start(r io.Reader) {
	fmt.Println(l)
	go l.start(r)
}

func (l Listener) start(r io.Reader) {

	var cmd string

	input := bufio.NewScanner(r)

	for input.Scan() {
		cmd = input.Text()

		handler, ok := l[cmd]

		if !ok {
			continue
		}

		handler()
	}

}
