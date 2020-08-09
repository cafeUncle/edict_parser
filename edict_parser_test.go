package edict_parser

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestEdictParser(t *testing.T) {
	word, err := ParseEdict("./edict2")

	if err != nil {
		log.Fatal(err.Error())
	}

	assert.NotNil(t, word, "word should not be nil")
}
