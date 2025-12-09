// A pacakage containing functions for interacting with the frontend
package bridge

import (
	"context"
	"time"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

type ConsoleMessage struct {
	Content string `json:"content"`
	RxTime  string `json:"rxTime"`
	Size    int    `json:"size"`
	Sender  string `json:"sender"`
}

func makeConsoleMessage(msg string, sender string) ConsoleMessage {
	return ConsoleMessage{
		Content: msg,
		RxTime:  time.Now().Format("15:04:05.000"),
		Size:    len(msg),
		Sender:  sender,
	}
}

func PresentMessage(ctx context.Context, msg, sender string) {
	cmsg := makeConsoleMessage(msg, sender)
	rt.EventsEmit(ctx, "console-message", cmsg)
}
