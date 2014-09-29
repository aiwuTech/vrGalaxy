package packet

import (
	"errors"
	"math"
)

const (
	PACKET_LIMIT = 65533 // 2^16 - 1 - 2 数据包的前两个字节标识数据包的长度
)

type Packet struct {
	pos  uint
	data []byte
}

// 数据包内容
func (p *Packet) Data() []byte {
	return p.data
}

// 数据包长度
func (p *Packet) Length() int {
	return len(p.data)
}

// 数据当前读取位置
func (p *Packet) Pos() uint {
	return p.pos
}

// 略过n个字节的数据
func (p *Packet) Seek(n uint) {
	p.pos += n
}

// 1为true，其余为false
func (p *Packet) ReadBool() (ret bool, err error) {
	b, _err := p.ReadByte()

	if b != byte(1) {
		return false, _err
	}

	return true, _err
}

// 读取1字节的数据
func (p *Packet) ReadByte() (ret byte, err error) {
	if p.pos >= uint(len(p.data)) {
		err = errors.New("read byte failed")
		return
	}

	ret = p.data[p.pos]
	p.pos++
	return
}

// 读取2字节数据，大端显示uint16
func (p *Packet) ReadU16() (ret uint16, err error) {
	if p.pos+2 > uint(len(p.data)) {
		err = errors.New("read uint16 failed")
		return
	}

	buf := p.data[p.pos : p.pos+2]
	ret = uint16(buf[0])<<8 | uint16(buf[1])
	p.pos += 2
	return
}

// 读取数据到byte数组，要读取的数据的长度为当前数据的前2个字节
func (p *Packet) ReadBytes() (ret []byte, err error) {
	if p.pos+2 > uint(len(p.data)) {
		err = errors.New("read bytes header failed")
		return
	}
	size, _ := p.ReadU16()
	if p.pos+uint(size) > uint(len(p.data)) {
		err = errors.New("read bytes data failed")
		return
	}

	ret = p.data[p.pos : p.pos+uint(size)]
	p.pos += uint(size)
	return
}

// 读取数据到string，要读取的数据的长度为当前数据的前2个字节
func (p *Packet) ReadString() (ret string, err error) {
	if p.pos+2 > uint(len(p.data)) {
		err = errors.New("read string header failed")
		return
	}

	size, _ := p.ReadU16()
	if p.pos+uint(size) > uint(len(p.data)) {
		err = errors.New("read string data failed")
		return
	}

	bytes := p.data[p.pos : p.pos+uint(size)]
	p.pos += uint(size)
	ret = string(bytes)
	return
}

// 读取2字节数据，转换成int16
func (p *Packet) ReadS16() (ret int16, err error) {
	_ret, _err := p.ReadU16()
	ret = int16(_ret)
	err = _err
	return
}

// 读取3字节数据，转换成uint32
func (p *Packet) ReadU24() (ret uint32, err error) {
	if p.pos+3 > uint(len(p.data)) {
		err = errors.New("read uint24 failed")
		return
	}

	buf := p.data[p.pos : p.pos+3]
	ret = uint32(buf[0])<<16 | uint32(buf[1])<<8 | uint32(buf[2])
	p.pos += 3
	return
}

// 读取3字节数据，转换成int32
func (p *Packet) ReadS24() (ret int32, err error) {
	_ret, _err := p.ReadU24()
	ret = int32(_ret)
	err = _err
	return
}

// 读取4字节数据，转换成uint32
func (p *Packet) ReadU32() (ret uint32, err error) {
	if p.pos+4 > uint(len(p.data)) {
		err = errors.New("read uint32 failed")
		return
	}

	buf := p.data[p.pos : p.pos+4]
	ret = uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3])
	p.pos += 4
	return
}

// 读取4字节数据，转换成int32
func (p *Packet) ReadS32() (ret int32, err error) {
	_ret, _err := p.ReadU32()
	ret = int32(_ret)
	err = _err
	return
}

