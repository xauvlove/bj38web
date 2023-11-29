package main

import (
	"encoding/json"
	"fmt"

	"github.com/afocus/captcha"
)

func main() {

	image := captcha.NewImage(128, 64)
	marshal1, _ := json.Marshal(image)
	s3 := string(marshal1)
	fmt.Printf("%v", s3)

	//cap := captcha.New()
	//cap.SetFont("comic.ttf")
	//cap.SetSize(128, 64)
	//cap.SetDisturbance(captcha.NORMAL)
	//cap.SetFrontColor(color.RGBA{132, 63, 0, 53})
	//cap.SetBkgColor(color.RGBA{42, 163, 110, 99})
	//img, s := cap.Create(4, captcha.UPPER)
	//marshal, _ := json.Marshal(img)
	//s2 := string(marshal)
	//fmt.Printf("%v", s2)
	//println(s)
}
