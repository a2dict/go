package idgen

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
	"os"
	"sync/atomic"
	"time"

	"github.com/a2dict/go/base36"
)

// Generator ...
type Generator func() string

// Next ...
func (g Generator) Next() string {
	return g()
}

var (
	hostHash []byte //2-bytes
)

// New return a id generator
// pre should be 2-bytes length
// generation rule: base36(pre[2bytes] + timestamp[4bytes] + hostHash[2bytes] + index[2bytes])
func New(pre []byte) Generator {
	if len(pre) != 2 {
		panic("pre should be 2-bytes length")
	}
	var idx uint32
	return func() string {
		b := bytes.Buffer{}
		b.Write(pre)
		ts := time.Now().Unix()
		b.Write(uint32ToBytes(uint32(ts)))
		b.Write(hostHash)
		i := atomic.AddUint32(&idx, 1)
		b.Write(uint16ToBytes(uint16(i)))

		bs := b.Bytes()
		id := base36.EncodeBytes(bs)
		return id
	}
}

func uint32ToBytes(i32 uint32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, i32)
	return buf
}

func uint16ToBytes(i16 uint16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, i16)
	return buf
}

func init() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	h := fnv.New32()
	h.Write([]byte(hostname))
	hostHash = h.Sum(nil)
}
