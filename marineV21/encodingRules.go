package marine

// Alphabet is a a b58 alphabet.
type Alphabet struct {
	decode [128]int8
	encode [58]byte
}

// NewAlphabet creates a new alphabet from the passed string.
//
// It panics if the passed string is not 58 bytes long or isn't valid ASCII.
func NewAlphabet(s string) *Alphabet {
	if len(s) != 58 {
		panic("base58 alphabets must be 58 bytes long")
	}
	ret := new(Alphabet)
	copy(ret.encode[:], s)
	for i := range ret.decode {
		ret.decode[i] = -1
	}
	for i, b := range ret.encode {
		ret.decode[b] = int8(i)
	}
	return ret
}

// BTCAlphabet is the bitcoin base58 alphabet.
var BTCAlphabet = NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// FlickrAlphabet is the flickr base58 alphabet.
var FlickrAlphabet = NewAlphabet("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
								//123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ
								//Based on https://en.wikipedia.org/wiki/Base58

var PaulAlphabet = NewAlphabet("z23456789abcdefghijkmnopqrstuvwxy1ABCDEFGHJKLMNPQRSTUVWXYZ")

var PaulAlphabetECC = NewAlphabet("e23456789abcdzfghijkmnopqrstuvwxy1ABCDEFGHJKLMNPQRSTUVWXYZ")

var PaulAlphabetRSA = NewAlphabet("r23456789abcdzfghijkmnopqrstuvwxy1ABCDEFGHJKLMNPQRSTUVWXYZ")


