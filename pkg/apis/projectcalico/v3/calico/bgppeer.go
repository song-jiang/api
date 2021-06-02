// Copyright (c) 2017,2020-2021 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package calico

import (
	k8sv1 "k8s.io/api/core/v1"

	"github.com/projectcalico/api/pkg/lib/numorstring"
)

const (
	KindBGPPeer     = "BGPPeer"
	KindBGPPeerList = "BGPPeerList"
)

// BGPPeerSpec contains the specification for a BGPPeer resource.
type BGPPeerSpec struct {
	// The node name identifying the Calico node instance that is targeted by this peer.
	// If this is not set, and no nodeSelector is specified, then this BGP peer selects all
	// nodes in the cluster.
	// +optional
	Node string `json:"node,omitempty" validate:"omitempty,name"`

	// Selector for the nodes that should have this peering.  When this is set, the Node
	// field must be empty.
	// +optional
	NodeSelector string `json:"nodeSelector,omitempty" validate:"omitempty,selector"`

	// The IP address of the peer followed by an optional port number to peer with.
	// If port number is given, format should be `[<IPv6>]:port` or `<IPv4>:<port>` for IPv4.
	// If optional port number is not set, and this peer IP and ASNumber belongs to a calico/node
	// with ListenPort set in BGPConfiguration, then we use that port to peer.
	// +optional
	PeerIP string `json:"peerIP,omitempty" validate:"omitempty,IP:port"`

	// The AS Number of the peer.
	// +optional
	ASNumber numorstring.ASNumber `json:"asNumber,omitempty"`

	// Selector for the remote nodes to peer with.  When this is set, the PeerIP and
	// ASNumber fields must be empty.  For each peering between the local node and
	// selected remote nodes, we configure an IPv4 peering if both ends have
	// NodeBGPSpec.IPv4Address specified, and an IPv6 peering if both ends have
	// NodeBGPSpec.IPv6Address specified.  The remote AS number comes from the remote
	// node's NodeBGPSpec.ASNumber, or the global default if that is not set.
	// +optional
	PeerSelector string `json:"peerSelector,omitempty" validate:"omitempty,selector"`

	// Option to keep the original nexthop field when routes are sent to a BGP Peer.
	// Setting "true" configures the selected BGP Peers node to use the "next hop keep;"
	// instead of "next hop self;"(default) in the specific branch of the Node on "bird.cfg".
	KeepOriginalNextHop bool `json:"keepOriginalNextHop,omitempty"`

	// Optional BGP password for the peerings generated by this BGPPeer resource.
	Password *BGPPassword `json:"password,omitempty" validate:"omitempty"`
	// Specifies whether and how to configure a source address for the peerings generated by
	// this BGPPeer resource.  Default value "UseNodeIP" means to configure the node IP as the
	// source address.  "None" means not to configure a source address.
	SourceAddress SourceAddress `json:"sourceAddress,omitempty" validate:"omitempty,sourceAddress"`
}

type SourceAddress string

const (
	SourceAddressUseNodeIP SourceAddress = "UseNodeIP"
	SourceAddressNone                    = "None"
)

// BGPPassword contains ways to specify a BGP password.
type BGPPassword struct {
	// Selects a key of a secret in the node pod's namespace.
	SecretKeyRef *k8sv1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}
