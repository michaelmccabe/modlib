package modlib

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println("received message", Config())
}
