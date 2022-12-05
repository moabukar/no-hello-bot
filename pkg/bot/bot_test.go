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
