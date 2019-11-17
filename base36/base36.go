package base36

// copy form github.com/martinlindhe/base36/
// edit just for lowcase

import (
	"math"
	"math/big"
	"strings"
)

var (
	base36 = []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z'}

	index = map[byte]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'a': 10, 'b': 11, 'c': 12, 'd': 13, 'e': 14,
		'f': 15, 'g': 16, 'h': 17, 'i': 18, 'j': 19,
		'k': 20, 'l': 21, 'm': 22, 'n': 23, 'o': 24,
		'p': 25, 'q': 26, 'r': 27, 's': 28, 't': 29,
		'u': 30, 'v': 31, 'w': 32, 'x': 33, 'y': 34,
		'z': 35,
	}
)

// Encode encodes a number to base36
func Encode(value uint64) string {

	var res [16]byte
	var i int
	for i = len(res) - 1; value != 0; i-- {
		res[i] = base36[value%36]
		value /= 36
	}
	return string(res[i+1:])
}

// Decode decodes a base36-encoded string
func Decode(s string) uint64 {

	res := uint64(0)
	l := len(s) - 1
	for idx := range s {
		c := s[l-idx]
		byteOffset := index[c]
		res += uint64(byteOffset) * uint64(math.Pow(36, float64(idx)))
	}
	return res
}

var bigRadix = big.NewInt(36)
var bigZero = big.NewInt(0)

// EncodeBytesAsBytes encodes a byte slice base36
func EncodeBytesAsBytes(b []byte) []byte {
	x := new(big.Int)
	x.SetBytes(b)

	answer := make([]byte, 0, len(b)*136/100)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, base36[mod.Int64()])
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, base36[0])
	}

	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return answer
}

// EncodeBytes encodes a byte slice to base36 string
func EncodeBytes(b []byte) string {
	return string(EncodeBytesAsBytes(b))
}

// DecodeToBytes decodes a modified base58 string to a byte slice, using alphabet.
func DecodeToBytes(b string) []byte {

	alphabet := string(base36)
	answer := big.NewInt(0)
	j := big.NewInt(1)

	for i := len(b) - 1; i >= 0; i-- {
		tmp := strings.IndexAny(alphabet, string(b[i]))
		if tmp == -1 {
			return []byte("")
		}
		idx := big.NewInt(int64(tmp))
		tmp1 := big.NewInt(0)
		tmp1.Mul(j, idx)

		answer.Add(answer, tmp1)
		j.Mul(j, bigRadix)
	}

	tmpval := answer.Bytes()

	var numZeros int
	for numZeros = 0; numZeros < len(b); numZeros++ {
		if b[numZeros] != alphabet[0] {
			break
		}
	}
	flen := numZeros + len(tmpval)
	val := make([]byte, flen, flen)
	copy(val[numZeros:], tmpval)

	return val
}
