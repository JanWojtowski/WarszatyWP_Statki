package utility

import (
	"context"
	tl "github.com/grupawp/termloop"
)

type button struct {
	*tl.Rectangle
	coord string
	ch    chan string
}

func NewRectangle(rec *tl.Rectangle) *button {
	return &button{Rectangle: rec}
}

func NewClickableRectangle(rec *tl.Rectangle, coord string, ch chan string) *button {
	return &button{Rectangle: rec, coord: coord, ch: ch}
}

func (c *button) Tick(e tl.Event) {
	if c.ch == nil || c.coord == "" {
		return
	}

	switch e.Key {
	case tl.MouseLeft:
		c.processClick(e)
	}

}

func (c *button) processClick(e tl.Event) {
	x, y := c.Position()
	w, h := c.Size()
	if e.MouseX >= x && e.MouseY >= y && e.MouseX <= (x+w) && e.MouseY <= (y+h) {
		select {
		case c.ch <- c.coord:
		default:
			// drop
		}
	}

}

func (b *button) Listen(ctx context.Context) string {
	select {
	case s := <-b.ch:
		return s
	case <-ctx.Done():
		return ""
	}
}
