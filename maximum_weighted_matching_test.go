package mwm

import (
	"fmt"
	"reflect"
	"testing"
)

// Helper function for formatting
func init() {
	// Import fmt for use in tests
}

// maxWeightMatchingList - helper function for tests
// Converts MaxWeightMatching result to list format as in the original Kotlin code
func maxWeightMatchingList(matcher *MaximumWeightedMatching, edges []GraphEdge, maxCardinality bool) []int64 {
	pairs := matcher.MaxWeightMatching(edges, maxCardinality)

	if len(pairs) == 0 {
		return []int64{}
	}

	// Determine the maximum node number to create array of needed size
	maxNode := int64(-1)
	for _, edge := range edges {
		if edge.Node1 > maxNode {
			maxNode = edge.Node1
		}
		if edge.Node2 > maxNode {
			maxNode = edge.Node2
		}
	}

	// Create result array initialized with -1
	result := make([]int64, maxNode+1)
	for i := range result {
		result[i] = -1
	}

	// Fill matchings
	for _, pair := range pairs {
		result[pair.First] = pair.Second
		result[pair.Second] = pair.First
	}

	return result
}

func Test10(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test11(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 0, Node2: 1, Weight: 1},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{1, 0}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test12(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 10},
		{Node1: 2, Node2: 3, Weight: 11},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, -1, 3, 2}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test13(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 5},
		{Node1: 2, Node2: 3, Weight: 11},
		{Node1: 3, Node2: 4, Weight: 5},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, -1, 3, 2, -1}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test14Maxcard(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 5},
		{Node1: 2, Node2: 3, Weight: 11},
		{Node1: 3, Node2: 4, Weight: 5},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, true)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 2, 1, 4, 3}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test16Negative(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 2},
		{Node1: 1, Node2: 3, Weight: -2},
		{Node1: 2, Node2: 3, Weight: 1},
		{Node1: 2, Node2: 4, Weight: -1},
		{Node1: 3, Node2: 4, Weight: -6},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 2, 1, -1, -1}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test16NegativeMaxcard(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 2},
		{Node1: 1, Node2: 3, Weight: -2},
		{Node1: 2, Node2: 3, Weight: 1},
		{Node1: 2, Node2: 4, Weight: -1},
		{Node1: 3, Node2: 4, Weight: -6},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, true)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 3, 4, 1, 2}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test20SBlossom1(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 8},
		{Node1: 1, Node2: 3, Weight: 9},
		{Node1: 2, Node2: 3, Weight: 10},
		{Node1: 3, Node2: 4, Weight: 7},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 2, 1, 4, 3}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test20SBlossom2(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 8},
		{Node1: 1, Node2: 3, Weight: 9},
		{Node1: 2, Node2: 3, Weight: 10},
		{Node1: 3, Node2: 4, Weight: 7},
		{Node1: 1, Node2: 6, Weight: 5},
		{Node1: 4, Node2: 5, Weight: 6},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 6, 3, 2, 5, 4, 1}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test21TBlossom1(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 9},
		{Node1: 1, Node2: 3, Weight: 8},
		{Node1: 2, Node2: 3, Weight: 10},
		{Node1: 1, Node2: 4, Weight: 5},
		{Node1: 4, Node2: 5, Weight: 4},
		{Node1: 1, Node2: 6, Weight: 3},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 6, 3, 2, 5, 4, 1}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test21TBlossom2(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 9},
		{Node1: 1, Node2: 3, Weight: 8},
		{Node1: 2, Node2: 3, Weight: 10},
		{Node1: 1, Node2: 4, Weight: 5},
		{Node1: 4, Node2: 5, Weight: 4},
		{Node1: 1, Node2: 6, Weight: 4},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 6, 3, 2, 5, 4, 1}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test21TBlossom3(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 9},
		{Node1: 1, Node2: 3, Weight: 8},
		{Node1: 2, Node2: 3, Weight: 10},
		{Node1: 1, Node2: 4, Weight: 5},
		{Node1: 4, Node2: 5, Weight: 4},
		{Node1: 3, Node2: 6, Weight: 4},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 2, 1, 6, 5, 4, 3}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test22SNest(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 9},
		{Node1: 1, Node2: 3, Weight: 9},
		{Node1: 2, Node2: 3, Weight: 10},
		{Node1: 2, Node2: 4, Weight: 8},
		{Node1: 3, Node2: 5, Weight: 8},
		{Node1: 4, Node2: 5, Weight: 10},
		{Node1: 5, Node2: 6, Weight: 6},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 3, 4, 1, 2, 6, 5}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test23SRelabelNest(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 10},
		{Node1: 1, Node2: 7, Weight: 10},
		{Node1: 2, Node2: 3, Weight: 12},
		{Node1: 3, Node2: 4, Weight: 20},
		{Node1: 3, Node2: 5, Weight: 20},
		{Node1: 4, Node2: 5, Weight: 25},
		{Node1: 5, Node2: 6, Weight: 10},
		{Node1: 6, Node2: 7, Weight: 10},
		{Node1: 7, Node2: 8, Weight: 8},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 2, 1, 4, 3, 6, 5, 8, 7}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test24SNestExpand(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 8},
		{Node1: 1, Node2: 3, Weight: 8},
		{Node1: 2, Node2: 3, Weight: 10},
		{Node1: 2, Node2: 4, Weight: 12},
		{Node1: 3, Node2: 5, Weight: 12},
		{Node1: 4, Node2: 5, Weight: 14},
		{Node1: 4, Node2: 6, Weight: 12},
		{Node1: 5, Node2: 7, Weight: 12},
		{Node1: 6, Node2: 7, Weight: 14},
		{Node1: 7, Node2: 8, Weight: 12},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 2, 1, 5, 6, 3, 4, 8, 7}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test25STExpand(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 23},
		{Node1: 1, Node2: 5, Weight: 22},
		{Node1: 1, Node2: 6, Weight: 15},
		{Node1: 2, Node2: 3, Weight: 25},
		{Node1: 3, Node2: 4, Weight: 22},
		{Node1: 4, Node2: 5, Weight: 25},
		{Node1: 4, Node2: 8, Weight: 14},
		{Node1: 5, Node2: 7, Weight: 13},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 6, 3, 2, 8, 7, 1, 5, 4}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test26SNestTExpand(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 19},
		{Node1: 1, Node2: 3, Weight: 20},
		{Node1: 1, Node2: 8, Weight: 8},
		{Node1: 2, Node2: 3, Weight: 25},
		{Node1: 2, Node2: 4, Weight: 18},
		{Node1: 3, Node2: 5, Weight: 18},
		{Node1: 4, Node2: 5, Weight: 13},
		{Node1: 4, Node2: 7, Weight: 7},
		{Node1: 5, Node2: 6, Weight: 7},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 8, 3, 2, 7, 6, 5, 4, 1}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test30TNastyExpand(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 45},
		{Node1: 1, Node2: 5, Weight: 45},
		{Node1: 2, Node2: 3, Weight: 50},
		{Node1: 3, Node2: 4, Weight: 45},
		{Node1: 4, Node2: 5, Weight: 50},
		{Node1: 1, Node2: 6, Weight: 30},
		{Node1: 3, Node2: 9, Weight: 35},
		{Node1: 4, Node2: 8, Weight: 35},
		{Node1: 5, Node2: 7, Weight: 26},
		{Node1: 9, Node2: 10, Weight: 5},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 6, 3, 2, 8, 7, 1, 5, 4, 10, 9}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test31TNasty2Expand(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 45},
		{Node1: 1, Node2: 5, Weight: 45},
		{Node1: 2, Node2: 3, Weight: 50},
		{Node1: 3, Node2: 4, Weight: 45},
		{Node1: 4, Node2: 5, Weight: 50},
		{Node1: 1, Node2: 6, Weight: 30},
		{Node1: 3, Node2: 9, Weight: 35},
		{Node1: 4, Node2: 8, Weight: 26},
		{Node1: 5, Node2: 7, Weight: 40},
		{Node1: 9, Node2: 10, Weight: 5},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 6, 3, 2, 8, 7, 1, 5, 4, 10, 9}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test32TExpandLeastslack(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 45},
		{Node1: 1, Node2: 5, Weight: 45},
		{Node1: 2, Node2: 3, Weight: 50},
		{Node1: 3, Node2: 4, Weight: 45},
		{Node1: 4, Node2: 5, Weight: 50},
		{Node1: 1, Node2: 6, Weight: 30},
		{Node1: 3, Node2: 9, Weight: 35},
		{Node1: 4, Node2: 8, Weight: 28},
		{Node1: 5, Node2: 7, Weight: 26},
		{Node1: 9, Node2: 10, Weight: 5},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 6, 3, 2, 8, 7, 1, 5, 4, 10, 9}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test33NestTNastyExpand(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 45},
		{Node1: 1, Node2: 7, Weight: 45},
		{Node1: 2, Node2: 3, Weight: 50},
		{Node1: 3, Node2: 4, Weight: 45},
		{Node1: 4, Node2: 5, Weight: 95},
		{Node1: 4, Node2: 6, Weight: 94},
		{Node1: 5, Node2: 6, Weight: 94},
		{Node1: 6, Node2: 7, Weight: 50},
		{Node1: 1, Node2: 8, Weight: 30},
		{Node1: 3, Node2: 11, Weight: 35},
		{Node1: 5, Node2: 9, Weight: 36},
		{Node1: 7, Node2: 10, Weight: 26},
		{Node1: 11, Node2: 12, Weight: 5},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 8, 3, 2, 6, 9, 4, 10, 1, 5, 7, 12, 11}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

