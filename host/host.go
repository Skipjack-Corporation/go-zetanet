package host

import (
	peer "github.com/go-zetanet-p2p/peer"
)

// Host is an object participating in a zetanet p2p network
type Host interface {
	ID() peer.SJID
}
