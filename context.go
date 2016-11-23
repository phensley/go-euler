package euler

import (
	"fmt"
	"io"
	"time"
)

// Context holds some common state for the solution and
// collects the answer.
type Context struct {
	id          string
	exists      bool
	description string
	out         io.Writer
	known       string
	answered    bool
	answer      string
	elapsed     time.Duration
}

// SetAnswer sets the answer to the problem
func (c *Context) SetAnswer(a string) {
	c.answered = true
	c.answer = a
}

// Exists indicates whether a solution exists
func (c *Context) Exists() bool {
	return c.exists
}

// ID returns the problem's identifier
func (c *Context) ID() string {
	return c.id
}

// Description returns the problem's description
func (c *Context) Description() string {
	return c.description
}

// Answer is the (hopefully correct) answer to the problem
func (c *Context) Answer() string {
	return c.answer
}

// IsAnswered indiciates an answer was provided
func (c *Context) IsAnswered() bool {
	return c.answered
}

// IsSolved indicates if the problem has been solved
func (c *Context) IsSolved() bool {
	return c.known != "" && c.answer == c.known
}

// Elapsed is the time taken to produce the answer
func (c *Context) Elapsed() time.Duration {
	return c.elapsed
}

func (c *Context) String() string {
	return fmt.Sprintf("id=%v exists=%v desc=%#v known=%#v answered=%v answer=%#v elapsed=%.3f",
		c.id,
		c.exists,
		c.description,
		c.known,
		c.answered,
		c.answer,
		c.elapsed.Seconds(),
	)
}
