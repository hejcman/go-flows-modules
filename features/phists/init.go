package phists

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
)

/*
╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
│ Init function │
╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯
*/

func init() {

	// ╭╶╶╶╶╶╶╶╶┬╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ PHists │ Sizes (SRC -> DST) │
	// ╰╴╴╴╴╴╴╴╴┴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewBasicList(
			"phistSrcSizes",
			ipfix.NewInformationElement(
				"phistSrcSizes",
				common.CesnetPen,
				1060,
				ipfix.Unsigned16Type,
				0),
			8),
		"histogram of interpacket sizes, SRC -> DST",
		flows.FlowFeature,
		func() flows.Feature {
			return &phistsSizes{
				phists: makePhists(true),
				sizes:  []uint16{0, 0, 0, 0, 0, 0, 0, 0}}
		},
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶┬╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ PHists │ Inter Packet Time (SRC -> DST) │
	// ╰╴╴╴╴╴╴╴╴┴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewBasicList(
			"phistSrcInterPacketTime",
			ipfix.NewInformationElement(
				"phistSrcInterPacketTime",
				common.CesnetPen,
				1061,
				ipfix.Unsigned16Type,
				0),
			8),
		"histogram of interpacket times, SRC -> DST",
		flows.FlowFeature,
		func() flows.Feature {
			return &phistsIpt{
				phists: makePhists(true),
				times:  []uint16{0, 0, 0, 0, 0, 0, 0, 0}}
		},
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶┬╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ PHists │ Sizes (DST -> SRC) │
	// ╰╴╴╴╴╴╴╴╴┴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewBasicList(
			"phistDstSizes",
			ipfix.NewInformationElement(
				"phistDstSizes",
				common.CesnetPen,
				1062,
				ipfix.Unsigned16Type,
				0),
			8),
		"histogram of interpacket sizes, DST -> SRC",
		flows.FlowFeature,
		func() flows.Feature {
			return &phistsSizes{
				phists: makePhists(false),
				sizes:  []uint16{0, 0, 0, 0, 0, 0, 0, 0}}
		},
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶┬╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ PHists │ Inter Packet Time (DST -> SRC) │
	// ╰╴╴╴╴╴╴╴╴┴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewBasicList(
			"phistDstInterPacketTime",
			ipfix.NewInformationElement(
				"phistDstInterPacketTime",
				common.CesnetPen,
				1063,
				ipfix.Unsigned32Type,
				0),
			8),
		"histogram of interpacket times, DST -> SRC",
		flows.FlowFeature,
		func() flows.Feature {
			return &phistsIpt{
				phists: makePhists(false),
				times:  []uint16{0, 0, 0, 0, 0, 0, 0, 0}}
		},
		flows.RawPacket)
}
