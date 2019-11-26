package teller

import "testing"

func TestTell(t *testing.T) {
    story := Tell()
    if story == "" {
       t.Errorf("Our teller is sick, come back tomorrow!")
    }
}
