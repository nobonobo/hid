# hid

Cross PLatform HID API

- Native(Linux, Windows, MacOS)
- Wasm(js)

## Global Types

- HIDDeviceFilter struct
  - VendorId
  - ProductId

## Global functions

- GetDevices() ([]\*HIDDevice, error)
- RequestDevice(options ...HIDDeviceFilter) ([]\*HIDDevice, error)

## type HIDDevice struct

Getters:

- Opened() bool
- VendorID() uint16
- ProductID() uint16
- Path() string
- ProductName() string

Methods:

- OnInputReport(callback func([]byte)
- Open() error
- Close() error
- SendReport(id byte, data []byte) error
- SendFeatureReport(id byte, data []byte) error
- ReceiveFeatureReport(id byte) ([]byte, error)
