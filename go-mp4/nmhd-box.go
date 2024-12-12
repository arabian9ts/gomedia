package mp4

import (
	"io"
)

type NullMediaHeaderBox struct {
	Box *FullBox
}

func NewNullMediaHeaderBox() *NullMediaHeaderBox {
	return &NullMediaHeaderBox{
		Box: NewFullBox([4]byte{'n', 'm', 'h', 'd'}, 0),
	}
}

func (nmhd *NullMediaHeaderBox) Size() uint64 {
	return nmhd.Box.Size()
}

func (nmhd *NullMediaHeaderBox) Decode(r io.Reader) (offset int, err error) {
	if offset, err = nmhd.Box.Decode(r); err != nil {
		return 0, err
	}
	buf := make([]byte, 4)
	if _, err = io.ReadFull(r, buf); err != nil {
		return
	}
	return 4, nil
}

func (nmhd *NullMediaHeaderBox) Encode() (int, []byte) {
	nmhd.Box.Box.Size = nmhd.Size()
	nmhd.Box.Flags[2] = 0
	offset, buf := nmhd.Box.Encode()
	return offset, buf
}

func makeNmhdBox() []byte {
	smhd := NewNullMediaHeaderBox()
	_, smhdbox := smhd.Encode()
	return smhdbox
}

func decodeNmhdBox(demuxer *MovDemuxer) (err error) {
	smhd := NullMediaHeaderBox{Box: new(FullBox)}
	_, err = smhd.Decode(demuxer.reader)
	return
}
