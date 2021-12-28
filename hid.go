package hid

type HIDDeviceFilter struct {
	VendorID  uint16
	ProductID uint16
	//UsagePage int16
	//Usage     int16
}

type hidDevice interface {
	Opened() bool
	VendorID() uint16
	ProductID() uint16
	Path() string
	ProductName() string

	Open() error
	Close() error
	SendReport(id byte, data []byte) error
	SendFeatureReport(id byte, data []byte) error
	ReceiveFeatureReport(id byte) ([]byte, error)
}

var _ hidDevice = (*HIDDevice)(nil)
