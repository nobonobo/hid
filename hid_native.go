//go:build !js
// +build !js

package hid

import (
	"fmt"
	"sync"
	"time"

	lib "github.com/flynn/hid"
)

func GetDevices() ([]*HIDDevice, error) {
	d, err := lib.Devices()
	if err != nil {
		return nil, err
	}
	res := []*HIDDevice{}
	for _, v := range d {
		res = append(res, &HIDDevice{DeviceInfo: v})
	}
	return res, nil
}

func RequestDevice(options ...HIDDeviceFilter) ([]*HIDDevice, error) {
	devices, err := lib.Devices()
	if err != nil {
		return nil, err
	}
	res := []*HIDDevice{}
	for _, d := range devices {
		if len(options) == 0 {
			res = append(res, &HIDDevice{DeviceInfo: d})
			continue
		}
		for _, o := range options {
			if d.VendorID == o.VendorID && d.ProductID == o.ProductID {
				res = append(res, &HIDDevice{DeviceInfo: d})
			}
		}
	}
	return res, nil
}

type request struct {
	id  byte
	res chan []byte
}

type HIDDevice struct {
	*lib.DeviceInfo
	mu   sync.RWMutex
	dev  lib.Device
	req  chan *request
	done chan struct{}
	// attribute EventHandler oninputreport;
	onInputReport func([]byte)
}

func (d *HIDDevice) OnInputReport(callback func([]byte)) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.onInputReport = callback
}

func (d *HIDDevice) Opened() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.dev != nil && d.dev.ReadError() == nil
}

func (d *HIDDevice) VendorID() uint16 {
	return d.DeviceInfo.VendorID
}

func (d *HIDDevice) ProductID() uint16 {
	return d.DeviceInfo.ProductID
}

func (d *HIDDevice) ProductName() string {
	return d.DeviceInfo.Product
}

func (d *HIDDevice) Path() string {
	return d.DeviceInfo.Path
}

// Promise<void> open();
func (d *HIDDevice) Open() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	dev, err := d.DeviceInfo.Open()
	if err != nil {
		return err
	}
	d.dev = dev
	d.req = make(chan *request)
	d.done = make(chan struct{})
	go func() {
		var req *request
		for {
			select {
			case <-d.done:
				return
			case data, ok := <-d.dev.ReadCh():
				if !ok {
					return
				}
				if d.onInputReport != nil {
					d.onInputReport(data)
				}
				if req == nil {
					select {
					case req = <-d.req:
					default:
					}
					if req != nil {
						if req.id == 0 {
							req.res <- data
							req = nil
						} else {
							if req.id == data[0] {
								req.res <- data[1:]
								req = nil
							}
						}
					}
				}
			}
		}
	}()
	return nil
}

// Promise<void> close();
func (d *HIDDevice) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.dev != nil {
		d.dev.Close()
		d.dev = nil
	}
	close(d.req)
	close(d.done)
	return nil
}

// Promise<void> sendReport([EnforceRange] octet reportId, BufferSource data);
func (d *HIDDevice) SendReport(id byte, data []byte) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if d.dev == nil {
		return fmt.Errorf("device not opened")
	}
	return d.dev.Write(append([]byte{id}, data...))
}

// Promise<void> sendFeatureReport([EnforceRange] octet reportId, BufferSource data);
func (d *HIDDevice) SendFeatureReport(id byte, data []byte) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if d.dev == nil {
		return fmt.Errorf("device not opened")
	}
	return d.dev.Write(append([]byte{id}, data...))
}

// Promise<DataView> receiveFeatureReport([EnforceRange] octet reportId);
func (d *HIDDevice) ReceiveFeatureReport(id byte) ([]byte, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if d.dev == nil {
		return nil, fmt.Errorf("device not opened")
	}
	if err := d.dev.ReadError(); err != nil {
		return nil, err
	}
	res := make(chan []byte, 1)
	d.req <- &request{id: id, res: res}
	time.AfterFunc(3+time.Second, func() { close(res) })
	b, ok := <-res
	if !ok {
		return nil, fmt.Errorf("timeout")
	}
	return b, nil
}
