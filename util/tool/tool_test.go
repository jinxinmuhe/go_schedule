package tool

import "testing"

func TestIPAcquire(t *testing.T) {
	ipAcquire()
	t.Log(IP)
}