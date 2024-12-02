package main

import (
    "testing"
    "os"
)

func TestSumSim(t *testing.T) {
    left := []int{3, 4, 2, 1, 3, 3}
    right := []int{4, 3, 5, 3, 9, 3}
    expectedSim := int64(31)

    sim, err := sumSim(left, right)
    if err != nil {
        t.Fatalf("sumSim returned an unexpected Error: %v", err)
    }

    if sim != expectedSim {
        t.Fatalf("sumSim: git %v, want %v", sim, expectedSim)
    }
}

func TestSumDiff(t *testing.T) {
	t.Run("Valid Inputs", func(t *testing.T) {
		left := []int{3, 4, 2, 1, 3, 3}
		right := []int{4, 3, 5, 3, 9, 3}
		expectedSum := int64(13)

		sum, err := sumDiff(left, right)
		if err != nil {
			t.Fatalf("sumDiff returned an unexpected error: %v", err)
		}

		if sum != expectedSum {
			t.Errorf("sumDiff: got %v, want %v", sum, expectedSum)
		}
	})

	t.Run("Different Lengths", func(t *testing.T) {
		left := []int{3, 4, 2, 1, 3}
		right := []int{4, 3, 5, 3, 9, 3}

		sum, err := sumDiff(left, right)

		if err == nil {
			t.Fatal("sumDiff: expected an error for lists of different lengths, but got nil")
		}

		// Validate the sum is zero
		if sum != 0 {
			t.Errorf("sumDiff: got %v, want 0 for lists of different lengths", sum)
		}
	})
}

func TestReadFile(t *testing.T) {
	input := `3 4
4 3
2 5
1 3
3 9
3 3`

	// Create a testinput file
	tmpfile, err := os.CreateTemp("", "testinput")
	if err != nil {
		t.Fatalf("Failed to create testinput file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up the file after the test


	if _, err := tmpfile.WriteString(input); err != nil {
		t.Fatalf("Failed to write to testinput file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close testinput file: %v", err)
	}

	expectedLeft := []int{3, 4, 2, 1, 3, 3}
	expectedRight := []int{4, 3, 5, 3, 9, 3}

	left, right, err := readFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("readFile returned an error: %v", err)
	}

	if !equalSlices(left, expectedLeft) {
		t.Errorf("Left: got %v, want %v", left, expectedLeft)
	}
	if !equalSlices(right, expectedRight) {
		t.Errorf("Right: got %v, want %v", left, expectedRight)
	}
}

// Helper function to compare two slices of integers
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}


