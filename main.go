package main

import (
	"time"

	"github.com/micmonay/keybd_event"
	"golang.org/x/sys/windows"
)

var (
	user32               = windows.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

func main() {
	for {
		val, _, _ := procGetAsyncKeyState.Call(uintptr(192)) // Tilde / backqoute

		if val&0x1 == 0 {
			continue
		}

		go runMacro()

		time.Sleep(1 * time.Millisecond)
	}
}

func runMacro() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	kb.SetKeys(keybd_event.VK_1, keybd_event.VK_2, keybd_event.VK_3, keybd_event.VK_4, keybd_event.VK_5)

	err = kb.Launching()
	if err != nil {
		panic(err)
	}
}
