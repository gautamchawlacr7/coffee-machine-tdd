package coffee_machine_test

import (
	"testing"

	"github.com/hellogautam/coffee_machine"
)

func TestShouldReturn0GivenEmptyString(t *testing.T) {
	res, _ := coffee_machine.Add("")
	if res != 0 {
		t.Errorf("expected to be 0 got %d", res)
	}
}
