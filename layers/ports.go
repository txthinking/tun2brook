// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"strconv"

	"github.com/google/gopacket"
)

// TCPPort is a port in a TCP layer.
type TCPPort uint16

// UDPPort is a port in a UDP layer.
type UDPPort uint16

// {TCP,UDP,SCTP}PortNames can be found in iana_ports.go

// String returns the port as "number(name)" if there's a well-known port name,
// or just "number" if there isn't.  Well-known names are stored in
// TCPPortNames.
func (a TCPPort) String() string {
	return strconv.Itoa(int(a))
}

// LayerType returns a LayerType that would be able to decode the
// application payload. It uses some well-known ports such as 53 for
// DNS.
//
// Returns gopacket.LayerTypePayload for unknown/unsupported port numbers.
func (a TCPPort) LayerType() gopacket.LayerType {
	return gopacket.LayerTypePayload
}

// String returns the port as "number(name)" if there's a well-known port name,
// or just "number" if there isn't.  Well-known names are stored in
// UDPPortNames.
func (a UDPPort) String() string {
	return strconv.Itoa(int(a))
}

// LayerType returns a LayerType that would be able to decode the
// application payload. It uses some well-known ports such as 53 for
// DNS.
//
// Returns gopacket.LayerTypePayload for unknown/unsupported port numbers.
func (a UDPPort) LayerType() gopacket.LayerType {
	return gopacket.LayerTypePayload
}
