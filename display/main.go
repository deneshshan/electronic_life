package display

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

//func main() {
//cmd := exec.Command("stty", "size")
//cmd.Stdin = os.Stdin
//out, err := cmd.Output()
//fmt.Printf("out: %#v\n", string(out))
//fmt.Printf("err: %#v\n", err)
//if err != nil {
//log.Fatal(err)
//}
//}

func getWindowSize()

func NewDisplay() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: 'v',
	})
	//level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))

	//Set the character at position (0, 0) on the entity.
	game.Screen().SetLevel(level)
	dat, _ := ioutil.ReadFile("lorry.txt")
	e := tl.NewEntityFromCanvas(1, 1, tl.CanvasFromString(string(dat)))
	game.Screen().AddEntity(e)

	game.Start()
}
