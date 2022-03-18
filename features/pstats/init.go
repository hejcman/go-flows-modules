package pstats

import (
	"github.com/hejcman/go-flows-modules/common"

	"github.com/CN-TU/go-flows/flows"
	"github.com/CN-TU/go-ipfix"
)

/*
╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
│ Common variables and definitions │
╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯
*/

var (
	// includeZeroes specifies whether packets with a payload length of 0 bytes should be included
	// in the packet statistics.
	includeZeroes = false

	// maxElemCount contains the number of packets, for which stats should be gathered.
	// In essence, it is the length of all the IPFIX lists outputed by this feature.
	maxElemCount uint = 30

	// skipDup skips duplicated (retransmitted) TCP packets.
	skipDup = false
)

/*
╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
│ Init function │
╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯
*/

func init() {

	// ╭╶╶╶╶╶╶╶╶┬╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ PStats │ Packet Payload Length │
	// ╰╴╴╴╴╴╴╴╴┴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewBasicList(
			"packetLength",
			ipfix.NewInformationElement(
				"packetLength",
				common.CesnetPen,
				1013,
				ipfix.Signed16Type,
				0),
			0),
		"sizes of the first packets",
		flows.FlowFeature,
		func() flows.Feature { return &pktLengths{pstats: makePstats()} },
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶┬╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ PStats │ Packet Payload Length │
	// ╰╴╴╴╴╴╴╴╴┴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewBasicList(
			"packetTime",
			ipfix.NewInformationElement(
				"packetTime",
				common.CesnetPen,
				1014,
				ipfix.DateTimeMillisecondsType,
				0),
			0),
		"timestamps of the first packets",
		flows.FlowFeature,
		func() flows.Feature { return &pktTimes{pstats: makePstats()} },
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶┬╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ PStats │ TCP Packet Flags │
	// ╰╴╴╴╴╴╴╴╴┴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewBasicList(
			"packetFlag",
			ipfix.NewInformationElement(
				"packetFlag",
				common.CesnetPen,
				1015,
				ipfix.Unsigned8Type,
				0),
			0),
		"TCP flags for each packet",
		flows.FlowFeature,
		func() flows.Feature { return &pktFlags{pstats: makePstats()} },
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶┬╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ PStats │ Packet directions │
	// ╰╴╴╴╴╴╴╴╴┴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewBasicList(
			"packetDirection",
			ipfix.NewInformationElement(
				"packetDirection",
				common.CesnetPen,
				1016,
				ipfix.Signed8Type,
				0),
			0),
		"directions of the first packets",
		flows.FlowFeature,
		func() flows.Feature { return &pktDirections{pstats: makePstats()} },
		flows.RawPacket)
}
