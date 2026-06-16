package im

import (
	"context"
	"testing"
)

func TestDetachMessageContextIgnoresCanceledParent(t *testing.T) {
	type contextKey string
	key := contextKey("trace")

	parent, cancel := context.WithCancel(context.WithValue(context.Background(), key, "keep-me"))
	cancel()

	got := detachMessageContext(parent)
	if err := got.Err(); err != nil {
		t.Fatalf("detached context should not be canceled, got err=%v", err)
	}
	if value := got.Value(key); value != "keep-me" {
		t.Fatalf("detached context should preserve values, got %v", value)
	}
}

func TestDetachMessageContextNil(t *testing.T) {
	got := detachMessageContext(nil)
	if got == nil {
		t.Fatal("detached nil context should return a background context")
	}
	if err := got.Err(); err != nil {
		t.Fatalf("background context should not be canceled, got err=%v", err)
	}
}
