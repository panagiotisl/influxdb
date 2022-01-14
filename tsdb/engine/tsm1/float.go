package tsm1

/*
This code is originally from: https://github.com/dgryski/go-tsz and has been modified to remove
the timestamp compression functionality.

It implements the float compression as presented in: http://www.vldb.org/pvldb/vol8/p1816-teller.pdf.
This implementation uses a sentinel value of NaN which means that float64 NaN cannot be stored using
this version.
*/

import (
	"bytes"
	"fmt"
	"github.com/dgryski/go-bitstream"
	"math"
	"math/bits"
)

// Note: an uncompressed format is not yet implemented.
// floatCompressedGorilla is a compressed format using the gorilla paper encoding
const floatCompressedGorilla = 1

const previousValues = 128
var previousValuesLog2 =  int(math.Log2(previousValues))
var threshold = 6 + uint64(previousValuesLog2)

// uvnan is the constant returned from math.NaN().
const uvnan = 0x7FF8000000000001

// FloatEncoder encodes multiple float64s into a byte slice.
type FloatEncoder struct {
	val [previousValues]float64
	i *index
	err error

	leading  uint64
	trailing uint64

	buf bytes.Buffer
	bw  *bitstream.BitWriter

	first    bool
	finished bool
	current  uint64
	index      uint64
	comparisonsCounter  uint64
}

// NewFloatEncoder returns a new FloatEncoder.
func NewFloatEncoder() *FloatEncoder {
	s := FloatEncoder{
		first:   true,
		leading: ^uint64(0),
		i:       createIndex(),

	}

	s.bw = bitstream.NewWriter(&s.buf)
	s.buf.WriteByte(floatCompressedGorilla << 4)

	return &s
}

// Reset sets the encoder back to its initial state.
func (s *FloatEncoder) Reset() {
	for i := 0; i < previousValues; i++ {
		s.val[i] = 0
	}
	s.i = createIndex()
	s.err = nil
	s.leading = ^uint64(0)
	s.trailing = 0
	s.comparisonsCounter = 0
	s.buf.Reset()
	s.buf.WriteByte(floatCompressedGorilla << 4)

	s.bw.Resume(0x0, 8)

	s.finished = false
	s.first = true
	s.current = 0
	s.index = 0
}

// Bytes returns a copy of the underlying byte buffer used in the encoder.
func (s *FloatEncoder) Bytes() ([]byte, error) {
	return s.buf.Bytes(), s.err
}

// Flush indicates there are no more values to encode.
func (s *FloatEncoder) Flush() {
	if !s.finished {
		// write an end-of-stream record
		s.finished = true
		s.Write(math.NaN())
		s.bw.Flush(bitstream.Zero)
	}
}


// Write encodes v to the underlying buffer.
func (s *FloatEncoder) Write(v float64) {
	//key := math.Float64bits(v) & uint64(0x1f)
	//_ = key
	// Only allow NaN as a sentinel value
	if math.IsNaN(v) && !s.finished {
		s.err = fmt.Errorf("unsupported value: NaN")
		return
	}
	if s.first {
		// first point
		s.val[s.current] = v
		s.i.addRecord(v, s.index)
		s.first = false
		//fmt.Printf("Value: %G, writing first as float64\n", s.val)
		s.bw.WriteBits(math.Float64bits(v), 64)
		return
	}
	previousIndex := s.i.getAll(v, s.index, previousValues, threshold)
	if previousIndex == previousValues {
		previousIndex = s.index % previousValues
	}

	vDelta := math.Float64bits(v) ^ math.Float64bits(s.val[previousIndex])


	if vDelta == 0 {
		s.bw.WriteBits(previousIndex, previousValuesLog2 + 2)
		//fmt.Printf("%d SV: 00%b\n", s.leading, previousIndex)
		s.leading = 65
	} else {

		leading := uint64(bits.LeadingZeros64(vDelta))
		trailing := uint64(bits.TrailingZeros64(vDelta))

		// Clamp number of leading zeros to avoid overflow when encoding
		leading &= 0x1F
		leadingRepresentation := uint64(0)
		if leading < 8 {
			leading = 0
		} else if leading < 12 {
			leading = 8
			leadingRepresentation = 1
		}  else if leading < 14 {
			leading = 12
			leadingRepresentation = 2
		}  else if leading < 16 {
			leading = 14
			leadingRepresentation = 3
		}  else if leading < 18 {
			leading = 16
			leadingRepresentation = 4
		}  else if leading < 20 {
			leading = 18
			leadingRepresentation = 5
		}  else if leading < 22 {
			leading = 20
			leadingRepresentation = 6
		} else if leading >= 22 {
			leading = 22
			leadingRepresentation = 7
		}

		if trailing > threshold {
			sigbits := 64 - leading - trailing
			s.bw.WriteBits(previousValues + previousIndex, previousValuesLog2 + 2)
			s.bw.WriteBits(leadingRepresentation, 3)
			s.bw.WriteBits(sigbits, 6)
			s.bw.WriteBits(vDelta>>trailing, int(sigbits))
			s.leading = 65
			//fmt.Printf("%d NV: 0%b%b%b%b\n", leading, previousValues + previousIndex, leadingRepresentation, sigbits, vDelta>>trailing)
		} else if leading == s.leading {
			s.bw.WriteBits(2, 2)
			s.bw.WriteBits(vDelta, int(64 - leading))
			//fmt.Printf("%d SL: %b%b\n", leading, 2, vDelta)
		} else {
			s.leading, s.trailing = leading, trailing
			s.bw.WriteBits(16 + 8 + leadingRepresentation, 5)
			s.bw.WriteBits(vDelta, int(64 - leading))
		}
	}
	s.current = (s.current + 1) % previousValues
	s.val[s.current] = v
	s.index = s.index + 1
	//fmt.Printf("Adding %v with %d\n", v, s.index)
	s.i.addRecord(v, s.index)
	//fmt.Printf("Total comparisons: %d\n", s.comparisonsCounter)
}

