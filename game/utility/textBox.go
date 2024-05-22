package utility

import (
	"context"
	"github.com/grupawp/termloop"
)

type InputField struct {
	*termloop.Text
	buffer string
	ch     chan string
}

func NewInputField(x, y int, ch chan string) *InputField {
	return &InputField{termloop.NewText(x, y, "", termloop.ColorWhite, termloop.ColorBlack), "", ch}
}

func (i *InputField) Tick(event termloop.Event) {
	if event.Type == termloop.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case termloop.KeyEnter:
			// User has pressed enter, process buffer and clear it
			i.processInput(i.buffer)
			i.buffer = ""
		case termloop.KeySpace:
			// User has pressed space, add a space to the buffer
			i.buffer += " "
		case termloop.KeyBackspace, termloop.KeyBackspace2:
			// User has pressed backspace, remove the last character from the buffer
			if len(i.buffer) > 0 {
				i.buffer = i.buffer[:len(i.buffer)-1]
			}
		default:
			// User has pressed a different key, append it to the buffer
			i.buffer += string(event.Ch)
		}
		i.SetText(i.buffer) // Update the visual text field
	}
}

func (i *InputField) processInput(input string) {
	select {
	case i.ch <- i.buffer:
	default:
		// drop
	}
}

func (i *InputField) Listen(ctx context.Context) string {
	select {
	case s := <-i.ch:
		return s
	case <-ctx.Done():
		return ""
	}
}

type channel struct {
	ch chan string
}

func TextTest() {
	game := termloop.NewGame()
	temp := channel{make(chan string)}
	level := termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorBlack,
	})

	inputField := NewInputField(10, 10, temp.ch)
	level.AddEntity(inputField)

	game.Screen().SetLevel(level)
	go func() {
		for {
			char := inputField.Listen(context.TODO())

			panic(char)

		}
	}()
	game.Start(context.TODO())
}
