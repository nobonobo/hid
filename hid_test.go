package hid

import "testing"

func TestBasic(t *testing.T) {
	devices := GetDevices()
	if len(devices) == 0 {
		t.Fatal("no devices found")
	}
	for _, d := range devices {
		if d.VendorID() == 0x0433 && d.ProductID() == 0x5740 {
			t.Logf("found device: %s", d.Path())
		}
	}
}