// 读取8字节数据，转换成uint64
func (p *Packet) ReadU64() (ret uint64, err error) {
	if p.pos+8 > uint(len(p.data)) {
		err = errors.New("read uint64 failed")
		return
	}

	ret = 0
	buf := p.data[p.pos : p.pos+8]
	for i, v := range buf {
		ret |= uint64(v) << uint((7-i)*8)
	}
	p.pos += 8
	return
}

// 读取8字节数据，转换成int64
func (p *Packet) ReadS64() (ret int64, err error) {
	_ret, _err := p.ReadU64()
	ret = int64(_ret)
	err = _err
	return
}

// 读取4字节数据，转换成float32
func (p *Packet) ReadFloat32() (ret float32, err error) {
	bits, _err := p.ReadU32()
	if _err != nil {
		return float32(0), _err
	}

	ret = math.Float32frombits(bits)
	if math.IsNaN(float64(ret)) || math.IsInf(float64(ret), 0) {
		return 0, nil
	}

	return ret, nil
}

// 读取8字节数据，转换成float64
func (p *Packet) ReadFloat64() (ret float64, err error) {
	bits, _err := p.ReadU64()
	if _err != nil {
		return float64(0), _err
	}

	ret = math.Float64frombits(bits)
	if math.IsNaN(ret) || math.IsInf(ret, 0) {
		return 0, nil
	}

	return ret, nil
}

// 写入n字节数据，初始化为0
func (p *Packet) WriteZeros(n int) {
	zeros := make([]byte, n)
	p.data = append(p.data, zeros...)
}

// 写入布尔值（1字节）
func (p *Packet) WriteBool(v bool) {
	if v {
		p.data = append(p.data, byte(1))
	} else {
		p.data = append(p.data, byte(0))
	}
}

// 写入byte，占位1字节
func (p *Packet) WriteByte(v byte) {
	p.data = append(p.data, v)
}

// 写入uint16数据，占位2字节
func (p *Packet) WriteU16(v uint16) {
	buf := make([]byte, 2)
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
	p.data = append(p.data, buf...)
}

// 写入[]byte，并写入2字节的数据长度
func (p *Packet) WriteBytes(v []byte) {
	p.WriteU16(uint16(len(v)))
	p.data = append(p.data, v...)
}

// 写入[]byte，不写入长度
func (p *Packet) WriteRawBytes(v []byte) {
	p.data = append(p.data, v...)
}

// 写入string，并写入2字节的数据长度
func (p *Packet) WriteString(v string) {
	bytes := []byte(v)
	p.WriteU16(uint16(len(bytes)))
	p.data = append(p.data, bytes...)
}

// 写入int16数据，占位2字节
func (p *Packet) WriteS16(v int16) {
	p.WriteU16(uint16(v))
}

// 写入uint32数据，占位3字节
func (p *Packet) WriteU24(v uint32) {
	buf := make([]byte, 3)
	buf[0] = byte(v >> 16)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v)
	p.data = append(p.data, buf...)
}

// 写入uint32数据，占位4字节
func (p *Packet) WriteU32(v uint32) {
	buf := make([]byte, 4)
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
	p.data = append(p.data, buf...)
}

// 写入int32数据，占位4字节
func (p *Packet) WriteS32(v int32) {
	p.WriteU32(uint32(v))
}

// 写入uint64数据，占位8字节
func (p *Packet) WriteU64(v uint64) {
	buf := make([]byte, 8)
	for i := range buf {
		buf[i] = byte(v >> uint((7-i)*8))
	}

	p.data = append(p.data, buf...)
}

// 写入int64数据，占位8字节
func (p *Packet) WriteS64(v int64) {
	p.WriteU64(uint64(v))
}

// 写入float32数据，占位4字节
func (p *Packet) WriteFloat32(f float32) {
	v := math.Float32bits(f)
	p.WriteU32(v)
}

// 写入float64数据，占位8字节
func (p *Packet) WriteFloat64(f float64) {
	v := math.Float64bits(f)
	p.WriteU64(v)
}

// 数据包读取器
func Reader(data []byte) *Packet {
	return &Packet{pos: 0, data: data}
}

// 数据包写入器
func Writer() *Packet {
	pkt := &Packet{pos: 0}
	pkt.data = make([]byte, 0, 128)
	return pkt
}
