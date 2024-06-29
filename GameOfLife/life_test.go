package main

import "testing"

var zeroNeighbor = [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
var oneNeighbor = [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 0}}
var twoNeighbor = [][]int{{1, 0, 0}, {0, 1, 0}, {0, 1, 0}}
var threeNeighbor = [][]int{{1, 0, 1}, {0, 1, 0}, {0, 1, 0}}
var fourNeighbor = [][]int{{1, 1, 1}, {0, 1, 0}, {0, 1, 0}}

func TestNextState0Neighbor(t *testing.T) {
	want := 0
	result := calculateNextState(zeroNeighbor, 1, 1)

	if result != 0 {
		t.Fatalf("calculateNextState with 0 neighbor returned %d, but %d was expected", result, want)
	}
}
func TestNextState1Neighbor(t *testing.T) {
	want := 0
	result := calculateNextState(oneNeighbor, 1, 1)

	if result != 0 {
		t.Fatalf("calculateNextState with 1 neighbor returned %d, but %d was expected", result, want)
	}
}
func TestNextState2Neighbor(t *testing.T) {
	want := 1
	result := calculateNextState(twoNeighbor, 1, 1)

	if result != 1 {
		t.Fatalf("calculateNextState with 2 neighbor returned %d, but %d was expected", result, want)
	}
}
func TestNextState3Neighbor(t *testing.T) {
	want := 1
	result := calculateNextState(threeNeighbor, 1, 1)

	if result != 1 {
		t.Fatalf("calculateNextState with 3 neighbor returned %d, but %d was expected", result, want)
	}
}
func TestNextState4Neighbor(t *testing.T) {
	want := 0
	result := calculateNextState(fourNeighbor, 1, 1)

	if result != 0 {
		t.Fatalf("calculateNextState with 4 neighbor returned %d, but %d was expected", result, want)
	}
}
