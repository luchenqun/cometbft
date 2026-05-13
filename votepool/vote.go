package votepool

import (
	"golang.org/x/crypto/sha3"
)

// DST is the BLS domain separation tag used by moca challenge and staking paths.
// Keep this minimal compatibility surface only; the full historical votepool
// subsystem is intentionally not migrated.
var DST = func() []byte {
	sum := sha3.NewLegacyKeccak256()
	sum.Write([]byte("BLS_SIG_BN254G1_XMD:SHA-256_SVDW_RO_NUL_"))
	return sum.Sum(nil)
}()
