package mp4

import (
	"encoding/binary"
)

// BitRateBox represents the `btrt` box structure
type BitRateBox struct {
	Box          *BasicBox
	BufferSizeDB uint32
	MaxBitrate   uint32
	AvgBitrate   uint32
}

// NewBitRateBox creates a new BitRateBox
func NewBitRateBox(bufferSizeDB, maxBitrate, avgBitrate uint32) *BitRateBox {
	return &BitRateBox{
		Box:          NewBasicBox([4]byte{'b', 't', 'r', 't'}),
		BufferSizeDB: bufferSizeDB,
		MaxBitrate:   maxBitrate,
		AvgBitrate:   avgBitrate,
	}
}

// Size returns the size of the BitRateBox
func (box *BitRateBox) Size() uint64 {
	return 8 + 12 // Box header + BufferSizeDB + MaxBitrate + AvgBitrate
}

// Encode encodes the BitRateBox into a byte slice
func (box *BitRateBox) Encode() (int, []byte) {
	box.Box.Size = box.Size()
	offset, buf := box.Box.Encode()
	binary.BigEndian.PutUint32(buf[offset:], box.BufferSizeDB)
	offset += 4
	binary.BigEndian.PutUint32(buf[offset:], box.MaxBitrate)
	offset += 4
	binary.BigEndian.PutUint32(buf[offset:], box.AvgBitrate)
	offset += 4
	return offset, buf
}
