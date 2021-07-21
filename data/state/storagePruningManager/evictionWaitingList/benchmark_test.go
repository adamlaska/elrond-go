package evictionWaitingList

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/data"
	"github.com/ElrondNetwork/elrond-go/data/mock"
	"github.com/ElrondNetwork/elrond-go/hashing/blake2b"
	"github.com/ElrondNetwork/elrond-go/marshal"
	"github.com/stretchr/testify/require"
)

var testHasher = blake2b.NewBlake2b()
var testHashes map[string]data.ModifiedHashes
var roothashes []string
var hashes [][]byte

var ewlEvict *evictionWaitingList
var ewlEvict2 *evictionWaitingListV2

func initTestHashes() {
	if testHashes == nil {
		testHashes, roothashes, hashes = generateTestHashes(10000, 100)
	}
}

func initEvictEWL() {
	if check.IfNil(ewlEvict) {
		initTestHashes()
		ewlEvict, _ = NewEvictionWaitingList(10000, mock.NewMemDbMock(), &marshal.GogoProtoMarshalizer{})

		for _, roothash := range roothashes {
			_ = ewlEvict.Put([]byte(roothash), testHashes[roothash])
		}
	}
}

func initEvictEWL2() {
	if check.IfNil(ewlEvict2) {
		initTestHashes()
		ewlEvict2, _ = NewEvictionWaitingListV2(10000, mock.NewMemDbMock(), &marshal.GogoProtoMarshalizer{})

		for _, roothash := range roothashes {
			_ = ewlEvict2.Put([]byte(roothash), testHashes[roothash])
		}
	}
}

func generateTestHashes(numRoothashes int, numHashesOnRoothash int) (map[string]data.ModifiedHashes, []string, [][]byte) {
	counter := 0
	results := make(map[string]data.ModifiedHashes, numRoothashes)
	resultsRoothashes := make([]string, 0, numRoothashes)
	resultsHashes := make([][]byte, 0, numRoothashes*numHashesOnRoothash)
	for i := 0; i < numRoothashes; i++ {
		rootHash := string(intToHash(counter))
		counter++

		var newHashes data.ModifiedHashes
		newHashes, counter = generateHashes(counter, numHashesOnRoothash)
		for h := range newHashes {
			resultsHashes = append(resultsHashes, []byte(h))
		}

		results[rootHash] = newHashes

		resultsRoothashes = append(resultsRoothashes, rootHash)
	}

	return results, resultsRoothashes, resultsHashes
}

func intToHash(value int) []byte {
	buff := make([]byte, 8)
	binary.BigEndian.PutUint64(buff, uint64(value))

	return testHasher.Compute(string(buff))
}

func generateHashes(counter int, numHashesOnRoothash int) (data.ModifiedHashes, int) {
	result := make(map[string]struct{}, numHashesOnRoothash)

	for i := 0; i < numHashesOnRoothash; i++ {
		hash := intToHash(counter)
		counter++

		result[string(hash)] = struct{}{}
	}

	return result, counter
}

func printMapOfModifiedHashes(hashes map[string]data.ModifiedHashes) {
	for key, val := range hashes {
		fmt.Printf("  %s\n", hex.EncodeToString([]byte(key)))

		for hash := range val {
			fmt.Printf("    %s\n", hex.EncodeToString([]byte(hash)))
		}

		fmt.Println()
	}
}

func BenchmarkEvictionWaitingList_Put(b *testing.B) {
	localEwl, err := NewEvictionWaitingList(10000, mock.NewMemDbMock(), &marshal.GogoProtoMarshalizer{})
	require.Nil(b, err)
	initTestHashes()
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		idx := i % len(roothashes)
		roothash := roothashes[idx]
		modifiedHashes := testHashes[roothash]

		b.StartTimer()
		err = localEwl.Put([]byte(roothash), modifiedHashes)
		b.StopTimer()
		require.Nil(b, err)
	}
}

func BenchmarkEvictionWaitingList_Evict(b *testing.B) {
	initEvictEWL()
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		idx := i % len(roothashes)
		roothash := roothashes[idx]

		b.StartTimer()
		evicted, err := ewlEvict.Evict([]byte(roothash))
		b.StopTimer()
		require.Nil(b, err)
		require.True(b, len(evicted) > 0)

		_ = ewlEvict.Put([]byte(roothash), testHashes[roothash])
	}
}

func BenchmarkEvictionWaitingList_ShouldKeep(b *testing.B) {
	initEvictEWL()
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		idx := i % len(hashes)
		hash := hashes[idx]

		b.StartTimer()
		_, err := ewlEvict.ShouldKeepHash(string(hash), data.TriePruningIdentifier(i%2))
		b.StopTimer()
		require.Nil(b, err)
	}
}

func BenchmarkEvictionWaitingListV2_Put(b *testing.B) {
	ewl, err := NewEvictionWaitingListV2(10000, mock.NewMemDbMock(), &marshal.GogoProtoMarshalizer{})
	require.Nil(b, err)
	initTestHashes()
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		idx := i % len(roothashes)
		roothash := roothashes[idx]
		modifiedHashes := testHashes[roothash]

		b.StartTimer()
		err = ewl.Put([]byte(roothash), modifiedHashes)
		b.StopTimer()
		require.Nil(b, err)
	}
}

func BenchmarkEvictionWaitingListV2_Evict(b *testing.B) {
	initEvictEWL2()
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		idx := i % len(roothashes)
		roothash := roothashes[idx]

		b.StartTimer()
		evicted, err := ewlEvict2.Evict([]byte(roothash))
		b.StopTimer()
		require.Nil(b, err)
		require.True(b, len(evicted) > 0)

		_ = ewlEvict2.Put([]byte(roothash), testHashes[roothash])
	}
}

func BenchmarkEvictionWaitingListV2_ShouldKeep(b *testing.B) {
	initEvictEWL2()
	b.ResetTimer()

	b.StopTimer()
	for i := 0; i < b.N; i++ {
		idx := i % len(hashes)
		hash := hashes[idx]

		b.StartTimer()
		_, err := ewlEvict2.ShouldKeepHash(string(hash), data.TriePruningIdentifier(i%2))
		b.StopTimer()
		require.Nil(b, err)
	}
}
