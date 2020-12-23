package main

type Registers struct {
	a byte
	b byte
	c byte
	d byte
	e byte
	f byte
	h byte
	l byte
}

func (reg *Registers) get_af() uint16 {
	return (uint16(reg.a) << 8) | (uint16(reg.f))

}
func (reg *Registers) set_af(val uint16) {
	reg.a = byte((val & 0xFF00) >> 8) 
	reg.f = byte(val & 0xFF)
}


func (reg *Registers) get_bc() uint16 {
	return (uint16(reg.b) << 8) | (uint16(reg.c))

}
func (reg *Registers) set_bc(val uint16) {
	reg.b = byte((val & 0xFF00) >> 8) 
	reg.c = byte(val & 0xFF)
}


func (reg *Registers) get_de() uint16 {
	return (uint16(reg.d) << 8) | (uint16(reg.e))

}
func (reg *Registers) set_de(val uint16) {
	reg.d = byte((val & 0xFF00) >> 8) 
	reg.e = byte(val & 0xFF)
}


func (reg *Registers) get_hl() uint16 {
	return (uint16(reg.h) << 8) | (uint16(reg.l))

}
func (reg *Registers) set_hl(val uint16) {
	reg.h = byte((val & 0xFF00) >> 8) 
	reg.l = byte(val & 0xFF)
}

const (
	zero_offset byte = 7
	subtract_offset byte = 6
	halfCarry_offset byte = 5
	carry_offset byte = 4
)

type FlagsRegister struct {
	zero byte
	subtract byte
	half_carry byte
	carry byte
}

func (flags *FlagsRegister) convert() byte {
	return 7
}

func (flags *FlagsRegister) convert(val byte) {
	
}