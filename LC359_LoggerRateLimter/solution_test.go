package lc359

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewLogger()

	// Test: first-time message should be printed
	if ok := logger.ShouldPrintMessage(1, "msg1"); !ok {
		t.Errorf("Expected true, got %v", ok)
	}

	// Test: within 10 seconds, should not print
	if ok := logger.ShouldPrintMessage(2, "msg1"); ok {
		t.Errorf("Expected false, got %v", ok)
	}

	// Test: after 10 seconds, should print
	if ok := logger.ShouldPrintMessage(11, "msg1"); !ok {
		t.Errorf("Expected true, got %v", ok)
	}

	// Test: new message, should print
	if ok := logger.ShouldPrintMessage(12, "msg2"); !ok {
		t.Errorf("Expected true, got %v", ok)
	}

	// Test: new message, should print and old messages should cleaned
	if ok := logger.ShouldPrintMessage(23, "msg3"); !ok {
		t.Errorf("Expected true, got %v", ok)
	}
	if len(logger.logMap) != 1 {
		t.Errorf("Expected one log, got %v", len(logger.logMap))
	}

}
