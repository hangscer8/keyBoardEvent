package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	hook "github.com/robotn/gohook"
	"github.com/vcaesar/keycode"
)

func main() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		chanHook := hook.Start()
		for ev := range chanHook {
			if ev.Kind != hook.KeyUp {
				continue
			}
			for keyStr, keyCode := range keycode.Keycode {
				if keyCode == ev.Keycode {
					result, _ := json.Marshal(KeyEvent{Kind: "KeyUp", KeyName: keyStr})
					fmt.Println(string(result))
				}
			}
		}
	}()
	<-signalCh
	hook.End()
	fmt.Println("process terminating")
}

type KeyEvent struct {
	Kind    string
	KeyName string
}
