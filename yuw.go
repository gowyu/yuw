package yuw

import (
	"fmt"
)

type Engine struct {

}

func New() *Engine {
	return &Engine {

	}
}

func Run() (Y *Engine) {
	Y = New()
	Y.Run()

	return
}

func (yuw *Engine) Run() {
	fmt.Println("Yuw Success!")
}
