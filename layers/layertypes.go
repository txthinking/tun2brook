// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"github.com/google/gopacket"
)

var (
	LayerTypeIPv4            = gopacket.RegisterLayerType(20, gopacket.LayerTypeMetadata{Name: "IPv4", Decoder: gopacket.DecodeFunc(decodeIPv4)})
	LayerTypeIPv6            = gopacket.RegisterLayerType(21, gopacket.LayerTypeMetadata{Name: "IPv6", Decoder: gopacket.DecodeFunc(decodeIPv6)})
	LayerTypeTCP             = gopacket.RegisterLayerType(44, gopacket.LayerTypeMetadata{Name: "TCP", Decoder: gopacket.DecodeFunc(decodeTCP)})
	LayerTypeUDP             = gopacket.RegisterLayerType(45, gopacket.LayerTypeMetadata{Name: "UDP", Decoder: gopacket.DecodeFunc(decodeUDP)})
	LayerTypeIPv6HopByHop    = gopacket.RegisterLayerType(46, gopacket.LayerTypeMetadata{Name: "IPv6HopByHop", Decoder: gopacket.DecodeFunc(decodeIPv6HopByHop)})
	LayerTypeIPv6Routing     = gopacket.RegisterLayerType(47, gopacket.LayerTypeMetadata{Name: "IPv6Routing", Decoder: gopacket.DecodeFunc(decodeIPv6Routing)})
	LayerTypeIPv6Fragment    = gopacket.RegisterLayerType(48, gopacket.LayerTypeMetadata{Name: "IPv6Fragment", Decoder: gopacket.DecodeFunc(decodeIPv6Fragment)})
	LayerTypeIPv6Destination = gopacket.RegisterLayerType(49, gopacket.LayerTypeMetadata{Name: "IPv6Destination", Decoder: gopacket.DecodeFunc(decodeIPv6Destination)})
)

var (
	// LayerClassIPNetwork contains TCP/IP network layer types.
	LayerClassIPNetwork = gopacket.NewLayerClass([]gopacket.LayerType{
		LayerTypeIPv4,
		LayerTypeIPv6,
	})
	// LayerClassIPTransport contains TCP/IP transport layer types.
	LayerClassIPTransport = gopacket.NewLayerClass([]gopacket.LayerType{
		LayerTypeTCP,
		LayerTypeUDP,
	})
	// LayerClassIPv6Extension contains IPv6 extension headers.
	LayerClassIPv6Extension = gopacket.NewLayerClass([]gopacket.LayerType{
		LayerTypeIPv6HopByHop,
		LayerTypeIPv6Routing,
		LayerTypeIPv6Fragment,
		LayerTypeIPv6Destination,
	})
)
