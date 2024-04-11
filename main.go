package main

import (
	"fmt"
	"log"
	"os/exec"

	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

var keys = map[hotkey.Key]string{
	hotkey.KeyQ: "/Applications/Kitty.app/",
	hotkey.KeyW: "/Applications/Google Chrome.app/",
	hotkey.KeyE: "/Applications/Slack.app/",
	hotkey.KeyR: "/Applications/Spotify.app/",
	hotkey.KeyZ: "/Applications/zoom.us.app/",
}

func main() { mainthread.Init(fn) }

func fn() {
	for key, app := range keys {
		hk := hotkey.New([]hotkey.Modifier{hotkey.ModCmd, hotkey.ModShift}, key)
		if err := hk.Register(); err != nil {
			panic(err)
		}
		fmt.Println("registered", app)

		go func(app string) {
			for range hk.Keydown() {
				_ = exec.Command("open", app).Run()
				log.Println("Switching to", app)
			}
		}(app)
	}

	select {}
}
