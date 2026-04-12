package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleHello(t *testing.T) {
	resp := handleHello(events.LambdaFunctionURLRequest{})

	if resp.StatusCode != 200 {
		t.Fatalf("status code = %d, want %d", resp.StatusCode, 200)
	}
	if resp.Body != "Hello, World!" {
		t.Fatalf("body = %q, want %q", resp.Body, "Hello, World!")
	}
}

func TestHandleKey(t *testing.T) {
	resp := handleKey(events.LambdaFunctionURLRequest{
		Headers: map[string]string{
			"Authorization": "abc123",
		},
	})

	if resp.StatusCode != 200 {
		t.Fatalf("status code = %d, want %d", resp.StatusCode, 200)
	}
	if resp.Body != "abc123" {
		t.Fatalf("body = %q, want %q", resp.Body, "abc123")
	}
}

func TestHandleNotFound(t *testing.T) {
	resp := handleNotFound(events.LambdaFunctionURLRequest{})

	if resp.StatusCode != 404 {
		t.Fatalf("status code = %d, want %d", resp.StatusCode, 404)
	}
	if resp.Body != "Not Found" {
		t.Fatalf("body = %q, want %q", resp.Body, "Not Found")
	}
}
