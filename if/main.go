package main

import (
	"fmt"
	"log"
	"sleep-as-a-serive/captcha"
)

func main() {
	c := captcha.NewCaptcha(0, 1, 1, 1)
	if err := c.Validate(); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print(c.GenerateCaptcha())
}
