package test

import (
	"context"
	"testing"

	"github.com/SuCicada/apprise-sdk-go/apprise"
)

func TestNotifier(t *testing.T) {
	notifier := apprise.Notifier{
		URL: "http://localhost:8080",
	}

	notifier.Send(context.Background(), &apprise.Message{
		Title: "title",
		Body:  "body",
		Tag:   "alert",
	})
}
