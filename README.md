# Maximum Weighted Matching (Go)

Implementation of the maximum weighted matching algorithm in Go, rewritten from the Kotlin version.

## Description

The maximum weighted matching algorithm finds a matching in a graph with the maximum total weight of edges. This is an implementation of Edmonds' Blossom algorithm.

## Features

- **Maximum weighted matching**: Finds a matching with maximum total weight
- **Maximum cardinality**: Can also find a matching with maximum number of edges
- **Negative weights support**: Algorithm works correctly with negative edge weights
- **Debug mode**: Ability to enable detailed algorithm execution output

## Installation

```bash
go get github.com/std000/mvm-go
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/std000/mvm-go"
)

func main() {
    // Create graph with edges
    edges := []mwm.GraphEdge{
        {Node1: 0, Node2: 1, Weight: 10},
        {Node1: 1, Node2: 2, Weight: 11},
        {Node1: 2, Node2: 3, Weight: 12},
        {Node1: 0, Node2: 3, Weight: 13},
    }
    
    // Create algorithm instance
    matcher := mwm.NewMaximumWeightedMatching()
    
    // Find maximum weighted matching
    result := matcher.MaxWeightMatching(edges, false)
    
    fmt.Printf("Result: %v\n", result)
}
```

### Maximum Cardinality Mode

```go
// Find matching with maximum number of edges
result := matcher.MaxWeightMatching(edges, true)
```

### Debug Mode

```go
matcher := mwm.NewMaximumWeightedMatching()
matcher.DebugMode = true // Enable detailed output
result := matcher.MaxWeightMatching(edges, false)
```

## API

### Data Types

#### GraphEdge
```go
type GraphEdge struct {
    Node1  int64  // First vertex
    Node2  int64  // Second vertex  
    Weight int64  // Edge weight
}
```

#### Pair
```go
type Pair struct {
    First  int64  // First vertex in pair
    Second int64  // Second vertex in pair
}
```

#### MaximumWeightedMatching
```go
type MaximumWeightedMatching struct {
    DebugMode bool  // Debug mode flag
}
```

### Methods

#### NewMaximumWeightedMatching
```go
func NewMaximumWeightedMatching() *MaximumWeightedMatching
```
Creates a new algorithm instance.

#### MaxWeightMatching
```go
func (mwm *MaximumWeightedMatching) MaxWeightMatching(edges []GraphEdge, maxCardinality bool) []Pair
```
Finds the maximum weighted matching.

**Parameters:**
- `edges` - slice of graph edges
- `maxCardinality` - if `true`, algorithm seeks matching with maximum number of edges; if `false` - with maximum weight

**Returns:** slice of vertex pairs forming the matching.

## Complexity

- **Time complexity**: O(n³), where n is the number of vertices
- **Space complexity**: O(n²)

## Testing

Run tests:

```bash
go test
```

Run tests with verbose output:

```bash
go test -v
```

Run benchmarks:

```bash
go test -bench=.
```

## Example Output

```
Graph with edges:
  Edge 0: 0 -- 1 (weight: 10)
  Edge 1: 1 -- 2 (weight: 11)  
  Edge 2: 2 -- 3 (weight: 12)
  Edge 3: 0 -- 3 (weight: 13)

Result:
  Pair 1: 0 -- 3 (weight: 13)
  Pair 2: 1 -- 2 (weight: 11)

Total matching weight: 24
```

## Algorithm

The implementation is based on Edmonds' Blossom algorithm for maximum weighted matching:

1. **Dual problem**: Uses the dual problem of linear programming
2. **Blossoms**: Handles odd-length cycles by contracting them into "blossoms"
3. **Augmenting paths**: Finds augmenting paths to improve the matching
4. **Dual updates**: Updates dual variables to maintain optimality

## License

MIT License

## Original Implementation

Rewritten from Kotlin version: https://github.com/PizzaMarinara/MaximumWeightedMatching 