func Test34NestRelabelExpand(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 1, Node2: 2, Weight: 40},
		{Node1: 1, Node2: 3, Weight: 40},
		{Node1: 2, Node2: 3, Weight: 60},
		{Node1: 2, Node2: 4, Weight: 55},
		{Node1: 3, Node2: 5, Weight: 55},
		{Node1: 4, Node2: 5, Weight: 50},
		{Node1: 1, Node2: 8, Weight: 15},
		{Node1: 5, Node2: 7, Weight: 30},
		{Node1: 7, Node2: 6, Weight: 10},
		{Node1: 8, Node2: 10, Weight: 10},
		{Node1: 4, Node2: 9, Weight: 30},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println(weightMatchList)
	expected := []int64{-1, 2, 1, 5, 9, 3, 7, 6, 10, 4, 8}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

// TestStarGraph - test of star graph (central vertex connected to all others)
func TestStarGraph(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 0, Node2: 1, Weight: 10},
		{Node1: 0, Node2: 2, Weight: 20},
		{Node1: 0, Node2: 3, Weight: 30},
		{Node1: 0, Node2: 4, Weight: 40},
		{Node1: 0, Node2: 5, Weight: 50},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println("Star graph:", weightMatchList)
	// In a star we can only take one edge with maximum weight
	expected := []int64{5, -1, -1, -1, -1, 0}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

