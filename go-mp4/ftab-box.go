package mp4

import (
	"encoding/binary"
)

// FontTableBox represents the `ftab` box structure
type FontTableBox struct {
	Box      *BasicBox
	FontID   uint16
	FontName string
}

// NewFontTableBox creates a new FontTableBox with the Serif font
func NewFontTableBox() *FontTableBox {
	return &FontTableBox{
		Box:      NewBasicBox([4]byte{'f', 't', 'a', 'b'}),
		FontID:   1,
		FontName: "Arial",
	}
}

// Size returns the size of the FontTableBox
func (box *FontTableBox) Size() uint64 {
	return 8 + 2 + 2 + 1 + uint64(len(box.FontName))
}

// Encode encodes the FontTableBox into a byte slice
func (box *FontTableBox) Encode() (int, []byte) {
	box.Box.Size = box.Size()
	offset, buf := box.Box.Encode()

	// font id
	binary.BigEndian.PutUint16(buf[offset:], 1) // font entry count
	offset += 2
	binary.BigEndian.PutUint16(buf[offset:], box.FontID)
	offset += 2

	// font name length
	buf[offset] = byte(len(box.FontName))
	offset++

	// font name
	copy(buf[offset:], box.FontName)
	offset += len(box.FontName)

	return offset, buf
}
