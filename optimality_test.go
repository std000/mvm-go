package mwm

import (
	"math/rand/v2"
	"testing"
)

// bruteForce returns the maximum total weight of any matching on the graph
// given as a weight matrix (0 = no edge). If perfect is true, only perfect
// matchings are considered.
func bruteForce(w [][]int64, used []bool, perfect bool) int64 {
	i := 0
	for i < len(w) && used[i] {
		i++
	}
	if i == len(w) {
		return 0
	}
	const impossible = int64(-1 << 62)
	used[i] = true
	best := impossible
	if !perfect {
		best = bruteForce(w, used, perfect) // leave i unmatched
	}
	for j := i + 1; j < len(w); j++ {
		if !used[j] && w[i][j] != 0 {
			used[j] = true
			if rest := bruteForce(w, used, perfect); rest != impossible {
				best = max(best, w[i][j]+rest)
			}
			used[j] = false
		}
	}
	used[i] = false
	return best
}

func matchingWeight(t *testing.T, w [][]int64, pairs []Pair) int64 {
	t.Helper()
	seen := map[int64]bool{}
	var total int64
	for _, p := range pairs {
		if seen[p.First] || seen[p.Second] {
			t.Fatalf("vertex matched twice in %v", pairs)
		}
		seen[p.First], seen[p.Second] = true, true
		a, b := min(p.First, p.Second), max(p.First, p.Second)
		if w[a][b] == 0 {
			t.Fatalf("pair %v is not an edge", p)
		}
		total += w[a][b]
	}
	return total
}

// Regression: this graph previously produced {0-5, 1-2, 3-4} (weight 2308)
// instead of the optimum {0-1, 2-3, 4-5} (weight 2312).
func TestSuboptimalRegression(t *testing.T) {
	weights := []int64{836, 822, 1, 624, 636, 836, 626, 1, 624, 640, 626, 1, 836, 822, 836}
	n := 6
	w := make([][]int64, n)
	for i := range w {
		w[i] = make([]int64, n)
	}
	var edges []GraphEdge
	k := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			w[i][j] = weights[k]
			edges = append(edges, GraphEdge{Node1: int64(i), Node2: int64(j), Weight: weights[k]})
			k++
		}
	}
	got := matchingWeight(t, w, NewMaximumWeightedMatching().MaxWeightMatching(edges, true))
	if got != 2312 {
		t.Fatalf("weight %d, want 2312", got)
	}
}

// Compares against exhaustive search on random graphs, both complete and
// sparse, with and without maxCardinality.
func TestRandomGraphsOptimal(t *testing.T) {
	rng := rand.New(rand.NewPCG(3, 4))
	for trial := 0; trial < 5000; trial++ {
		n := 2 + rng.IntN(11)
		maxW := 1 + rng.IntN(1000)
		density := 0.3 + 0.7*rng.Float64()
		maxCard := rng.IntN(2) == 0

		w := make([][]int64, n)
		for i := range w {
			w[i] = make([]int64, n)
		}
		var edges []GraphEdge
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if rng.Float64() < density {
					x := int64(1 + rng.IntN(maxW))
					w[i][j] = x
					edges = append(edges, GraphEdge{Node1: int64(i), Node2: int64(j), Weight: x})
				}
			}
		}

		res := NewMaximumWeightedMatching().MaxWeightMatching(edges, maxCard)
		got := matchingWeight(t, w, res)

		var want int64
		if maxCard {
			// Optimum among matchings of maximum cardinality: find that
			// cardinality, then the best weight at it, by brute force
			// over a transformed problem (large bonus per edge).
			const bonus = 1 << 40
			wb := make([][]int64, n)
			for i := range wb {
				wb[i] = make([]int64, n)
				for j := range wb[i] {
					if w[i][j] != 0 {
						wb[i][j] = w[i][j] + bonus
					}
				}
			}
			best := bruteForce(wb, make([]bool, n), false)
			want = best % bonus
			if card := best / bonus; int64(len(res)) != card {
				t.Fatalf("trial %d: cardinality %d, want %d", trial, len(res), card)
			}
		} else {
			want = bruteForce(w, make([]bool, n), false)
		}
		if got != want {
			t.Fatalf("trial %d (n=%d, maxCard=%v): weight %d, optimum %d", trial, n, maxCard, got, want)
		}
	}
}
