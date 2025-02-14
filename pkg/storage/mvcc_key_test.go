// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package storage

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"testing/quick"

	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/util/hlc"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/cockroach/pkg/util/shuffle"
	"github.com/stretchr/testify/require"
)

// Verify the sort ordering of successive keys with metadata and
// versioned values. In particular, the following sequence of keys /
// versions:
//
// a
// a<t=max>
// a<t=1>
// a<t=0>
// a\x00
// a\x00<t=max>
// a\x00<t=1>
// a\x00<t=0>
func TestMVCCKeys(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	aKey := roachpb.Key("a")
	a0Key := roachpb.Key("a\x00")
	keys := mvccKeys{
		mvccKey(aKey),
		mvccVersionKey(aKey, hlc.Timestamp{WallTime: math.MaxInt64}),
		mvccVersionKey(aKey, hlc.Timestamp{WallTime: 1}),
		mvccVersionKey(aKey, hlc.Timestamp{Logical: 1}),
		mvccKey(a0Key),
		mvccVersionKey(a0Key, hlc.Timestamp{WallTime: math.MaxInt64}),
		mvccVersionKey(a0Key, hlc.Timestamp{WallTime: 1}),
		mvccVersionKey(a0Key, hlc.Timestamp{Logical: 1}),
	}
	sortKeys := make(mvccKeys, len(keys))
	copy(sortKeys, keys)
	shuffle.Shuffle(sortKeys)
	sort.Sort(sortKeys)
	if !reflect.DeepEqual(sortKeys, keys) {
		t.Errorf("expected keys to sort in order %s, but got %s", keys, sortKeys)
	}
}

func TestMVCCKeyCompare(t *testing.T) {
	defer leaktest.AfterTest(t)()

	a1 := MVCCKey{roachpb.Key("a"), hlc.Timestamp{Logical: 1}}
	a2 := MVCCKey{roachpb.Key("a"), hlc.Timestamp{Logical: 2}}
	b0 := MVCCKey{roachpb.Key("b"), hlc.Timestamp{Logical: 0}}
	b1 := MVCCKey{roachpb.Key("b"), hlc.Timestamp{Logical: 1}}
	b2 := MVCCKey{roachpb.Key("b"), hlc.Timestamp{Logical: 2}}
	b2S := MVCCKey{roachpb.Key("b"), hlc.Timestamp{Logical: 2, Synthetic: true}}

	testcases := map[string]struct {
		a      MVCCKey
		b      MVCCKey
		expect int
	}{
		"equal":               {a1, a1, 0},
		"key lt":              {a1, b1, -1},
		"key gt":              {b1, a1, 1},
		"time lt":             {a2, a1, -1}, // MVCC timestamps sort in reverse order
		"time gt":             {a1, a2, 1},  // MVCC timestamps sort in reverse order
		"empty time lt set":   {b0, b1, -1}, // empty MVCC timestamps sort before non-empty
		"set time gt empty":   {b1, b0, 1},  // empty MVCC timestamps sort before non-empty
		"key time precedence": {a1, b2, -1}, // a before b, but 2 before 1; key takes precedence
		"synthetic equal":     {b2, b2S, 0}, // synthetic bit does not affect ordering
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expect, tc.a.Compare(tc.b))
			require.Equal(t, tc.expect == 0, tc.a.Equal(tc.b))
			require.Equal(t, tc.expect < 0, tc.a.Less(tc.b))
			require.Equal(t, tc.expect > 0, tc.b.Less(tc.a))

			// Comparators on encoded keys should be identical.
			aEnc, bEnc := EncodeMVCCKey(tc.a), EncodeMVCCKey(tc.b)
			require.Equal(t, tc.expect, EngineKeyCompare(aEnc, bEnc))
			require.Equal(t, tc.expect == 0, EngineKeyEqual(aEnc, bEnc))
		})
	}
}

func TestMVCCKeyCompareRandom(t *testing.T) {
	defer leaktest.AfterTest(t)()

	f := func(aGen, bGen randMVCCKey) bool {
		a, b := MVCCKey(aGen), MVCCKey(bGen)
		aEnc, bEnc := EncodeMVCCKey(a), EncodeMVCCKey(b)

		cmp := a.Compare(b)
		cmpEnc := EngineKeyCompare(aEnc, bEnc)
		eq := a.Equal(b)
		eqEnc := EngineKeyEqual(aEnc, bEnc)
		lessAB := a.Less(b)
		lessBA := b.Less(a)

		if cmp != cmpEnc {
			t.Logf("cmp (%v) != cmpEnc (%v)", cmp, cmpEnc)
			return false
		}
		if eq != eqEnc {
			t.Logf("eq (%v) != eqEnc (%v)", eq, eqEnc)
			return false
		}
		if (cmp == 0) != eq {
			t.Logf("(cmp == 0) (%v) != eq (%v)", cmp == 0, eq)
			return false
		}
		if (cmp < 0) != lessAB {
			t.Logf("(cmp < 0) (%v) != lessAB (%v)", cmp < 0, lessAB)
			return false
		}
		if (cmp > 0) != lessBA {
			t.Logf("(cmp > 0) (%v) != lessBA (%v)", cmp > 0, lessBA)
			return false
		}
		return true
	}
	require.NoError(t, quick.Check(f, nil))
}