// TestCompleteGraph4 - test of complete graph on 4 vertices
func TestCompleteGraph4(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 0, Node2: 1, Weight: 1},
		{Node1: 0, Node2: 2, Weight: 2},
		{Node1: 0, Node2: 3, Weight: 3},
		{Node1: 1, Node2: 2, Weight: 4},
		{Node1: 1, Node2: 3, Weight: 5},
		{Node1: 2, Node2: 3, Weight: 6},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println("Complete graph 4:", weightMatchList)
	// Optimal matching: (0,1) weight 1 + (2,3) weight 6 = 7
	expected := []int64{1, 0, 3, 2}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

// TestCycleGraph6 - test of cycle graph on 6 vertices
func TestCycleGraph6(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 0, Node2: 1, Weight: 10},
		{Node1: 1, Node2: 2, Weight: 15},
		{Node1: 2, Node2: 3, Weight: 20},
		{Node1: 3, Node2: 4, Weight: 25},
		{Node1: 4, Node2: 5, Weight: 30},
		{Node1: 5, Node2: 0, Weight: 5},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println("Cycle graph 6:", weightMatchList)
	// Optimal matching: (0,1) weight 10 + (2,3) weight 20 + (4,5) weight 30 = 60
	expected := []int64{1, 0, 3, 2, 5, 4}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

// TestDisconnectedComponents - test of graph with several separate components
func TestDisconnectedComponents(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		// First component: triangle
		{Node1: 0, Node2: 1, Weight: 10},
		{Node1: 1, Node2: 2, Weight: 20},
		{Node1: 0, Node2: 2, Weight: 15},
		// Second component: simple edge
		{Node1: 3, Node2: 4, Weight: 30},
		// Third component: another edge
		{Node1: 5, Node2: 6, Weight: 25},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println("Disconnected components:", weightMatchList)
	expected := []int64{-1, 2, 1, 4, 3, 6, 5}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

// TestEqualWeights - test with equal weights of all edges
func TestEqualWeights(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 0, Node2: 1, Weight: 5},
		{Node1: 1, Node2: 2, Weight: 5},
		{Node1: 2, Node2: 3, Weight: 5},
		{Node1: 3, Node2: 0, Weight: 5},
		{Node1: 0, Node2: 2, Weight: 5},
		{Node1: 1, Node2: 3, Weight: 5},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println("Equal weights:", weightMatchList)
	// Any maximum matching should have 2 edges
	nonMinusOneCount := 0
	for _, val := range weightMatchList {
		if val != -1 {
			nonMinusOneCount++
		}
	}
	if nonMinusOneCount != 4 { // 2 pairs = 4 non-empty elements
		t.Errorf("Expected 4 matched vertices, got %d", nonMinusOneCount)
	}
}

