package main

import (
	"encoding/binary"
	"encoding/json"
	"strconv"
)

//"encoding/binary"

// C3 3C - 2 byte identifier

// Cmd – (1 byte) Command: 2 – UDP Sensor Data; 5 – UDP Simulated Sensor Data (Wifi Sensor Utility).

// // DATA 1:
// PktCnt* – (2 bytes) packet count. The device will increment this count every time it transmits a UDP PassThru packet.
// MAC – (18 bytes – null terminated string) device MAC address. If the MAC address does not apply this field will contain a unique identifier for the device. If not used, this field will be set to all zeros. (ex: “00:23:b4:39:03:47”) (NULL terminated)
// reserved – (8 bytes) set all bytes to 0.
// Locator1 – character that represents where a sensor packet entered the repeater network. (" ", "a"-"z" and "A"- "Z"). Normally set to NULL(0) for Wifi sensors.
// Locator2 - character that represents where a sensor packet entered the repeater network. (" ", "a"-"z" and "A"- "Z"). Will be identical to Locator1. Normally set to NULL(0) for Wifi sensors.
// Sensor Pkt – (29 bytes) sensor packet. (includes the CR terminator) See the document "Point Six Wireless Transmitter Packet-Data Specification " for more information about specific sensors.

// // DATA 2:
// Org – originator type that generated the packet. 0 – Wifi Sensor; 1 – Point Manager; 2 – Ethernet Point Re-peater; 3 - Application
// Transmissions* – (3 bytes) number of transmissions since last battery reset. 0 if no battery support.
// Max Transmissions+ – (3 bytes) maximum number of transmissions for the power source (0 to 16777216 where 0 is unlimited)
// Period* – (2 bytes) transmit interval in seconds.

// 	Alarm – (1 byte) sensor is in alarm state: 0 – no alarm
// 	// 	Bit 0: I/O 1 – low alarm
// 	// 	Bit 1: I/O 1 – high alarm
// 	// 	Bit 2: I/O 2 – low alarm
// 	// 	Bit 3: I/O 2 – high alarm
// 	// 	Bit 4: I/O 1 – low alarm reset: 0 - reset
// 	// 	Bit 5: I/O 1 – high alarm reset: 0 - reset
// 	// 	Bit 6: I/O 2 – low alarm reset: 0 - reset
// 	// 	Bit 7: I/O 2 – high alarm reset: 0 - reset
// 	Reserved – (2 bytes) set all bytes to 0.

// // //  * Most significant byte is first.
// // //  UDP Sensor Packets that include only Data1 are 63 bytes. UDP Sensor Packets that include Data1 and Data2 are 75 bytes. Older sensors contained Data1 but not Data2. Newer sensors include Data1 and Data2.

//commandBuffer[52] = 48;
//commandBuffer[53] = 56;
//commandBuffer[54] = 70;
//commandBuffer[55] = 53;

type PointSixMessage struct {
	CMD           uint16
	MAC           string
	Loc1          uint8
	Loc2          uint8
	SensorPacket  []byte
	Org           byte
	Transmissions []byte
	Period        uint16
	Alarm         byte

	Temperature float32
}

// The wrapper of your app
func parse(packet []byte) []byte {

	var x PointSixMessage
	x.CMD = binary.BigEndian.Uint16(packet[2:4])
	x.MAC = string(packet[6:23])
	x.Loc1 = packet[32]
	x.Loc2 = packet[33]
	x.SensorPacket = packet[34:63]

	var tempHex = string(packet[52:56])
	var tempInt, _ = strconv.ParseInt(tempHex, 16, 32)
	x.Temperature = (float32(tempInt) * 0.0977) - 200

	x.Org = packet[63]
	x.Transmissions = packet[64:67]
	x.Period = binary.BigEndian.Uint16(packet[70:72])
	x.Alarm = packet[72]

	retString, err := json.Marshal(x)
	if err != nil {
		panic("dang")
	}

	return retString
}
