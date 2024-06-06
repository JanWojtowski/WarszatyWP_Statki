package utility

import (
	"context"
	"github.com/grupawp/termloop"
)

type InputField struct {
	*termloop.Text
	buffer    string
	ch        chan string
	maxlenght int
}

func NewInputField(x, y int, ch chan string, maxlenght int) *InputField {
	return &InputField{termloop.NewText(x, y, "", termloop.ColorWhite, termloop.ColorBlack), "", ch, maxlenght}
}

func (i *InputField) Tick(event termloop.Event) {
	if event.Type == termloop.EventKey {
		switch event.Key {
		case termloop.KeyEnter:
			i.processInput(i.buffer)
			i.buffer = ""
		case termloop.KeySpace:
			i.buffer += " "
		case termloop.KeyBackspace, termloop.KeyBackspace2:
			if len(i.buffer) > 0 {
				i.buffer = i.buffer[:len(i.buffer)-1]
			}
		default:
			if len(i.buffer) <= i.maxlenght {
				i.buffer += string(event.Ch)
			}
		}
		i.SetText(i.buffer)
	}
}

func (i *InputField) processInput(input string) {
	select {
	case i.ch <- i.buffer:
	default:
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

	inputField := NewInputField(10, 10, temp.ch, 10)
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
