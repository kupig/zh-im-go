package connect

import (
	"errors"
	"io"
)

// read/write buffer
type buffer struct {
	buff  []byte
	start int
	end   int
}

func NewBuffer(bytes []byte) buffer {
	return buffer{bytes, 0, 0}
}

func (b *buffer) PushBuffer(c io.Reader) (int, error) {
	n, err := c.Read(b.buff[b.end:])
	if err != nil {
		return n, err
	}
	b.end += n
	return n, nil
}

func (b *buffer) GetBuffer(start, end int) ([]byte, error) {
	if b.end-b.start >= end-start {
		buf := b.buff[b.start+start : b.start+end]
		b.start += end
		return buf, nil
	}

	return nil, errors.New("start or end postion is error")
}

func (b *buffer) PopBuffer(msgTypeLen, msgDatalen int) {
	if b.start != 0 {
		len := msgTypeLen + msgDatalen

		copy(b.buff, b.buff[len:])
		b.end -= len
		b.start = 0
	}
}

func (b *buffer) Write(c []byte, len int) {

}
