package gmesh

import (
	"github.com/briandowns/spinner"
	"time"
)

const (
	// Version -> Muss noch in ne Config
	Version = "0.1.0-alpha"
	ApiUrl  = "https://gme.sh/"
	Prefix  = "https://gme.sh/"
)

type CLI struct {
	Pipe string
}

func New() *CLI {
	return &CLI{}
}

func (c *CLI) Run() (err error) {
	// check pipe
	err = c.ReadPipe()
	if err != nil {
		return
	}
	// show help
	return c.ParseArgs()
}

func newSpinner() *spinner.Spinner {
	return spinner.New(spinner.CharSets[9], 100*time.Millisecond)
}
