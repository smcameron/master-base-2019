package main

import (
	"fmt"

	bw "github.com/HackRVA/master-base-2019/badgewrangler"
	log "github.com/HackRVA/master-base-2019/filelogging"
	irp "github.com/HackRVA/master-base-2019/irpacket"
	"github.com/HackRVA/master-base-2019/serial"
	term "github.com/nsf/termbox-go"
)

var logger = log.Ger

const (
	listening = iota
	ignoring  = iota
)

const (
	erroring = iota
	correct  = iota
)

func reset() {
	term.Sync() // cosmetic purpose?
}

func main() {
	// Set up input and output channels
	packetsIn := make(chan *irp.Packet)
	packetsOut := make(chan *irp.Packet)

	serial.SetConnected(false)
	serial.SetDebug(false)

	serial.OpenPort("/dev/ttyUSB1", 9600)

	go serial.ReadSerial(packetsIn)
	go serial.WriteSerial(packetsOut)
	bw.SetDebug(true)

	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()

	connStatus := ignoring
	errStatus := correct
	nextConnStatusLabel := "Listen"
	nextErrStatusLabel := "IR Error"

	gameData := &bw.GameData{
		BadgeID: uint16(333),
		GameID:  uint16(1234),
		Hits: []*bw.Hit{
			{BadgeID: uint16(101), Timestamp: uint16(33), Team: uint8(2)},
			{BadgeID: uint16(101), Timestamp: uint16(103), Team: uint8(2)},
			{BadgeID: uint16(101), Timestamp: uint16(203), Team: uint8(2)},
			{BadgeID: uint16(101), Timestamp: uint16(303), Team: uint8(2)},
			{BadgeID: uint16(101), Timestamp: uint16(403), Team: uint8(2)},
			{BadgeID: uint16(101), Timestamp: uint16(503), Team: uint8(2)},
			{BadgeID: uint16(101), Timestamp: uint16(603), Team: uint8(2)},
			{BadgeID: uint16(101), Timestamp: uint16(703), Team: uint8(2)}}}

	/*
		gameData := &bw.GameData{
			BadgeID: uint16(333),
			GameID:  uint16(1234),
			Hits:    []*bw.Hit{}}
	*/

	go bw.BadgeHandlePackets(packetsIn, packetsOut, gameData)

	reset()
keyPressListenerLoop:
	for {
		switch connStatus {
		case listening:
			fmt.Println("Listening to base station")
			serial.SetConnected(true)
			nextConnStatusLabel = "Ignore"
		case ignoring:
			fmt.Println("Ignoring base station")
			serial.SetConnected(false)
			nextConnStatusLabel = "Listen"
		}

		switch errStatus {
		case correct:
			fmt.Println("Sending Correct Data")
			bw.SetIrErr(false)
			nextErrStatusLabel = "IR Error"
		case erroring:
			fmt.Println("Simulating IR Error in Data")
			bw.SetIrErr(true)
			nextErrStatusLabel = "No Error"
		}

		fmt.Println("F5:", nextConnStatusLabel, " F6:", nextErrStatusLabel, " Esc: Quit")

		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break keyPressListenerLoop
			case term.KeyF5:
				if connStatus == listening {
					connStatus = ignoring
				} else {
					connStatus = listening
				}
			case term.KeyF6:
				if errStatus == erroring {
					errStatus = correct
				} else {
					errStatus = erroring
				}
			}

		case term.EventError:
			panic(ev.Err)
		}
		reset()
	}
}
