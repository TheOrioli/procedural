package wave

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

// channel is a wrapper for the number of channels in the wave file
type channel uint16

const (
	riffID                    = "RIFF"
	riffType                  = "WAVE"
	fmtID                     = "fmt "
	fmtSize            uint32 = 0x10 // fmt size is constant
	fmtCompressionCode uint16 = 1
	// Mono channel
	Mono channel = 1
	// Stereo channel
	Stereo channel = 2
	dataID         = "data"
)

type chunkHeader struct {
	id   string
	size uint32
}

func (c chunkHeader) write(w io.Writer) {
	binary.Write(w, binary.LittleEndian, []byte(c.id))
	binary.Write(w, binary.LittleEndian, c.size)
}

type fmtChunk struct {
	chunkHeader
	CompressionCode uint16
	Channels        channel
	SampleRate      uint32
	BytesPerSecond  uint32
	BytesPerSample  uint16
	BitsPerSample   uint16
}

func (f fmtChunk) write(w io.Writer) {
	f.chunkHeader.write(w)
	binary.Write(w, binary.LittleEndian, f.CompressionCode)
	binary.Write(w, binary.LittleEndian, f.Channels)
	binary.Write(w, binary.LittleEndian, f.SampleRate)
	binary.Write(w, binary.LittleEndian, f.BytesPerSecond)
	binary.Write(w, binary.LittleEndian, f.BytesPerSample)
	binary.Write(w, binary.LittleEndian, f.BitsPerSample)
}

type dataChunk struct {
	chunkHeader
	Bytes []byte
}

func (d dataChunk) write(w io.Writer) {
	d.chunkHeader.write(w)
	binary.Write(w, binary.LittleEndian, d.Bytes)
}

// Wave struct represents the format of the wave file
// Wave implements io.Writer for writing to the underlying Data array
type Wave struct {
	chunkHeader
	riffType string
	format   fmtChunk
	Data     dataChunk
}

// New creates a new Wave with calculated field values for the format
func New(channels channel, sampleRate uint32) Wave {
	w := Wave{
		chunkHeader: chunkHeader{
			id: riffID,
		},
		riffType: riffType,
		format: fmtChunk{
			chunkHeader: chunkHeader{
				id:   fmtID,
				size: fmtSize,
			},
			CompressionCode: fmtCompressionCode,
			Channels:        channels,
			SampleRate:      sampleRate,
			BytesPerSecond:  sampleRate * uint32(channels),
			BytesPerSample:  uint16(channels),
			BitsPerSample:   uint16(channels) * 8,
		},
		Data: dataChunk{
			chunkHeader: chunkHeader{
				id: dataID,
			},
		},
	}
	return w
}

// Reader return the entire wave struct in a byte buffer
// encoded in little endian
func (w Wave) Reader() io.Reader {
	w.Data.size = uint32(len(w.Data.Bytes))
	w.size = w.Data.size + w.format.size + 4

	var b bytes.Buffer

	w.chunkHeader.write(&b)
	binary.Write(&b, binary.LittleEndian, []byte(w.riffType))
	w.format.write(&b)
	w.Data.write(&b)
	return &b
}

func (w *Wave) Write(p []byte) (int, error) {
	if len(p)%int(w.format.BytesPerSample) != 0 {
		return 0, errors.New("The given array doesn't match the file format")
	}

	w.Data.Bytes = append(w.Data.Bytes, p...)
	return len(p), nil
}