// FloatDecoder decodes a byte slice into multiple float64 values.
type FloatDecoder struct {
	val uint64
	values [previousValues]uint64

	leading  uint64
	trailing uint64
	current uint64

	br BitReader
	b  []byte

	first    bool
	finished bool

	err error
}

// SetBytes initializes the decoder with b. Must call before calling Next().
func (it *FloatDecoder) SetBytes(b []byte) error {
	var v uint64
	if len(b) == 0 {
		v = uvnan
	} else {
		// first byte is the compression type.
		// we currently just have gorilla compression.
		it.br.Reset(b[1:])

		var err error
		v, err = it.br.ReadBits(64)
		if err != nil {
			return err
		}
	}

	// Reset all fields.
	for i := 0; i < previousValues; i++ {
		it.values[i] = 0
	}
	it.val = v
	it.current = 0
	it.values[it.current] = v
	it.leading = 0
	it.trailing = 0
	it.b = b
	it.first = true
	it.finished = false
	it.err = nil

	return nil
}

// Next returns true if there are remaining values to read.
func (it *FloatDecoder) Next() bool {
	if it.err != nil || it.finished {
		return false
	}

	if it.first {
		it.first = false

		// mark as finished if there were no values.
		if it.values[it.current] == uvnan { // IsNaN
			it.finished = true
			return false
		}

		return true
	}

	// read compressed value
	var bit bool
	if it.br.CanReadBitFast() {
		bit = it.br.ReadBitFast()
	} else if v, err := it.br.ReadBit(); err != nil {
		it.err = err
		return false
	} else {
		bit = v
	}
	if !bit {
		var bit bool
		if it.br.CanReadBitFast() {
			bit = it.br.ReadBitFast()
		} else if v, err := it.br.ReadBit(); err != nil {
			it.err = err
			return false
		} else {
			bit = v
		}
		// read index of previous value
		index, err := it.br.ReadBits(uint(previousValuesLog2))
		if err != nil {
			it.err = err
			return false
		}
		//fmt.Printf("INDEX: %b\n", index)
		if !bit {
			it.val = it.values[index]
			it.current = (it.current + 1) % previousValues
			it.values[it.current] = it.val
			//fmt.Printf("%d 1SV 00%b\n", it.leading, index)
		} else {
			bits, err := it.br.ReadBits(3)
			if err != nil {
				it.err = err
				return false
			}
			//fmt.Printf("%d 2NV: 01%b%b", it.leading, index, bits)
			it.leading = getLeadingBits(bits)
			bits, err = it.br.ReadBits(6)
			if err != nil {
				it.err = err
				return false
			}
			//fmt.Printf("%b", bits)
			mbits := bits
			// 0 significant bits here means we overflowed and we actually need 64; see comment in encoder
			if mbits == 0 {
				mbits = 64
			}
			it.trailing = 64 - it.leading - mbits

			sigbits, err := it.br.ReadBits(uint(mbits))
			//fmt.Printf("%b\n", sigbits)
			if err != nil {
				it.err = err
				return false
			}
			vbits := it.values[index]
			vbits ^= (sigbits << it.trailing)

			if vbits == uvnan { // IsNaN
				it.finished = true
				return false
			}
			it.val = vbits
			it.current = (it.current + 1) % previousValues
			it.values[it.current] = vbits
		}
	} else {
		var bit bool
		if it.br.CanReadBitFast() {
			bit = it.br.ReadBitFast()
		} else if v, err := it.br.ReadBit(); err != nil {
			it.err = err
			return false
		} else {
			bit = v
		}
		if !bit {
			//fmt.Printf("%d 3SL 10", it.leading)
			mbits := 64 - it.leading
			// 0 significant bits here means we overflowed and we actually need 64; see comment in encoder
			if mbits == 0 {
				mbits = 64
			}
			it.trailing = 0
		} else {
			bits, err := it.br.ReadBits(3)
			if err != nil {
				it.err = err
				return false
			}
			it.leading = getLeadingBits(bits)
			mbits := 64 - it.leading
			//fmt.Printf("%d %d 4DL 11%b ", it.leading, mbits, bits)
			// 0 significant bits here means we overflowed and we actually need 64; see comment in encoder
			if mbits == 0 {
				mbits = 64
			}
			it.trailing = 0
		}

		mbits := uint(64 - it.leading - it.trailing)
		bits, err := it.br.ReadBits(mbits)
		if err != nil {
			it.err = err
			return false
		}
		//fmt.Printf("%b\n", bits)
		//fmt.Printf("%b\n", it.val)
		vbits := it.val
		vbits ^= (bits << it.trailing)
		//fmt.Printf("%b\n\n", vbits)

		if vbits == uvnan { // IsNaN
			it.finished = true
			return false
		}
		it.val = vbits
		it.current = (it.current + 1) % previousValues
		it.values[it.current] = it.val
	}

	return true
}

func getLeadingBits(bits uint64) uint64 {
	switch bits {
	case 0:
		return 0
	case 1:
		return 8
	case 2:
		return 12
	case 3:
		return 14
	case 4:
		return 16
	case 5:
		return 18
	case 6:
		return 20
	case 7:
		return 22
	}
	return 0
}

// Values returns the current float64 value.
func (it *FloatDecoder) Values() float64 {
	return math.Float64frombits(it.val)
}

// Error returns the current decoding error.
func (it *FloatDecoder) Error() error {
	return it.err
}