// randMVCCKey is a quick.Generator for MVCCKey.
type randMVCCKey MVCCKey

func (k randMVCCKey) Generate(r *rand.Rand, size int) reflect.Value {
	k.Key = []byte([...]string{"a", "b", "c"}[r.Intn(3)])
	k.Timestamp.WallTime = r.Int63n(5)
	k.Timestamp.Logical = r.Int31n(5)
	if !k.Timestamp.IsEmpty() {
		// NB: the zero timestamp cannot be synthetic.
		k.Timestamp.Synthetic = r.Intn(2) != 0
	}
	return reflect.ValueOf(k)
}

func TestEncodeDecodeMVCCKeyAndTimestampWithLength(t *testing.T) {
	defer leaktest.AfterTest(t)()

	testcases := map[string]struct {
		key     string
		ts      hlc.Timestamp
		encoded string // hexadecimal
	}{
		"empty":                  {"", hlc.Timestamp{}, "00"},
		"only key":               {"foo", hlc.Timestamp{}, "666f6f00"},
		"no key":                 {"", hlc.Timestamp{WallTime: 1643550788737652545}, "0016cf10bc0505574109"},
		"walltime":               {"foo", hlc.Timestamp{WallTime: 1643550788737652545}, "666f6f0016cf10bc0505574109"},
		"logical":                {"foo", hlc.Timestamp{Logical: 65535}, "666f6f0000000000000000000000ffff0d"},
		"synthetic":              {"foo", hlc.Timestamp{Synthetic: true}, "666f6f00000000000000000000000000010e"},
		"walltime and logical":   {"foo", hlc.Timestamp{WallTime: 1643550788737652545, Logical: 65535}, "666f6f0016cf10bc050557410000ffff0d"},
		"walltime and synthetic": {"foo", hlc.Timestamp{WallTime: 1643550788737652545, Synthetic: true}, "666f6f0016cf10bc0505574100000000010e"},
		"logical and synthetic":  {"foo", hlc.Timestamp{Logical: 65535, Synthetic: true}, "666f6f0000000000000000000000ffff010e"},
		"all":                    {"foo", hlc.Timestamp{WallTime: 1643550788737652545, Logical: 65535, Synthetic: true}, "666f6f0016cf10bc050557410000ffff010e"},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {

			// Test Encode/DecodeMVCCKey.
			expect, err := hex.DecodeString(tc.encoded)
			require.NoError(t, err)
			if len(expect) == 0 {
				expect = nil
			}

			mvccKey := MVCCKey{Key: []byte(tc.key), Timestamp: tc.ts}

			encoded := EncodeMVCCKey(mvccKey)
			require.Equal(t, expect, encoded)
			require.Equal(t, len(encoded), encodedMVCCKeyLength(mvccKey))
			require.Equal(t, len(encoded),
				encodedMVCCKeyPrefixLength(mvccKey.Key)+encodedMVCCTimestampSuffixLength(mvccKey.Timestamp))

			decoded, err := DecodeMVCCKey(encoded)
			require.NoError(t, err)
			require.Equal(t, mvccKey, decoded)

			// Test EncodeMVCCKeyPrefix.
			expectPrefix, err := hex.DecodeString(tc.encoded[:2*len(tc.key)+2])
			require.NoError(t, err)
			require.Equal(t, expectPrefix, EncodeMVCCKeyPrefix(roachpb.Key(tc.key)))
			require.Equal(t, len(expectPrefix), encodedMVCCKeyPrefixLength(roachpb.Key(tc.key)))

			// Test encode/decodeMVCCTimestampSuffix too, since we can trivially do so.
			expectTS, err := hex.DecodeString(tc.encoded[2*len(tc.key)+2:])
			require.NoError(t, err)
			if len(expectTS) == 0 {
				expectTS = nil
			}

			encodedTS := EncodeMVCCTimestampSuffix(tc.ts)
			require.Equal(t, expectTS, encodedTS)
			require.Equal(t, len(encodedTS), encodedMVCCTimestampSuffixLength(tc.ts))

			decodedTS, err := decodeMVCCTimestampSuffix(encodedTS)
			require.NoError(t, err)
			require.Equal(t, tc.ts, decodedTS)

			// Test encode/decodeMVCCTimestamp as well, for completeness.
			if len(expectTS) > 0 {
				expectTS = expectTS[:len(expectTS)-1]
			}

			encodedTS = encodeMVCCTimestamp(tc.ts)
			require.Equal(t, expectTS, encodedTS)
			require.Equal(t, len(encodedTS), encodedMVCCTimestampLength(tc.ts))

			decodedTS, err = decodeMVCCTimestamp(encodedTS)
			require.NoError(t, err)
			require.Equal(t, tc.ts, decodedTS)
		})
	}
}

