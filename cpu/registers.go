package cpu

type ByteReg struct {
	val uint8
}

type TwoByteReg struct {
	high ByteReg
	low  ByteReg
}

func (reg *ByteReg) SetVal(val uint8) {
	reg.val = val
}

func (reg *ByteReg) GetVal() uint8 {
	return reg.val
}

func (reg *ByteReg) Increment() {
	reg.val++
}

func (reg *Bytereg) Decrement() {
	reg.val--
}

func (reg *TwoByteReg) SetHigh(value uint8) {
	reg.high.val = value
}

func (reg *TwoByteReg) GetHigh() uint8 {
	return reg.high.val
}

func (reg *TwoByteReg) GetHighReg() *ByteReg {
	return &reg.high
}

func (reg *TwoByteReg) SetLow(value uint8) {
	reg.low.val = value
}

func (reg *TwoByteReg) GetLow() uint8 {
	return reg.low.val
}

func (reg *TwoByteReg) GetLowReg() *ByteReg {
	return &reg.low
}

func (reg *TwoByteReg) SetVal(val uint16) {
	reg.low.SetValue(uint8(value & 0xFF))
	reg.high.SetValue(uint8((value >> 8) & 0xFF))
}

func (reg *TwoByteReg) GetVal() uint16 {
	high := uint16(reg.high.GetVal())
	low := uint16(reg.low.GetVal())
	return (high << 8) | low
}

func (reg *TwoByteReg) Increment() {
	value := reg.GetVal()
	value++
	reg.SetValue(value)
}

func (reg *TwoByteReg) Decrement() {
	value := reg.GetVal()
	value--
	reg.SetValue(value)
}
