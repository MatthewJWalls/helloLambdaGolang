package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// this is how they laid out testing in the official example

func TestHandler(t *testing.T) {
	
	tests := []struct {
		request Request
		expect  string
		err     error
	}{
		{
			request: Request{Message: "what"},
			expect: "you sent the message what",
			err: nil,
		},
	}
	
	for _, test := range tests {
		response, err := Handler(test.request) 
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Message)
	}
	
}


// this is how I'd lay it out

func TestBetter(t *testing.T) {
	response, err := Handler(Request{Message: "what"})
	assert.Equal(t, response.Message, "you sent the message what")
	assert.Equal(t, response.Ok, true)
	assert.Equal(t, err, nil)	
}
