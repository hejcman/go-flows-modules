package biflows

import (
	"github.com/hejcman/go-flows-modules/common"

	"github.com/CN-TU/go-flows/flows"
	"github.com/CN-TU/go-ipfix"
)

/*
╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
│ Init function │
╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯
*/

func init() {

	// ╭╶╶╶╶╶╶╶╶╶┬╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶┬╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ Biflows │ octetDeltaCount │ https://datatracker.ietf.org/doc/html/rfc5102#section-5.10.1 │
	// ╰╴╴╴╴╴╴╴╴╴┴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴┴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewInformationElement(
			"octetDeltaCount",
			common.BiflowsPen,
			1,
			ipfix.Unsigned64Type,
			0),
		"Biflow version",
		flows.FlowFeature,
		func() flows.Feature { return &octetDeltaCount{} },
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶╶┬╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶┬╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ Biflows │ packetDeltaCount │ https://datatracker.ietf.org/doc/html/rfc5102#section-5.10.7 │
	// ╰╴╴╴╴╴╴╴╴╴┴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴┴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewInformationElement(
			"packetDeltaCount",
			common.BiflowsPen,
			2,
			ipfix.Unsigned64Type,
			0),
		"Biflow version",
		flows.FlowFeature,
		func() flows.Feature { return &packetDeltaCount{} },
		flows.RawPacket)
}
