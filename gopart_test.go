package gopart

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	var partitionTests = []struct {
		collectionLen int
		partitionSize int
		step          int
		expRanges     []IdxRange
	}{
		// evenly split
		{9, 3, 3, []IdxRange{{0, 3}, {3, 6}, {6, 9}}},
		// uneven partition
		{13, 5, 5, []IdxRange{{0, 5}, {5, 10}, {10, 13}}},
		// large partition size
		{13, 19, 19, []IdxRange{{0, 13}}},
		// zero partition size
		{7, 0, 0, nil},
		// negative partition size
		{7, -4, -4, nil},
		// same size
		{19, 19, 19, []IdxRange{{0, 19}}},
		// zero collection length
		{0, 19, 19, nil},
		// overlapping partitions with step 1
		{5, 3, 1, []IdxRange{{0, 3}, {1, 4}, {2, 5}}},
		// overlapping partitions with step 2
		{8, 3, 2, []IdxRange{{0, 3}, {2, 5}, {4, 7}, {6, 8}}},
		// partitions with gaps (remainder)
		{5, 3, 4, []IdxRange{{0, 3}, {4, 5}}},
		// partitions with gaps (outside)
		{5, 3, 6, []IdxRange{{0, 3}}},
	}

	for _, tt := range partitionTests {
		actChannel := PartitionWithStep(tt.collectionLen, tt.partitionSize, tt.step)
		var actRange []IdxRange
		for idxRange := range actChannel {
			actRange = append(actRange, idxRange)
		}

		if !reflect.DeepEqual(actRange, tt.expRanges) {
			t.Errorf("expected %d, actual %d", tt.expRanges, actRange)
		}
	}
}

func Benchmark100kPartitions(b *testing.B) { benchmarkPartition(17*1e5+11, b) }
func Benchmark10kPartitions(b *testing.B)  { benchmarkPartition(17*1e4+11, b) }
func Benchmark1kPartitions(b *testing.B)   { benchmarkPartition(17*1e3+11, b) }
func Benchmark100Partitions(b *testing.B)  { benchmarkPartition(17*1e2+11, b) }
func Benchmark10Partitions(b *testing.B)   { benchmarkPartition(17*1e1+11, b) }

func benchmarkPartition(length int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		for range Partition(length, 17) {
		}
	}
}
