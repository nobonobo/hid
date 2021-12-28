//go:build ignore
// +build ignore

package hid

func (d *HIDDevice) Collections() []*HIDCollectionInfo {
	res := []*HIDCollectionInfo{}
	v := d.Value.Get("collections")
	for i := 0; i < v.Length(); i++ {
		res = append(res, &HIDCollectionInfo{v.Index(i)})
	}
	return res
}

type HIDCollectionInfo struct {
	js.Value
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
		res = append(res, &HIDCollectionInfo{Value:v.Index(i)})
	}
	return res
}

func (c *HIDCollectionInfo) InputReports() []*HIDReportInfo {
	res := []*HIDReportInfo{}
	v := c.Value.Get("inputReports")
	for i := 0; i < v.Length(); i++ {
		res = append(res, &HIDReportInfo{Value:v.Index(i)})
	}
	return res
}

func (c *HIDCollectionInfo) OutputReports() []*HIDReportInfo {
	res := []*HIDReportInfo{}
	v := c.Value.Get("outputReports")
	for i := 0; i < v.Length(); i++ {
		res = append(res, &HIDReportInfo{Value:v.Index(i)})
	}
	return res
}

func (c *HIDCollectionInfo) FeatureReports() []*HIDReportInfo {
	res := []*HIDReportInfo{}
	v := c.Value.Get("featureReports")
	for i := 0; i < v.Length(); i++ {
		res = append(res, &HIDReportInfo{Value:v.Index(i)})
	}
	return res
}

type HIDReportInfo struct {
	js.Value
}

func (ri *HIDReportInfo) ReportId() byte {
	return byte(ri.Get("reportId").Int())
}

func (ri *HIDReportInfo) Items() []*HIDReportItem {
	res := []*HIDReportItem{}
	v := ri.Value.Get("items")
	for i := 0; i < v.Length(); i++ {
		res = append(res, &HIDReportItem{v.Index(i)})
	}
	return res
}

type HIDReportItem struct {
	js.Value
}

func (ri *HIDReportItem) IsAbsolute() bool {
	return ri.Get("isAbsolute").Bool()
}

func (ri *HIDReportItem) IsArray() bool {
	return ri.Get("isArray").Bool()
}

func (ri *HIDReportItem) IsBufferedBytes() bool {
	return ri.Get("isBufferedBytes").Bool()
}

func (ri *HIDReportItem) IsConstant() bool {
	return ri.Get("isConstant").Bool()
}

func (ri *HIDReportItem) IsLinear() bool {
	return ri.Get("isLinear").Bool()
}

func (ri *HIDReportItem) IsRange() bool {
	return ri.Get("isRange").Bool()
}

func (ri *HIDReportItem) IsVolatile() bool {
	return ri.Get("isVolatile").Bool()
}

func (ri *HIDReportItem) HasNull() bool {
	return ri.Get("hasNull").Bool()
}

func (ri *HIDReportItem) HasPreferredState() bool {
	return ri.Get("hasPreferredState").Bool()
}

func (ri *HIDReportItem) Wrap() bool {
	return ri.Get("wrap").Bool()
}

func (ri *HIDReportItem) Usages() []uint32 {
	res := []uint32{}
	v := ri.Get("usages")
	for i := 0; i < v.Length(); i++ {
		res = append(res, v.Index(i).Int())
	}
	return res
}

func (ri *HIDReportItem) UsageMinimum() uint32 {
	return uint32(ri.Value.Get("usageMinimum").Int())
}

func (ri *HIDReportItem) UsageMaximum() uint32 {
	return uint32(ri.Value.Get("usageMaximum").Int())
}

func (ri *HIDReportItem) ReportSize() uint16 {
	return uint16(ri.Value.Get("reportSize").Int())
}

func (ri *HIDReportItem) ReportCount() uint16 {
	return uint16(ri.Value.Get("reportCount").Int())
}

func (ri *HIDReportItem) UnitExponent() byte {
	return byte(ri.Value.Get("unitExponent").Int())
}

func (ri *HIDReportItem) UnitSystem() HIDUnitSystem {
	return &HIDUnitSystem{Value: ri.Get("unitSystem")}
}

func (ri *HIDReportItem) UnitFactorLengthExponent() byte {
	return byte(ri.Value.Get("unitFactorLengthExponent").Int())
}

func (ri *HIDReportItem) UnitFactorMassExponent() byte {
	return byte(ri.Value.Get("unitFactorMassExponent").Int())
}

func (ri *HIDReportItem) UnitFactorTimeExponent() byte {
	return byte(ri.Value.Get("unitFactorTimeExponent").Int())
}

func (ri *HIDReportItem) UnitFactorTemperatureExponent() byte {
	return byte(ri.Value.Get("unitFactorTemperatureExponent").Int())
}

func (ri *HIDReportItem) UnitFactorCurrentExponent() byte {
	return byte(ri.Value.Get("unitFactorCurrentExponent").Int())
}

func (ri *HIDReportItem) UnitFactorLuminousIntensityExponent() byte {
	return byte(ri.Value.Get("unitFactorLuminousIntensityExponent").Int())
}

func (ri *HIDReportItem) LogicalMinimum() int32 {
	return int32(ri.Value.Get("logicalMinimum").Int())
}

func (ri *HIDReportItem) LogicalMaximum() int32 {
	return int32(ri.Value.Get("logicalMaximum").Int())
}

func (ri *HIDReportItem) PhysicalMinimum() int32 {
	return int32(ri.Value.Get("physicalMinimum").Int())
}

func (ri *HIDReportItem) PhysicalMaximum() int32 {
	return int32(ri.Value.Get("physicalMaximum").Int())
}

func (ri *HIDReportItem) Strings() []string {
	res := []string{}
	v := ri.Value.Get("strings")
	for i:=0; i<v.Length(); i++ {
		res = append(res, v.Index(i).String())
	return res
}

type HIDUnitSystem struct {
	js.Value
}
