package irpacket

import (
	"fmt"
	"testing"
)

const start uint8 = 1
const command uint8 = 1
const address uint8 = 0x13
const badgeid uint16 = 0x000
const payload uint16 = 0x4000

func testRawPacket() RawPacket {
	return RawPacket(StartBit(start) |
		CommandBit(command) |
		AddressBits(address) |
		BadgeIDBits(badgeid) |
		PayloadBits(payload))
}

func TestBitShifting(t *testing.T) {
	fmt.Println("See how the bits pack")
	fmt.Println()
	fmt.Println("Badge packet is 32-bits:")
	fmt.Println("1 start bit")
	fmt.Println("1 cmd bit")
	fmt.Println("5 address bits (like port number)")
	fmt.Println("9 badge id bits")
	fmt.Println("16 payload bits")
	fmt.Println()
	fmt.Printf("start   - %#6x - %6[1]d - %08[1]b\n", start)
	fmt.Printf("command - %#6x - %6[1]d - %08[1]b\n", command)
	fmt.Printf("address - %#6x - %6[1]d - %08[1]b\n", address)
	fmt.Printf("badgeid - %#6x - %6[1]d - %016[1]b\n", badgeid)
	fmt.Printf("payload - %#6x - %6[1]d - %016[1]b\n", payload)
	fmt.Println()
	fmt.Printf("(start   & 0x01)  << 31   - %032b - %#[1]x\n", StartBit(start))
	fmt.Printf("(command & 0x01)  << 30   - %032b - %#[1]x\n", CommandBit(command))
	fmt.Printf("(address & 0x01f) << 25   - %032b - %#[1]x\n", AddressBits(address))
	fmt.Printf("(badgeid & 0x1ff) << 16   - %032b - %#[1]x\n", BadgeIDBits(badgeid))
	fmt.Printf("(payload & 0x0ffff)       - %032b - %#[1]x\n", PayloadBits(payload))
	fmt.Println()
	fmt.Printf("bits or'd together        - %032b - %#[1]x\n", testRawPacket())

	byte1 := uint8(testRawPacket() & 0x0ff)
	byte2 := uint8((testRawPacket() >> 8) & 0x0ff)
	byte3 := uint8((testRawPacket() >> 16) & 0x0ff)
	byte4 := uint8((testRawPacket() >> 24) & 0x0ff)
	fmt.Printf("Bytes: %#x, %#x, %#x, %#x\n", byte1, byte2, byte3, byte4)
	fmt.Printf("Bytes: %d, %d, %d, %d\n", byte1, byte2, byte3, byte4)

	fmt.Println("Byte Slice:", RawPacketToBytes(testRawPacket()))

	fmt.Printf("12 bit integer mask: %#7x - %016[1]b\n", 0x0800)
}

func TestReadPacket(t *testing.T) {

	testPacket := ReadPacket(testRawPacket())

	fmt.Println()
	PrintPacket(testPacket)
	fmt.Println()
	testPacket.Print()
	fmt.Println()

	if testPacket.Start != start {
		t.Errorf("readPacket(testRawPacket()).Start = start")
	}

	if testPacket.Command != command {
		t.Errorf("readPacket(testRawPacket()).Command = command")
	}

	if testPacket.Address != address {
		t.Errorf("readPacket(testRawPacket()).Address = address")
	}

	if testPacket.BadgeID != badgeid {
		t.Errorf("readPacket(testRawPacket()).BadgeID = badgeid")
	}

	if testPacket.Payload != payload {
		t.Errorf("readPacket(testRawPacket()).Payload = payload")
	}
}

func TestBuildPacket(t *testing.T) {

	testPacket := BuildPacket(badgeid, payload)

	if testPacket.Start != start {
		t.Errorf("testPacket.Start = start")
	}

	if testPacket.Command != command {
		t.Errorf("testPacket.Command = command")
	}

	if testPacket.Address != address {
		t.Errorf("testPacket.Address = address")
	}

	if testPacket.BadgeID != badgeid {
		t.Errorf("testPacket.BadgeID = badgeid")
	}

	if testPacket.Payload != payload {
		t.Errorf("testPacket.Payload = payload")
	}
}

func TestWritePacket(t *testing.T) {

	testPacket := BuildPacket(badgeid, payload)

	if WritePacket(testPacket) != testRawPacket() {
		t.Errorf("writePacket(testPacket()) = testRawPacket")
	}
}
