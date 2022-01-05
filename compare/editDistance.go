package compare

import (
	"errors"
	"strings"
)

//DamerauLevenshtein computes the Damerau-Levenshtein distance between two
// strings. The returned value - distance - is the number of insertions,
// deletions, substitutions, and transpositions it takes to transform one
// string (s1) into another (s2). Each step in the transformation "costs"
// one distance point. It is similar to the Optimal String Alignment,
// algorithm, but is more complex because it allows multiple edits on
// substrings.
//
//Based on: https://github.com/antzucaro/matchr/blob/b04723ef80f0/damerau_levenshtein.go#L15
func DamerauLevenshtein(s1 string, s2 string, caseSensitive bool) (distance int) {

	if !caseSensitive {
		s1 = strings.ToLower(s1)
		s2 = strings.ToLower(s2)
	}

	// index by code point, not byte
	r1 := []rune(s1)
	r2 := []rune(s2)

	// the maximum possible distance
	inf := len(r1) + len(r2)

	// if one string is blank, we needs insertions
	// for all characters in the other one
	if len(r1) == 0 {
		return len(r2)
	}

	if len(r2) == 0 {
		return len(r1)
	}

	// construct the edit-tracking matrix
	matrix := make([][]int, len(r1))
	for i := range matrix {
		matrix[i] = make([]int, len(r2))
	}

	// seen characters
	seenRunes := make(map[rune]int)

	if r1[0] != r2[0] {
		matrix[0][0] = 1
	}

	seenRunes[r1[0]] = 0
	for i := 1; i < len(r1); i++ {
		deleteDist := matrix[i-1][0] + 1
		insertDist := (i+1)*1 + 1
		var matchDist int
		if r1[i] == r2[0] {
			matchDist = i
		} else {
			matchDist = i + 1
		}
		matrix[i][0] = minInt(minInt(deleteDist, insertDist), matchDist)
	}

	for j := 1; j < len(r2); j++ {
		deleteDist := (j + 1) * 2
		insertDist := matrix[0][j-1] + 1
		var matchDist int
		if r1[0] == r2[j] {
			matchDist = j
		} else {
			matchDist = j + 1
		}

		matrix[0][j] = minInt(minInt(deleteDist, insertDist), matchDist)
	}

	for i := 1; i < len(r1); i++ {
		var maxSrcMatchIndex int
		if r1[i] == r2[0] {
			maxSrcMatchIndex = 0
		} else {
			maxSrcMatchIndex = -1
		}

		for j := 1; j < len(r2); j++ {
			swapIndex, ok := seenRunes[r2[j]]
			jSwap := maxSrcMatchIndex
			deleteDist := matrix[i-1][j] + 1
			insertDist := matrix[i][j-1] + 1
			matchDist := matrix[i-1][j-1]
			if r1[i] != r2[j] {
				matchDist += 1
			} else {
				maxSrcMatchIndex = j
			}

			// for transpositions
			var swapDist int
			if ok && jSwap != -1 {
				iSwap := swapIndex
				var preSwapCost int
				if iSwap == 0 && jSwap == 0 {
					preSwapCost = 0
				} else {
					preSwapCost = matrix[maxInt(0, iSwap-1)][maxInt(0, jSwap-1)]
				}
				swapDist = i + j + preSwapCost - iSwap - jSwap - 1
			} else {
				swapDist = inf
			}
			matrix[i][j] = minInt(minInt(minInt(deleteDist, insertDist), matchDist), swapDist)
		}
		seenRunes[r1[i]] = i
	}

	return matrix[len(r1)-1][len(r2)-1]
}

// Levenshtein computes the Levenshtein distance between two
// strings. The returned value - distance - is the number of insertions,
// deletions, and substitutions it takes to transform one
// string (s1) into another (s2). Each step in the transformation "costs"
// one distance point.
//
//Based on: https://github.com/antzucaro/matchr/blob/b04723ef80f01d4977524a1de94197868f4c5907/levenshtein.go
func Levenshtein(s1 string, s2 string, caseSensitive bool) (distance int) {

	if !caseSensitive {
		s1 = strings.ToLower(s1)
		s2 = strings.ToLower(s2)
	}

	// index by code point, not byte
	r1 := []rune(s1)
	r2 := []rune(s2)

	rows := len(r1) + 1
	cols := len(r2) + 1

	var d1 int
	var d2 int
	var d3 int
	var i int
	var j int
	dist := make([]int, rows*cols)

	for i = 0; i < rows; i++ {
		dist[i*cols] = i
	}

	for j = 0; j < cols; j++ {
		dist[j] = j
	}

	for j = 1; j < cols; j++ {
		for i = 1; i < rows; i++ {
			if r1[i-1] == r2[j-1] {
				dist[(i*cols)+j] = dist[((i-1)*cols)+(j-1)]
			} else {
				d1 = dist[((i-1)*cols)+j] + 1
				d2 = dist[(i*cols)+(j-1)] + 1
				d3 = dist[((i-1)*cols)+(j-1)] + 1

				dist[(i*cols)+j] = minInt(d1, minInt(d2, d3))
			}
		}
	}

	distance = dist[(cols*rows)-1]

	return
}

// Hamming computes the Hamming distance between two equal-length strings.
// This is the number of times the two strings differ between characters at
// the same index. This implementation is based off of the algorithm
// description found at http://en.wikipedia.org/wiki/Hamming_distance.
//
//Based on: https://github.com/antzucaro/matchr/blob/b04723ef80f01d4977524a1de94197868f4c5907/hamming.go
func Hamming(s1 string, s2 string, caseSensitive bool) (distance int, err error) {
	if !caseSensitive {
		s1 = strings.ToLower(s1)
		s2 = strings.ToLower(s2)
	}

	// index by code point, not byte
	r1 := []rune(s1)
	r2 := []rune(s2)

	if len(r1) != len(r2) {
		err = errors.New("Hamming distance of different sized strings.")
		return
	}

	for i, v := range r1 {
		if r2[i] != v {
			distance += 1
		}
	}
	return
}
