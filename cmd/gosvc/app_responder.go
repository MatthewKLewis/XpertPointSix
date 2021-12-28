package main

import "strconv"

func createResponseBytes(s server, packet []byte) []byte {

	switch packet[3] {
	case 0:
		response := make([]byte, 4)
		return response
	case 1:
		response := make([]byte, 4)
		return response
	case 2: //Command 2 is the standard reporting packet from PointSix
		response := make([]byte, 4)
		var firstByte, _ = strconv.Atoi("C3")
		response[0] = byte(firstByte)
		var secondByte, _ = strconv.Atoi("3C")
		response[1] = byte(secondByte)
		response[2] = 0
		response[3] = 6
		return response
	case 3:
		response := make([]byte, 4)
		return response
	case 4:
		response := make([]byte, 4)
		return response
	case 5:
		response := make([]byte, 4)
		return response
	default:
		response := make([]byte, 4)
		return response
	}
}
