package bot_test

import (
	"testing"

	"github.com/moabukar/no-hello-bot/pkg/bot"
)

func TestGreet(t *testing.T) {
	expected := "Hello"
	actual := bot.Greet()
	if actual != expected {
		t.Errorf("Greet() got %v want %v", actual, expected)
	}
}

func TestGetResponse(t *testing.T) {
	input := "hello"
	expected := "Hey, how can I help you?"
	actual := bot.GetResponse(input)
	if actual != expected {
		t.Errorf("GetResponse() got %v want %v", actual, expected)
	}
}

func TestHandleMessage(t *testing.T) {
	input := "hey"
	expected := "Hi there!"
	actual := bot.HandleMessage(input)
	if actual != expected {
		t.Errorf("HandleMessage() got %v want %v", actual, expected)
	}
}
