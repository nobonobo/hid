//go:build ignore
// +build ignore

package hid

func (d *HIDDevice) Collections() []*HIDCollectionInfo {
	res := []*HIDCollectionInfo{}
	return res
}

type HIDCollectionInfo struct {
	usagePage      uint16
	usage          uint16
	type_          byte
	children       []*HIDCollectionInfo
	inputReports   []*HIDReportInfo
	outputReports  []*HIDReportInfo
	featureReports []*HIDReportInfo
}

func (c *HIDCollectionInfo) UsagePage() uint16 {
	return uint16(c.Get("usagePage").Int())
}

func (c *HIDCollectionInfo) Usage() uint16 {
	return uint16(c.Get("usage").Int())
}

func (c *HIDCollectionInfo) Type() byte {
	return byte(c.Get("type").Int())
}

func (c *HIDCollectionInfo) Children() []*HIDCollectionInfo {
	res := []*HIDCollectionInfo{}
	v := c.Value.Get("children")
	for i := 0; i < v.Length(); i++ {
		res = append(res, &HIDCollectionInfo{Value: v.Index(i)})
	}
	return res
}

func (c *HIDCollectionInfo) InputReports() []*HIDReportInfo {
	res := []*HIDReportInfo{}
	v := c.Value.Get("inputReports")
	for i := 0; i < v.Length(); i++ {
		res = append(res, &HIDReportInfo{Value: v.Index(i)})
	}
	return res
}

func (c *HIDCollectionInfo) OutputReports() []*HIDReportInfo {
	res := []*HIDReportInfo{}
	v := c.Value.Get("outputReports")
	for i := 0; i < v.Length(); i++ {
		res = append(res, &HIDReportInfo{Value: v.Index(i)})
	}
	return res
}

func (c *HIDCollectionInfo) FeatureReports() []*HIDReportInfo {
	res := []*HIDReportInfo{}
	v := c.Value.Get("featureReports")
	for i := 0; i < v.Length(); i++ {
		res = append(res, &HIDReportInfo{Value: v.Index(i)})
	}
	return res
}

type HIDReportInfo struct {
	reportId byte
	items    []*HIDReportItem
}

func (ri *HIDReportInfo) ReportId() byte {
	return ri.reportId
}

func (ri *HIDReportInfo) Items() []*HIDReportItem {
	return ri.items
}

type HIDReportItem struct {
	isAbsolute                    bool
	isArray                       bool
	isBufferedBytes               bool
	isConstant                    bool
	isLinear                      bool
	isRange                       bool
	isVolatile                    bool
	hasNull                       bool
	hasPreferredState             bool
	wrap                          bool
	usages                        []uint32
	usageMinimum                  uint32
	usageMaximum                  uint32
	reportSize                    uint16
	reportCount                   uint16
	unitExponent                  byte
	unitSystem                    *HIDUnitSystem
	unitFactorLengthExponent      byte
	unitFactorMassExponent        byte
	unitFactorTimeExponent        byte
	unitFactorTemperatureExponent byte
	unitFactorCurrentExponent     byte
	unitFactorLuminousExponent    byte
	logicalMinimum                int32
	logicalMaximum                int32
	physicalMinimum               int32
	physicalMaximum               int32
	strings                       []string
}

func (ri *HIDReportItem) IsAbsolute() bool {
	return ri.isAbsolute
}

func (ri *HIDReportItem) IsArray() bool {
	return ri.isArray
}

func (ri *HIDReportItem) IsBufferedBytes() bool {
	return ri.isBufferedBytes
}

func (ri *HIDReportItem) IsConstant() bool {
	return ri.isConstant
}

func (ri *HIDReportItem) IsLinear() bool {
	return ri.isLinear
}

func (ri *HIDReportItem) IsRange() bool {
	return ri.isRange
}

func (ri *HIDReportItem) IsVolatile() bool {
	return ri.isVolatile
}

func (ri *HIDReportItem) HasNull() bool {
	return ri.hasNull
}

func (ri *HIDReportItem) HasPreferredState() bool {
	return ri.hasPreferredState
}

func (ri *HIDReportItem) Wrap() bool {
	return ri.wrap
}

func (ri *HIDReportItem) Usages() []uint32 {
	return ri.usages
}

func (ri *HIDReportItem) UsageMinimum() uint32 {
	return ri.usageMinimum
}

func (ri *HIDReportItem) UsageMaximum() uint32 {
	return ri.usageMaximum
}

func (ri *HIDReportItem) ReportSize() uint16 {
	return ri.reportSize
}

func (ri *HIDReportItem) ReportCount() uint16 {
	return uri.reportCount
}

func (ri *HIDReportItem) UnitExponent() byte {
	return ri.unitExponent
}

func (ri *HIDReportItem) UnitSystem() *HIDUnitSystem {
	return &HIDUnitSystem{}
}

func (ri *HIDReportItem) UnitFactorLengthExponent() byte {
	return ri.unitFactorLengthExponent
}

func (ri *HIDReportItem) UnitFactorMassExponent() byte {
	return ri.unitFactorMassExponent
}

func (ri *HIDReportItem) UnitFactorTimeExponent() byte {
	return ri.unitFactorTimeExponent
}

func (ri *HIDReportItem) UnitFactorTemperatureExponent() byte {
	return ri.unitFactorTemperatureExponent
}

func (ri *HIDReportItem) UnitFactorCurrentExponent() byte {
	return ri.unitFactorCurrentExponent
}

func (ri *HIDReportItem) UnitFactorLuminousIntensityExponent() byte {
	return ri.unitFactorLuminousExponent
}

func (ri *HIDReportItem) LogicalMinimum() int32 {
	return ri.logicalMinimum
}

func (ri *HIDReportItem) LogicalMaximum() int32 {
	return ri.logicalMaximum
}

func (ri *HIDReportItem) PhysicalMinimum() int32 {
	return ri.physicalMinimum
}

func (ri *HIDReportItem) PhysicalMaximum() int32 {
	return ri.physicalMaximum
}

func (ri *HIDReportItem) Strings() []string {
	return ri.strings
}

type HIDUnitSystem struct{}
