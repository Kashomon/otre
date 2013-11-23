package core

import testing

func TestOutput(t *testing) {
	output := TestOutput()
	if output != "Foo" {
		t.Errorf("shit")
	}
}
