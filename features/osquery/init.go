package osquery

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

	// ╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ Process Name │
	// ╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewInformationElement(
			"OSQueryProgramName",
			common.CesnetPen,
			852,
			ipfix.StringType,
			ipfix.VariableLength),
		"the process which created the flow",
		flows.FlowFeature,
		func() flows.Feature { return &processFeature{} },
		flows.RawPacket,
	)

	// ╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ Kernel version │
	// ╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewInformationElement(
			"OSQueryKernelVersion",
			common.CesnetPen,
			861,
			ipfix.StringType,
			ipfix.VariableLength),
		"kernel version",
		flows.FlowFeature,
		func() flows.Feature { return prepareKernelFeature("version") },
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶╶╮
	// │ OS Name │
	// ╰╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewInformationElement(
			"OSQueryOSName",
			common.CesnetPen,
			854,
			ipfix.StringType,
			ipfix.VariableLength),
		"distribution or product name",
		flows.FlowFeature,
		func() flows.Feature { return prepareOsFeature("major", ipfix.StringType) },
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ OS Major Version │
	// ╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewInformationElement(
			"OSQueryOSMajor",
			common.CesnetPen,
			855,
			ipfix.Unsigned16Type,
			0),
		"major release version",
		flows.FlowFeature,
		func() flows.Feature { return prepareOsFeature("major", ipfix.Unsigned16Type) },
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ OS Minor Version │
	// ╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewInformationElement(
			"OSQueryOSMinor",
			common.CesnetPen,
			856,
			ipfix.Unsigned16Type,
			0),
		"minor release version",
		flows.FlowFeature,
		func() flows.Feature { return prepareOsFeature("minor", ipfix.Unsigned16Type) },
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ OS Patch Version │
	// ╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewInformationElement(
			"OSQueryOSBuild",
			common.CesnetPen,
			857,
			ipfix.StringType,
			ipfix.VariableLength),
		"optional build-specific or variant string",
		flows.FlowFeature,
		func() flows.Feature { return prepareOsFeature("build", ipfix.StringType) },
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ OS Platform │
	// ╰╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewInformationElement(
			"OSQueryOSPlatform",
			common.CesnetPen,
			858,
			ipfix.StringType,
			ipfix.VariableLength),
		"os platform or id",
		flows.FlowFeature,
		func() flows.Feature { return prepareOsFeature("platform", ipfix.StringType) },
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ OS Platform like │
	// ╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewInformationElement(
			"OSQueryOSPlatformLike",
			common.CesnetPen,
			859,
			ipfix.StringType,
			ipfix.VariableLength),
		"closely related platforms",
		flows.FlowFeature,
		func() flows.Feature { return prepareOsFeature("platform_like", ipfix.StringType) },
		flows.RawPacket)

	// ╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
	// │ OS Architecture │
	// ╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯

	flows.RegisterFeature(
		ipfix.NewInformationElement(
			"OSQueryOSArch",
			common.CesnetPen,
			860,
			ipfix.StringType,
			ipfix.VariableLength),
		"os architecture",
		flows.FlowFeature,
		func() flows.Feature { return prepareOsFeature("arch", ipfix.StringType) },
		flows.RawPacket)

}
