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

// SJIDSlice for sorting peers
type SJIDSlice []SJID

func (es SJIDSlice) Len() int           { return len(es) }
func (es SJIDSlice) Swap(i, j int)      { es[i], es[j] = es[j], es[i] }
func (es SJIDSlice) Less(i, j int) bool { return string(es[i]) < string(es[j]) }
