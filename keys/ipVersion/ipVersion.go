package packet

func ipVersionKey(packet Buffer, scratch, _ []byte) (int, int) {
	network := packet.NetworkLayer()
	if network == nil {
		return 0, 0
	}

	// Only the first 4 bits are the version. In this case, we want to check whether
	// it is IPv4. Therefore, by doing AND with 0100 0000, we either get 64 if it is IPv4,
	// or 0 otherwise.
	tmp := network.LayerContents()[0] & 64
	if tmp == 64 {
		scratch[0] = 4
	} else {
		scratch[0] = 6
	}
	return 1, 0
}

func init() {
	RegisterStringKey(
		"ipVersion",
		"IP version.",
		KeyTypeSource,
		KeyLayerNetwork,
		func(string) KeyFunc { return ipVersionKey })
}
