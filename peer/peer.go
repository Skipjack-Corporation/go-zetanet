package peer

import "errors"

var (
	// ErrEmptyPeerID is an error for empty peer ID.
	ErrEmptyPeerID = errors.New("empty peer ID")
)

// SJID is a zetanet p2p identity.
type SJID string

// Validate checks if ID is empty or not.
func (id SJID) Validate() error {
	if id == SJID("") {
		return ErrEmptyPeerID
	}

	return nil
}

// IDSlice for sorting peers
type IDSlice []SJID

func (es IDSlice) Len() int           { return len(es) }
func (es IDSlice) Swap(i, j int)      { es[i], es[j] = es[j], es[i] }
func (es IDSlice) Less(i, j int) bool { return string(es[i]) < string(es[j]) }
