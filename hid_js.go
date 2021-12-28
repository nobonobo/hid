//go:build js
// +build js

package hid

import (
	"syscall/js"

	"github.com/nobonobo/jsutil"
)

var hidObj js.Value

func init() {
	hidObj = js.Global().Get("navigator").Get("hid")
	if hidObj.IsUndefined() {
		panic("WebHID is not supported")
	}
}

func GetDevices() ([]*HIDDevice, error) {
	devices, err := jsutil.Await(hidObj.Call("getDevices"))
	if err != nil {
		return nil, err
	}
	res := []*HIDDevice{}
	for _, v := range devices {
		res = append(res, &HIDDevice{DeviceInfo: v})
	}
	return res, nil
}

func RequestDevice(options ...HIDDeviceFilter) ([]*HIDDevice, error) {
	opts := []js.Value{}
	for _, o := range options {
		m := map[string]interface{}{
			"vendorId":  o.VendorID,
			"productId": o.ProductID,
		}
		opts = append(opts, js.ValueOf(o))
	
	devices, err := jsutil.Await(hidObj.Call("requestDevice", opts))
	if err != nil {
		return nil, err
	}
	res := []*HIDDevice{}
	for _, v := range devices {
		res = append(res, &HIDDevice{DeviceInfo: v})
	}
	return res, nil
}

type HIDDevice struct {
	js.Value
	// attribute EventHandler oninputreport;
	onInputReport js.Func
}

func (d *HIDDevice) OnInputReport(callback func([]byte) {
	if d.onInputReport != nil {
		d.onInputReport.Release()
	}
	d.onInputReport = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		callback(args[0].Bytes())
		return nil
	}
}

func (d *HIDDevice) Opened() bool {
	return byte(d.Value.Get("opened").Bool())
}

func (d *HIDDevice) VendorId() uint16 {
	return uint16(d.Value.Get("vendorId").Int())
}

func (d *HIDDevice) ProductId() uint16 {
	return uint16(d.Value.Get("productId").Int())
}

func (d *HIDDevice) ProductName() string {
	return d.Value.Get("productName").String()
}

func (d *HIDDevice) Path() string {
	return d.Value.Get("productName").String()
}

// Promise<void> open();
func (d *HIDDevice) Open() error {
	_, err := jsutil.Await(d.Value.Call("open"))
	if err != nil {
		return err
	}
	return nil
}

// Promise<void> close();
func (d *HIDDevice) Close() error {
	if d.onInputReport != nil {
		d.onInputReport.Release()
	}
	_, err := jsutil.Await(d.Value.Call("close"))
	if err != nil {
		return err
	}
	return nil
}

// Promise<void> sendReport([EnforceRange] octet reportId, BufferSource data);
func (d *HIDDevice) SendReport(id byte, data []byte) error {
	_, err := jsutil.Await(d.Value.Call("sendReport", id, jsutil.Bytes2JS(data)))
	if err != nil {
		return err
	}
	return nil
}

// Promise<void> sendFeatureReport([EnforceRange] octet reportId, BufferSource data);
func (d *HIDDevice) SendFeatureReport(id byte, data []byte) error {
	_, err := jsutil.Await(d.Value.Call("sendFeatureReport", id, jsutil.Bytes2JS(data)))
	if err != nil {
		return err
	}
	return nil
}

// Promise<DataView> receiveFeatureReport([EnforceRange] octet reportId);
func (d *HIDDevice) ReceiveFeatureReport(id byte) ([]byte, error) {
	data, err := jsutil.Await(d.Value.Call("receiveFeatureReport", id, js.TypedArrayOf(data)))
	if err != nil {
		return nil, err
	}
	return jsutil.JS2Bytes(data), nil
}
