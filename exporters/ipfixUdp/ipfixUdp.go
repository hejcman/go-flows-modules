package ipfixUdp

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/CN-TU/go-flows/flows"
	"github.com/CN-TU/go-flows/util"
	"github.com/CN-TU/go-ipfix"
	"log"
	"net"
	"os"
	"strconv"
)

type ipfixUdpExporter struct {
	// id is the unique identificator of this exporter.
	id string
	// writer is responsible for writing the incoming flows to the buffer.
	writer *ipfix.MessageStream
	// buffer is a temporary store of the exported IPFix messages before they are sent over UDP.
	buffer *bytes.Buffer
	// dstAddress is the destination IP (v4 or v6) of the IPFix collector.
	dstAddress net.IP
	// dstPort is the destination port of the IPFix collector.
	dstPort uint16
	// observationID is the ID of this current exporter that should be included in the exported IPFix flows.
	observationID uint32
	// allocated is used for temporary storing the IPFix elements.
	allocated map[string]ipfix.InformationElement
	// templates stores the IDs of the exported IPFIX templates.
	templates []int
}

/*
╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
│ Public functions │
╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯
*/

// Init is called when initializing the module. It is used to allocate the necessary variables.
func (e *ipfixUdpExporter) Init() {

	var err error

	// Making the UDP connection to the collector
	addr := e.dstAddress.String() + ":" + strconv.Itoa(int(e.dstPort))
	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Fatal("Cannot establish connection with collector: ", err)
	}

	// Creating the message stream to which we will write.
	e.allocated = make(map[string]ipfix.InformationElement)
	e.buffer = new(bytes.Buffer)
	e.writer, err = ipfix.MakeMessageStream(conn, e.dstPort, e.observationID)
	if err != nil {
		log.Fatal("Couldn't create ipfixUdp message stream: ", err)
	}
}

// ID returns the ID of the exporter.
func (e *ipfixUdpExporter) ID() string {
	return e.id
}

// Fields is used once to let the exporter know which fields it will export (can be used for writing a CSV header).
// Whilst empty, this function must be implemented for compatibility with go-flows.
func (e *ipfixUdpExporter) Fields([]string) {}

// Export is called everytime a flow should be exported.
func (e *ipfixUdpExporter) Export(template flows.Template, features []interface{}, _ flows.DateTimeNanoseconds) {
	id := template.ID()
	if id >= len(e.templates) {
		e.templates = append(e.templates, make([]int, id-len(e.templates)+1)...)
	}
	templateID := e.templates[id]
	if templateID == 0 {
		var err error
		templateID, err = e.writer.AddTemplate(0, template.InformationElements()...)
		if err != nil {
			log.Panic(err)
		}
		e.templates[id] = templateID
	}
	err := e.writer.SendData(0, templateID, features...)
	if err != nil {
		log.Fatal("Unable to send ipfixUdp data: ", err)
	}
}

// Finish is called when the exporter can be destroyed. Write remaining data and wait until shutdown.
func (e *ipfixUdpExporter) Finish() {
	err := e.writer.Flush(0)
	if err != nil {
		log.Fatal("Unable to flush ipfixUdp data: ", err)
	}
}

/*
╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
│ Private functions │
╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯
*/

// newIpfixUdpExporter parses the arguments passed by the user and prepares the default values for the exporter module.
func newIpfixUdpExporter(args []string) (arguments []string, ret util.Module, err error) {
	// Parsing the module arguments and extracting the IP and Port of the collector.
	set := flag.NewFlagSet("ipfixUdp", flag.ExitOnError)
	set.Usage = func() { ipfixUdpHelp("ipfixUdp") }
	dstAddrArg := set.String("addr", "", "The destination collector IP address.")
	dstPortArg := set.Uint("port", 0, "The destination collector port.")
	observationIdArg := set.Uint("observationid", 0, "The observation ID of this exporter.")
	err = set.Parse(args)
	if err != nil {
		return nil, nil, err
	}

	// Checking that both the address and port have been set.
	if set.NArg() < 3 {
		return nil, nil, errors.New("ipfixUdp export needs the address and port as arguments")
	}

	// Checking the the IP address is not malformed.
	addr := net.ParseIP(*dstAddrArg)
	if addr == nil {
		return nil, nil, errors.New("unable to parse the ip address, make sure it is in the correct format")
	}

	// Checking the range of the port.
	if *dstPortArg > 65535 {
		return nil, nil, errors.New("unable to parse the port, make sure it is in the range <0,65535>")
	}

	// Passing the remaining arguments to the other modules
	arguments = set.Args() //[2:]

	ipfix.LoadIANASpec()
	ret = &ipfixUdpExporter{
		id:            "ipfixUdp|",
		writer:        nil,
		dstAddress:    addr,
		dstPort:       uint16(*dstPortArg),
		observationID: uint32(*observationIdArg),
	}
	return
}

func ipfixUdpHelp(name string) {
	_, _ = fmt.Fprintf(os.Stderr, `
The %s exporter writes the output to a UDP socket.

As argument, the output port and IP address are needed.

Usage:
  export %s -addr 192.168.1.1 -port 12345 [-observationid 123]

Flags:
  -addr string
        The destination IP addres.
  -port uint
        The destination port in the range <0, 65535>.
  -observationid uint
        The ID of this exporter.
`, name, name)
}

/*
╭╶╶╶╶╶╶╶╶╶╶╶╶╶╶╶╮
│ Init function │
╰╴╴╴╴╴╴╴╴╴╴╴╴╴╴╴╯
*/

func init() {
	flows.RegisterExporter(
		"ipfixUdp",
		"Exports flows to an IPFix collector over UDP.",
		newIpfixUdpExporter,
		ipfixUdpHelp)
}