func TestDecodeUnnormalizedMVCCKey(t *testing.T) {
	defer leaktest.AfterTest(t)()

	testcases := map[string]struct {
		encoded       string // hex-encoded
		expected      MVCCKey
		equalToNormal bool
	}{
		"zero logical": {
			encoded:       "666f6f0016cf10bc05055741000000000d",
			expected:      MVCCKey{Key: []byte("foo"), Timestamp: hlc.Timestamp{WallTime: 1643550788737652545, Logical: 0}},
			equalToNormal: true,
		},
		"zero walltime and logical": {
			encoded:  "666f6f000000000000000000000000000d",
			expected: MVCCKey{Key: []byte("foo"), Timestamp: hlc.Timestamp{WallTime: 0, Logical: 0}},
			// We could normalize this form in EngineKeyEqual and EngineKeyCompare,
			// but doing so is not worth losing the fast-path byte comparison between
			// keys that only contain (at most) a walltime.
			equalToNormal: false,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			encoded, err := hex.DecodeString(tc.encoded)
			require.NoError(t, err)

			decoded, err := DecodeMVCCKey(encoded)
			require.NoError(t, err)
			require.Equal(t, tc.expected, decoded)

			// Re-encode the key into its normal form.
			reencoded := EncodeMVCCKey(decoded)
			require.NotEqual(t, encoded, reencoded)
			require.Equal(t, tc.equalToNormal, EngineKeyEqual(encoded, reencoded))
			require.Equal(t, tc.equalToNormal, EngineKeyCompare(encoded, reencoded) == 0)
		})
	}
}

func TestDecodeMVCCKeyErrors(t *testing.T) {
	defer leaktest.AfterTest(t)()

	testcases := map[string]struct {
		encoded   string // hex-encoded
		expectErr string
	}{
		"empty input":                     {"", "invalid encoded mvcc key: "},
		"lone length suffix":              {"01", "invalid encoded mvcc key: "},
		"invalid timestamp length":        {"ab00ffff03", "invalid encoded mvcc key: ab00ffff03 bad timestamp ffff"},
		"invalid timestamp length suffix": {"ab00ffffffffffffffff0f", "invalid encoded mvcc key: ab00ffffffffffffffff0f"},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			encoded, err := hex.DecodeString(tc.encoded)
			require.NoError(t, err)

			_, err = DecodeMVCCKey(encoded)
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.expectErr)
		})
	}
}

func TestDecodeMVCCTimestampSuffixErrors(t *testing.T) {
	defer leaktest.AfterTest(t)()

	testcases := map[string]struct {
		encoded   string // hex-encoded
		expectErr string
	}{
		"invalid length":        {"ffff03", "bad timestamp ffff"},
		"invalid length suffix": {"ffffffffffffffff0f", "bad timestamp: found length suffix 15, actual length 9"},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			encoded, err := hex.DecodeString(tc.encoded)
			require.NoError(t, err)

			_, err = decodeMVCCTimestampSuffix(encoded)
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.expectErr)
		})
	}
}

var benchmarkEncodeMVCCKeyResult []byte

func BenchmarkEncodeMVCCKey(b *testing.B) {
	keys := map[string][]byte{
		"empty": {},
		"short": []byte("foo"),
		"long":  bytes.Repeat([]byte{1}, 4096),
	}
	timestamps := map[string]hlc.Timestamp{
		"empty":            {},
		"walltime":         {WallTime: 1643550788737652545},
		"walltime+logical": {WallTime: 1643550788737652545, Logical: 4096},
		"all":              {WallTime: 1643550788737652545, Logical: 4096, Synthetic: true},
	}
	buf := make([]byte, 0, 65536)
	for keyDesc, key := range keys {
		for tsDesc, ts := range timestamps {
			mvccKey := MVCCKey{Key: key, Timestamp: ts}
			b.Run(fmt.Sprintf("key=%s/ts=%s", keyDesc, tsDesc), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					EncodeMVCCKeyToBuf(buf, mvccKey)
				}
			})
		}
	}
	benchmarkEncodeMVCCKeyResult = buf // avoid compiler optimizing away function call
}

var benchmarkDecodeMVCCKeyResult MVCCKey

func BenchmarkDecodeMVCCKey(b *testing.B) {
	keys := map[string][]byte{
		"empty": {},
		"short": []byte("foo"),
		"long":  bytes.Repeat([]byte{1}, 4096),
	}
	timestamps := map[string]hlc.Timestamp{
		"empty":            {},
		"walltime":         {WallTime: 1643550788737652545},
		"walltime+logical": {WallTime: 1643550788737652545, Logical: 4096},
		"all":              {WallTime: 1643550788737652545, Logical: 4096, Synthetic: true},
	}
	var mvccKey MVCCKey
	var err error
	for keyDesc, key := range keys {
		for tsDesc, ts := range timestamps {
			encoded := EncodeMVCCKey(MVCCKey{Key: key, Timestamp: ts})
			b.Run(fmt.Sprintf("key=%s/ts=%s", keyDesc, tsDesc), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					mvccKey, err = DecodeMVCCKey(encoded)
					if err != nil { // for performance
						require.NoError(b, err)
					}
				}
			})
		}
	}
	benchmarkDecodeMVCCKeyResult = mvccKey // avoid compiler optimizing away function call
}
