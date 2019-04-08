package common

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

/*GlobalPositionIntCov The filtered global position (e.g. fused GPS and accelerometers). The position is in GPS-frame (right-handed, Z-up). It  is designed as scaled integer message since the resolution of float is not sufficient. NOTE: This message is intended for onboard networks / companion computers and higher-bandwidth links and optimized for accuracy and completeness. Please use the GLOBAL_POSITION_INT message for a minimal subset. */
type GlobalPositionIntCov struct {
	/*TimeUsec Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude the number. */
	TimeUsec uint64
	/*Lat Latitude */
	Lat int32
	/*Lon Longitude */
	Lon int32
	/*Alt Altitude in meters above MSL */
	Alt int32
	/*RelativeAlt Altitude above ground */
	RelativeAlt int32
	/*Vx Ground X Speed (Latitude) */
	Vx float32
	/*Vy Ground Y Speed (Longitude) */
	Vy float32
	/*Vz Ground Z Speed (Altitude) */
	Vz float32
	/*Covariance Covariance matrix (first six entries are the first ROW, next six entries are the second row, etc.) */
	Covariance [36]float32
	/*EstimatorType Class id of the estimator this estimate originated from. */
	EstimatorType uint8
	/*HasExtensionFieldValues indicates if this message has any extensions and  */
	HasExtensionFieldValues bool
}

func (m *GlobalPositionIntCov) String() string {
	var builder strings.Builder
	var buffer bytes.Buffer

	writer := tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)

	builder.WriteString("Name:\t%v/%v\n")
	// Output field values based on the decoded message type
	builder.WriteString("TimeUsec:\t%v [us]\n")
	builder.WriteString("Lat:\t%v [degE7]\n")
	builder.WriteString("Lon:\t%v [degE7]\n")
	builder.WriteString("Alt:\t%v [mm]\n")
	builder.WriteString("RelativeAlt:\t%v [mm]\n")
	builder.WriteString("Vx:\t%v [m/s]\n")
	builder.WriteString("Vy:\t%v [m/s]\n")
	builder.WriteString("Vz:\t%v [m/s]\n")
	builder.WriteString("Covariance:\t%v \n")
	builder.WriteString("EstimatorType:\t%v \n")
	format := builder.String()

	fmt.Fprintf(
		writer,
		format,
		m.GetDialect(),
		m.GetMessageName(),
		m.TimeUsec,
		m.Lat,
		m.Lon,
		m.Alt,
		m.RelativeAlt,
		m.Vx,
		m.Vy,
		m.Vz,
		m.Covariance,
		m.EstimatorType,
	)

	writer.Flush()
	return string(buffer.Bytes())
}

// GetVersion gets the MAVLink version of the Message contents
func (m *GlobalPositionIntCov) GetVersion() int {
	if m.HasExtensionFieldValues {
		return 2
	}

	return 1
}

// GetDialect gets the name of the dialect that defines the Message
func (m *GlobalPositionIntCov) GetDialect() string {
	return "common"
}

// GetMessageName gets the name of the Message
func (m *GlobalPositionIntCov) GetMessageName() string {
	return "GlobalPositionIntCov"
}

// GetID gets the ID of the Message
func (m *GlobalPositionIntCov) GetID() uint32 {
	return 63
}

// HasExtensionFields returns true if the message definition contained extensions; false otherwise
func (m *GlobalPositionIntCov) HasExtensionFields() bool {
	return false
}

func (m *GlobalPositionIntCov) getV1Length() int {
	return 181
}

func (m *GlobalPositionIntCov) getIOSlice() []byte {
	return make([]byte, 181+1)
}

// Read sets the field values of the message from the raw message payload
func (m *GlobalPositionIntCov) Read(frame mavlink2.Frame) (err error) {
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

	// Get a slice of bytes long enough for the all the GlobalPositionIntCov fields
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
func (m *GlobalPositionIntCov) Write(version int) (output []byte, err error) {
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
