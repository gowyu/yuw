package yuw

import (
	"fmt"
	_ "github.com/gowyu/yuw/modules"
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
