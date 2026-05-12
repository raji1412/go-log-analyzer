package analyzer

import (
	"os"
	"testing"
)

func TestLogCounter(t *testing.T) {
	logContent := `INFO User logged in
ERROR Database connection failed
INFO Request received
WARNING Disk almost full
ERROR Timeout occurred`

	err := os.WriteFile("logContent.txt", []byte(logContent), 0777)
	if err != nil {
		t.Fatalf("Error writing to logContent.txt: %v", err)
	}

	result := LogCounter("logContent.txt", "")
	if len(result) < 1 {
		t.Fatalf("Log counter result is empty.")
	}

	if result["info"] != 2 {
		t.Fatalf("Log counter result is wrong for info.")
	}

	if result["error"] != 2 {
		t.Fatalf("Log counter result is wrong for error.")
	}

	if result["warning"] != 1 {
		t.Fatalf("Log counter result is wrong for warning.")
	}
}

func TestLogCounter_FilterLevel(t *testing.T) {

	logContent := `INFO Start
ERROR Crash
INFO Request
ERROR Timeout`

	fileName := "test_logs_filter.txt"

	err := os.WriteFile(fileName, []byte(logContent), 0644)
	if err != nil {
		t.Fatalf("failed to create test log file: %v", err)
	}

	defer os.Remove(fileName)

	result := LogCounter(fileName, "error")

	if result["error"] != 2 {
		t.Errorf("expected 2 error logs, got %d", result["error"])
	}

	if len(result) != 1 {
		t.Errorf("expected only error logs in result")
	}
}
