package hellonow

import "testing"

func TestNow(t *testing.T) {
	if Now() == false {
		t.Error("Execution of Now has failed")
	}
}
