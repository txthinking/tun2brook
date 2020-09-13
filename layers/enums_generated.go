// Copyright 2012 Google, Inc. All rights reserved.

package layers

// Created by gen2.go, don't edit manually
// Generated at 2017-10-23 10:20:24.458771856 -0600 MDT m=+0.001159033

import (
	"fmt"

	"github.com/google/gopacket"
)

func init() {
	initUnknownTypesForIPProtocol()
	initUnknownTypesForProtocolFamily()
	initActualTypeData()
}

// Decoder calls IPProtocolMetadata.DecodeWith's decoder.
func (a IPProtocol) Decode(data []byte, p gopacket.PacketBuilder) error {
	return IPProtocolMetadata[a].DecodeWith.Decode(data, p)
}

// String returns IPProtocolMetadata.Name.
func (a IPProtocol) String() string {
	return IPProtocolMetadata[a].Name
}

// LayerType returns IPProtocolMetadata.LayerType.
func (a IPProtocol) LayerType() gopacket.LayerType {
	return IPProtocolMetadata[a].LayerType
}

type errorDecoderForIPProtocol int

func (a *errorDecoderForIPProtocol) Decode(data []byte, p gopacket.PacketBuilder) error {
	return a
}
func (a *errorDecoderForIPProtocol) Error() string {
	return fmt.Sprintf("Unable to decode IPProtocol %d", int(*a))
}

var errorDecodersForIPProtocol [256]errorDecoderForIPProtocol
var IPProtocolMetadata [256]EnumMetadata

func initUnknownTypesForIPProtocol() {
	for i := 0; i < 256; i++ {
		errorDecodersForIPProtocol[i] = errorDecoderForIPProtocol(i)
		IPProtocolMetadata[i] = EnumMetadata{
			DecodeWith: &errorDecodersForIPProtocol[i],
			Name:       "UnknownIPProtocol",
		}
	}
}

// Decoder calls ProtocolFamilyMetadata.DecodeWith's decoder.
func (a ProtocolFamily) Decode(data []byte, p gopacket.PacketBuilder) error {
	return ProtocolFamilyMetadata[a].DecodeWith.Decode(data, p)
}

// String returns ProtocolFamilyMetadata.Name.
func (a ProtocolFamily) String() string {
	return ProtocolFamilyMetadata[a].Name
}

// LayerType returns ProtocolFamilyMetadata.LayerType.
func (a ProtocolFamily) LayerType() gopacket.LayerType {
	return ProtocolFamilyMetadata[a].LayerType
}

type errorDecoderForProtocolFamily int

func (a *errorDecoderForProtocolFamily) Decode(data []byte, p gopacket.PacketBuilder) error {
	return a
}
func (a *errorDecoderForProtocolFamily) Error() string {
	return fmt.Sprintf("Unable to decode ProtocolFamily %d", int(*a))
}

var errorDecodersForProtocolFamily [256]errorDecoderForProtocolFamily
var ProtocolFamilyMetadata [256]EnumMetadata

func initUnknownTypesForProtocolFamily() {
	for i := 0; i < 256; i++ {
		errorDecodersForProtocolFamily[i] = errorDecoderForProtocolFamily(i)
		ProtocolFamilyMetadata[i] = EnumMetadata{
			DecodeWith: &errorDecodersForProtocolFamily[i],
			Name:       "UnknownProtocolFamily",
		}
	}
}
