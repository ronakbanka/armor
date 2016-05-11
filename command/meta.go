package command

import (
	"io"
	"log"
)

// Meta stores meta-option used in all subcommands
type Meta struct {
	// Name is executable name of command
	Name string

	// Debug shows debug logs
	Debug bool

	OutStream, ErrStream io.Writer
}

func (m *Meta) Debugf(format string, args ...interface{}) {
	if m.Debug {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}
