package matrixpilot

/*
Generated using mavgen - https://github.com/ArduPilot/pymavlink/

Copyright 2019 queue-b <https://github.com/queue-b>

Permission is hereby granted, free of charge, to any person obtaining a copy
of the generated software (the "Generated Software"), to deal
in the Generated Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Generated Software, and to permit persons to whom the Generated
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Generated Software.

THE GENERATED SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE GENERATED SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE GENERATED SOFTWARE.
*/

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"text/tabwriter"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/util"
)

/*SerialUdbExtraF19 Backwards compatible version of SERIAL_UDB_EXTRA F19 format */
type SerialUdbExtraF19 struct {
	/*SueAileronOutputChannel SUE aileron output channel */
	SueAileronOutputChannel uint8
	/*SueAileronReversed SUE aileron reversed */
	SueAileronReversed uint8
	/*SueElevatorOutputChannel SUE elevator output channel */
	SueElevatorOutputChannel uint8
	/*SueElevatorReversed SUE elevator reversed */
	SueElevatorReversed uint8
	/*SueThrottleOutputChannel SUE throttle output channel */
	SueThrottleOutputChannel uint8
	/*SueThrottleReversed SUE throttle reversed */
	SueThrottleReversed uint8
	/*SueRudderOutputChannel SUE rudder output channel */
	SueRudderOutputChannel uint8
	/*SueRudderReversed SUE rudder reversed */
	SueRudderReversed uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *SerialUdbExtraF19) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("SueAileronOutputChannel:\t%v \n")
	builder.WriteString("SueAileronReversed:\t%v \n")
	builder.WriteString("SueElevatorOutputChannel:\t%v \n")
	builder.WriteString("SueElevatorReversed:\t%v \n")
	builder.WriteString("SueThrottleOutputChannel:\t%v \n")
	builder.WriteString("SueThrottleReversed:\t%v \n")
	builder.WriteString("SueRudderOutputChannel:\t%v \n")
	builder.WriteString("SueRudderReversed:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.SueAileronOutputChannel,
		m.SueAileronReversed,
		m.SueElevatorOutputChannel,
		m.SueElevatorReversed,
		m.SueThrottleOutputChannel,
		m.SueThrottleReversed,
		m.SueRudderOutputChannel,
		m.SueRudderReversed,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *SerialUdbExtraF19) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *SerialUdbExtraF19) GetDialect() string {
	return "matrixpilot"
}

// GetMessageName gets the name of the Message
func (m *SerialUdbExtraF19) GetMessageName() string {
	return "SerialUdbExtraF19"
}

// GetID gets the ID of the Message
func (m *SerialUdbExtraF19) GetID() uint32 {
	return 185
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *SerialUdbExtraF19) HasExtensionFields() bool {
	return false
}

func (m *SerialUdbExtraF19) getV1Length() int {
	return 8
}

func (m *SerialUdbExtraF19) getIOSlice() []byte {
	return make([]byte, 8+1)
}

// Read sets the field values of the message from the raw message payload
func (m *SerialUdbExtraF19) Read(frame mavlink2.Frame) (err error) {
	version := frame.GetVersion()

	// Ensure only Version 1 or Version 2 were specified
	if version != 1 && version != 2 {
		err = mavlink2.ErrUnsupportedVersion
		return
	}

	// Don't attempt to Read V2 messages from V1 frames
	if m.GetID() > 255 && version < 2 {
		err = mavlink2.ErrDecodeV2MessageV1Frame
		return
	}

	// binary.Read can panic; swallow the panic and return a sane error
	defer func() {
		if r := recover(); r != nil {
			err = mavlink2.ErrPrivateField
		}
	}()

	// Get a slice of bytes long enough for the all the SerialUdbExtraF19 fields
	// binary.Read requires enough bytes in the reader to read all fields, even if
	// the fields are just zero values. This also simplifies handling MAVLink2
	// extensions and trailing zero truncation.
	ioSlice := m.getIOSlice()

	copy(ioSlice, frame.GetMessageBytes())

	// Indicate if
	if version == 2 && m.HasExtensionFields() {
		ioSlice[len(ioSlice)-1] = 1
	}

	reader := bytes.NewReader(ioSlice)

	err = binary.Read(reader, binary.LittleEndian, m)

	return
}

// Write encodes the field values of the message to a byte array
func (m *SerialUdbExtraF19) Write(version int) (output []byte, err error) {
	var buffer bytes.Buffer

	// Ensure only Version 1 or Version 2 were specified
	if version != 1 && version != 2 {
		err = mavlink2.ErrUnsupportedVersion
		return
	}

	// Don't attempt to Write V2 messages to V1 bodies
	if m.GetID() > 255 && version < 2 {
		err = mavlink2.ErrEncodeV2MessageV1Frame
		return
	}

	err = binary.Write(&buffer, binary.LittleEndian, *m)
	if err != nil {
		return
	}

	// V1 uses fixed message lengths and does not include any extension fields
	// Truncate the byte slice to the correct length
	if version == 1 {
		output = buffer.Bytes()[:m.getV1Length()]
	}

	// V2 uses variable message lengths and includes extension fields
	// The variable length is caused by truncating any trailing zeroes from
	// the end of the message before it is added to a frame
	if version == 2 {
		output = util.TruncateV2(buffer.Bytes())
	}

	return

}