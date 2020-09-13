// Copyright 2012 Google, Inc. All rights reserved.
// Copyright 2009-2011 Andreas Krennmair. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"fmt"

	"github.com/google/gopacket"
)

// EnumMetadata keeps track of a set of metadata for each enumeration value
// for protocol enumerations.
type EnumMetadata struct {
	// DecodeWith is the decoder to use to decode this protocol's data.
	DecodeWith gopacket.Decoder
	// Name is the name of the enumeration value.
	Name string
	// LayerType is the layer type implied by the given enum.
	LayerType gopacket.LayerType
}

// IPProtocol is an enumeration of IP protocol values, and acts as a decoder
// for any type it supports.
type IPProtocol uint8

const (
	IPProtocolIPv6HopByHop    IPProtocol = 0
	IPProtocolIPv4            IPProtocol = 4
	IPProtocolTCP             IPProtocol = 6
	IPProtocolUDP             IPProtocol = 17
	IPProtocolIPv6            IPProtocol = 41
	IPProtocolIPv6Routing     IPProtocol = 43
	IPProtocolIPv6Fragment    IPProtocol = 44
	IPProtocolNoNextHeader    IPProtocol = 59
	IPProtocolIPv6Destination IPProtocol = 60
)

// ProtocolFamily is the set of values defined as PF_* in sys/socket.h
type ProtocolFamily uint8

const (
	ProtocolFamilyIPv4 ProtocolFamily = 2
	// BSDs use different values for INET6... glory be.  These values taken from
	// tcpdump 4.3.0.
	ProtocolFamilyIPv6BSD     ProtocolFamily = 24
	ProtocolFamilyIPv6FreeBSD ProtocolFamily = 28
	ProtocolFamilyIPv6Darwin  ProtocolFamily = 30
	ProtocolFamilyIPv6Linux   ProtocolFamily = 10
)

// Decode a raw v4 or v6 IP packet.
func decodeIPv4or6(data []byte, p gopacket.PacketBuilder) error {
	version := data[0] >> 4
	switch version {
	case 4:
		return decodeIPv4(data, p)
	case 6:
		return decodeIPv6(data, p)
	}
	return fmt.Errorf("Invalid IP packet version %v", version)
}

func initActualTypeData() {
	// Each of the XXXTypeMetadata arrays contains mappings of how to handle enum
	// values for various enum types in gopacket/layers.
	// These arrays are actually created by gen2.go and stored in
	// enums_generated.go.
	//
	// So, EthernetTypeMetadata[2] contains information on how to handle EthernetType
	// 2, including which name to give it and which decoder to use to decode
	// packet data of that type.  These arrays are filled by default with all of the
	// protocols gopacket/layers knows how to handle, but users of the library can
	// add new decoders or override existing ones.  For example, if you write a better
	// TCP decoder, you can override IPProtocolMetadata[IPProtocolTCP].DecodeWith
	// with your new decoder, and all gopacket/layers decoding will use your new
	// decoder whenever they encounter that IPProtocol.

	// Here we link up all enumerations with their respective names and decoders.
	IPProtocolMetadata[IPProtocolIPv4] = EnumMetadata{DecodeWith: gopacket.DecodeFunc(decodeIPv4), Name: "IPv4", LayerType: LayerTypeIPv4}
	IPProtocolMetadata[IPProtocolTCP] = EnumMetadata{DecodeWith: gopacket.DecodeFunc(decodeTCP), Name: "TCP", LayerType: LayerTypeTCP}
	IPProtocolMetadata[IPProtocolUDP] = EnumMetadata{DecodeWith: gopacket.DecodeFunc(decodeUDP), Name: "UDP", LayerType: LayerTypeUDP}
	IPProtocolMetadata[IPProtocolIPv6] = EnumMetadata{DecodeWith: gopacket.DecodeFunc(decodeIPv6), Name: "IPv6", LayerType: LayerTypeIPv6}
	IPProtocolMetadata[IPProtocolIPv6HopByHop] = EnumMetadata{DecodeWith: gopacket.DecodeFunc(decodeIPv6HopByHop), Name: "IPv6HopByHop", LayerType: LayerTypeIPv6HopByHop}
	IPProtocolMetadata[IPProtocolIPv6Routing] = EnumMetadata{DecodeWith: gopacket.DecodeFunc(decodeIPv6Routing), Name: "IPv6Routing", LayerType: LayerTypeIPv6Routing}
	IPProtocolMetadata[IPProtocolIPv6Fragment] = EnumMetadata{DecodeWith: gopacket.DecodeFunc(decodeIPv6Fragment), Name: "IPv6Fragment", LayerType: LayerTypeIPv6Fragment}
	IPProtocolMetadata[IPProtocolIPv6Destination] = EnumMetadata{DecodeWith: gopacket.DecodeFunc(decodeIPv6Destination), Name: "IPv6Destination", LayerType: LayerTypeIPv6Destination}
	IPProtocolMetadata[IPProtocolNoNextHeader] = EnumMetadata{DecodeWith: gopacket.DecodePayload, Name: "NoNextHeader", LayerType: gopacket.LayerTypePayload}
}
