package tcp_util

import (
	"reflect"
	"unsafe"
)

type TcpBuffer struct {
	buff []byte
	Data []byte
}

func NewTcpBuffer(data []byte) *TcpBuffer {
	return &TcpBuffer{
		buff: data,
		Data: data,
	}
}

func (this *TcpBuffer) GetRawBuff() []byte {
	return this.buff
}

func (this *TcpBuffer) WriteUint(pack int, value uint64) *TcpBuffer {
	switch pack {
	case 1:
		this.WriteUint8(uint8(value))
	case 2:
		this.WriteUint16(uint16(value))
	case 4:
		this.WriteUint32(uint32(value))
	case 8:
		this.WriteUint64(uint64(value))
	default:
		panic("pack != 1 || pack != 2 || pack != 4 || pack != 8")
	}

	return this
}

func (this *TcpBuffer) WriteInt8(value int8) *TcpBuffer {
	if len(this.Data) < 1 {
		panic("index out of range")
	}
	this.Data[0] = byte(value)
	this.Data = this.Data[1:]
	return this
}

func (this *TcpBuffer) WriteUint8(value uint8) *TcpBuffer {
	if len(this.Data) < 1 {
		panic("index out of range")
	}
	this.Data[0] = byte(value)
	this.Data = this.Data[1:]
	return this
}

func (this *TcpBuffer) WriteInt16(value int16) *TcpBuffer {
	if len(this.Data) < 2 {
		panic("index out of range")
	}
	*(*int16)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&this.Data)).Data)) = value
	this.Data = this.Data[2:]
	return this
}

func (this *TcpBuffer) WriteUint16(value uint16) *TcpBuffer {
	if len(this.Data) < 2 {
		panic("index out of range")
	}
	*(*uint16)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&this.Data)).Data)) = value
	this.Data = this.Data[2:]
	return this
}

func (this *TcpBuffer) WriteInt32(value int32) *TcpBuffer {
	if len(this.Data) < 4 {
		panic("index out of range")
	}
	*(*int32)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&this.Data)).Data)) = value
	this.Data = this.Data[4:]
	return this
}

func (this *TcpBuffer) WriteUint32(value uint32) *TcpBuffer {
	if len(this.Data) < 4 {
		panic("index out of range")
	}
	*(*uint32)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&this.Data)).Data)) = value
	this.Data = this.Data[4:]
	return this
}

func (this *TcpBuffer) WriteInt64(value int64) *TcpBuffer {
	if len(this.Data) < 8 {
		panic("index out of range")
	}
	*(*int64)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&this.Data)).Data)) = value
	this.Data = this.Data[8:]
	return this
}

func (this *TcpBuffer) WriteUint64(value uint64) *TcpBuffer {
	if len(this.Data) < 8 {
		panic("index out of range")
	}
	*(*uint64)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&this.Data)).Data)) = value
	this.Data = this.Data[8:]
	return this
}

func (this *TcpBuffer) WriteBytes(data []byte) *TcpBuffer {
	if len(this.Data) < len(data) {
		panic("index out of range")
	}
	copy(this.Data, data)
	this.Data = this.Data[len(data):]
	return this
}

func (this *TcpBuffer) WriteBytes8(data []byte) *TcpBuffer {
	this.WriteUint8(uint8(len(data)))
	this.WriteBytes(data)
	return this
}

func (this *TcpBuffer) WriteBytes16(data []byte) *TcpBuffer {
	this.WriteUint16(uint16(len(data)))
	this.WriteBytes(data)
	return this
}

func (this *TcpBuffer) WriteBytes32(data []byte) *TcpBuffer {
	this.WriteUint32(uint32(len(data)))
	this.WriteBytes(data)
	return this
}

func (this *TcpBuffer) RawLen() int {
	return len(this.buff)
}

func (this *TcpBuffer) Len() int {
	return len(this.Data)
}

func (this *TcpBuffer) Seek(n int) *TcpBuffer {
	this.Data = this.Data[n:]
	return this
}

func (this *TcpBuffer) ReadInt8() int8 {
	var result = int8(this.Data[0])
	this.Data = this.Data[1:]
	return result
}

func (this *TcpBuffer) ReadUint8() uint8 {
	var result = uint8(this.Data[0])
	this.Data = this.Data[1:]
	return result
}

func (this *TcpBuffer) ReadInt16() int16 {
	if len(this.Data) < 2 {
		panic("index out of range")
	}
	var result = *(*int16)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&this.Data)).Data))
	this.Data = this.Data[2:]
	return result
}

func (this *TcpBuffer) ReadUint16() uint16 {
	if len(this.Data) < 2 {
		panic("index out of range")
	}
	var result = *(*uint16)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&this.Data)).Data))
	this.Data = this.Data[2:]
	return result
}

func (this *TcpBuffer) ReadInt32() int32 {
	if len(this.Data) < 4 {
		panic("index out of range")
	}
	var result = *(*int32)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&this.Data)).Data))
	this.Data = this.Data[4:]
	return result
}

func (this *TcpBuffer) ReadUint32() uint32 {
	if len(this.Data) < 4 {
		panic("index out of range")
	}
	var result = *(*uint32)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&this.Data)).Data))
	this.Data = this.Data[4:]
	return result
}

func (this *TcpBuffer) ReadInt64() int64 {
	if len(this.Data) < 8 {
		panic("index out of range")
	}
	var result = *(*int64)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&this.Data)).Data))
	this.Data = this.Data[8:]
	return result
}

func (this *TcpBuffer) ReadUint64() uint64 {
	if len(this.Data) < 8 {
		panic("index out of range")
	}
	var result = *(*uint64)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&this.Data)).Data))
	this.Data = this.Data[8:]
	return result
}

func (this *TcpBuffer) ReadBytes(n int) []byte {
	var result = this.Data[:n]
	this.Data = this.Data[n:]
	return result
}

func (this *TcpBuffer) ReadBytes8() []byte {
	var n = this.ReadUint8()
	return this.ReadBytes(int(n))
}

func (this *TcpBuffer) ReadBytes16() []byte {
	var n = this.ReadUint16()
	return this.ReadBytes(int(n))
}

func (this *TcpBuffer) ReadBytes32() []byte {
	var n = this.ReadUint32()
	return this.ReadBytes(int(n))
}

type TcpOutput struct {
	TcpBuffer
	owner *TcpConn
}

func (this *TcpOutput) Send() error {
	return this.owner.sendRaw(this.buff)
}

type TcpInput struct {
	TcpBuffer
}

func NewTcpInput(data []byte) *TcpInput {
	return &TcpInput{
		TcpBuffer: TcpBuffer{
			buff: data,
			Data: data,
		},
	}
}
