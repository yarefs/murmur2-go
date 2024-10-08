package murmur

import (
	"github.com/stretchr/testify/assert"
	"hash/fnv"
	"testing"
)

func TestHashReuse(t *testing.T) {
	hash := New32(123)

	for i := 0; i < 10; i += 1 {
		_, err := hash.Write([]byte("zztop"))
		assert.NoError(t, err)
	}

	s := hash.Sum32()
	assert.Equal(t, s, hash.Sum32())

	// once we reset the hash

	hash.Reset()

	// write some stuff
	_, err := hash.Write([]byte("foo"))
	assert.NoError(t, err)

	first := hash.Sum32()

	// reset it again
	hash.Reset()
	_, err = hash.Write([]byte("foo"))
	assert.NoError(t, err)

	second := hash.Sum32()

	// then the same result
	assert.Equal(t, first, second)
}

func TestFixedHashes(t *testing.T) {
	assert.Equal(t, uint32(1412061192), MurmurHash2([]byte("foo"), 123))
	assert.Equal(t, uint32(1878194508), MurmurHash2([]byte("zztop"), 123))
	assert.Equal(t, uint32(1777016281), MurmurHash2([]byte("foobarbaz"), 234))
	assert.Equal(t, uint32(1668928339), MurmurHash2([]byte("blam"), 777))
}

// -----------------------------------------------------------------------------

var sampleBytes = []byte("hardly a good test, but hey.")

func BenchmarkMurmurHash2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash2(sampleBytes, 42)
	}
}

func BenchmarkMurmurHash2A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash2A(sampleBytes, 42)
	}
}

func BenchmarkMurmurHash64A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash64A(sampleBytes, 42)
	}
}

func BenchmarkHash32_Murmur2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := New32(42)
		h.Write(sampleBytes)
		h.Write(sampleBytes)
		h.Write(sampleBytes)
		h.Sum32()
	}
}

// Benchmark "hash/fnv" to get a comparison on speed.

func BenchmarkHash32_FNV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := fnv.New32()
		h.Write(sampleBytes)
		h.Write(sampleBytes)
		h.Write(sampleBytes)
		h.Sum32()
	}
}

func BenchmarkHash32_FNV1a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := fnv.New32a()
		h.Write(sampleBytes)
		h.Write(sampleBytes)
		h.Write(sampleBytes)
		h.Sum32()
	}
}
