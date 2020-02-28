package iobuffer

import (
	"errors"
	"io"
)

type Buffer struct {
	buff  []byte
	start int
	end   int
}

func NewBuffer(bytes []byte) Buffer {
	return Buffer{bytes, 0, 0}
}

func (b *Buffer) PushBuffer(c io.Reader) (int, error) {
	n, err := c.Read(b.buff[b.end:])
	if err != nil {
		return n, err
	}
	b.end += n
	return n, nil
}

func (b *Buffer) GetBuffer(start, end int) ([]byte, error) {
	if b.end-b.start >= end-start {
		buf := b.buff[b.start+start : b.start+end]
		b.start += end
		return buf, nil
	}

	return nil, errors.New("start or end postion is error")
}

func (b *Buffer) PopBuffer(msgTypeLen, msgDatalen int) {
	if b.start != 0 {
		len := msgTypeLen + msgDatalen

		copy(b.buff, b.buff[len:])
		b.end -= len
		b.start = 0
	}
}

func (b *Buffer) Write(c []byte, len int) {

}