// TestZeroWeights - test with zero weights
func TestZeroWeights(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 0, Node2: 1, Weight: 0},
		{Node1: 1, Node2: 2, Weight: 0},
		{Node1: 2, Node2: 3, Weight: 0},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println("Zero weights:", weightMatchList)
	// With zero weights algorithm can find matching with total weight 0
	expected := []int64{1, 0, 3, 2}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

// TestSingleEdge - test with single edge
func TestSingleEdge(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 5, Node2: 10, Weight: 100},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println("Single edge:", weightMatchList)
	expected := []int64{-1, -1, -1, -1, -1, 10, -1, -1, -1, -1, 5}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

// TestLargeWeights - test with very large weights
func TestLargeWeights(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 0, Node2: 1, Weight: 1000000},
		{Node1: 1, Node2: 2, Weight: 2000000},
		{Node1: 0, Node2: 2, Weight: 1500000},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println("Large weights:", weightMatchList)
	// Should choose edge with weight 2000000
	expected := []int64{-1, 2, 1}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

// TestPath - test of path (chain)
func TestPath(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 0, Node2: 1, Weight: 10},
		{Node1: 1, Node2: 2, Weight: 20},
		{Node1: 2, Node2: 3, Weight: 15},
		{Node1: 3, Node2: 4, Weight: 25},
		{Node1: 4, Node2: 5, Weight: 12},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println("Path:", weightMatchList)
	// Optimal matching: (1,2) weight 20 + (3,4) weight 25 = 45
	expected := []int64{-1, 2, 1, 4, 3, -1}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}

// TestMaxCardinalityVsMaxWeight - comparison of maximum cardinality and maximum weight
func TestMaxCardinalityVsMaxWeight(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 0, Node2: 1, Weight: 100}, // Very large weight
		{Node1: 2, Node2: 3, Weight: 1},   // Small weight
		{Node1: 4, Node2: 5, Weight: 1},   // Small weight
	}

	// Maximum weight
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println("Max weight:", weightMatchList)
	expectedWeight := []int64{1, 0, 3, 2, 5, 4}
	if !reflect.DeepEqual(weightMatchList, expectedWeight) {
		t.Errorf("Expected %v, got %v", expectedWeight, weightMatchList)
	}

	// Maximum cardinality
	cardMatchList := maxWeightMatchingList(matcher, edges, true)
	fmt.Println("Max cardinality:", cardMatchList)
	expectedCard := []int64{1, 0, 3, 2, 5, 4}
	if !reflect.DeepEqual(cardMatchList, expectedCard) {
		t.Errorf("Expected %v, got %v", expectedCard, cardMatchList)
	}
}

// TestComplexBlossom - more complex case with blossom
func TestComplexBlossom(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 0, Node2: 1, Weight: 50},
		{Node1: 1, Node2: 2, Weight: 60},
		{Node1: 2, Node2: 0, Weight: 55},
		{Node1: 0, Node2: 3, Weight: 40},
		{Node1: 1, Node2: 4, Weight: 45},
		{Node1: 2, Node2: 5, Weight: 35},
		{Node1: 3, Node2: 6, Weight: 30},
		{Node1: 4, Node2: 7, Weight: 25},
		{Node1: 5, Node2: 8, Weight: 20},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println("Complex blossom:", weightMatchList)
	// Expect algorithm to find optimal solution
	// Check that all matchings are correct
	for i, val := range weightMatchList {
		if val != -1 && weightMatchList[val] != int64(i) {
			t.Errorf("Inconsistent matching at vertex %d: points to %d, but %d doesn't point back", i, val, val)
		}
	}
}

// TestMixedPositiveNegative - test with mixed positive and negative weights
func TestMixedPositiveNegative(t *testing.T) {
	matcher := NewMaximumWeightedMatching()
	edges := []GraphEdge{
		{Node1: 0, Node2: 1, Weight: 10},
		{Node1: 1, Node2: 2, Weight: -5},
		{Node1: 2, Node2: 3, Weight: 15},
		{Node1: 0, Node2: 3, Weight: 8},
		{Node1: 0, Node2: 2, Weight: -3},
		{Node1: 1, Node2: 3, Weight: 12},
	}
	weightMatchList := maxWeightMatchingList(matcher, edges, false)
	fmt.Println("Mixed positive negative:", weightMatchList)
	expected := []int64{1, 0, 3, 2}
	if !reflect.DeepEqual(weightMatchList, expected) {
		t.Errorf("Expected %v, got %v", expected, weightMatchList)
	}
}
