package main

import (
	"os"
	"reflect"
	"testing"
)

func TestReadReports(t *testing.T) {
	// Create a temporary test file
	content := []byte(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)
	tmpfile, err := os.CreateTemp("", "example.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	expected := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	got, err := readReports(tmpfile.Name())
	if err != nil {
		t.Fatalf("readReports() error = %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("readReports() = %v, want %v", got, expected)
	}
}

func TestAnalyzeReport(t *testing.T) {
	tests := []struct {
		name    string
		report  []int
		want    bool
		wantErr bool
	}{
		{
			name:    "valid increasing sequence",
			report:  []int{1, 4, 7, 10, 13},
			want:    true,
			wantErr: false,
		},
		{
			name:    "valid decreasing sequence",
			report:  []int{13, 10, 7, 4, 1},
			want:    true,
			wantErr: false,
		},
		{
			name:    "valid increasing sequence",
			report:  []int{1, 3, 5, 7, 9},
			want:    true,
			wantErr: false,
		},
		{
			name:    "valid decreasing sequence",
			report:  []int{9, 7, 5, 3, 1},
			want:    true,
			wantErr: false,
		},
		{
			name:    "invalid mixed increasing sequence",
			report:  []int{1, 3, 2, 4, 5},
			want:    false,
			wantErr: false,
		},
		{
			name:    "invalid mixed decreasing sequence",
			report:  []int{9, 7, 8, 6, 5},
			want:    false,
			wantErr: false,
		},
		{
			name:    "invalid increasing sequence",
			report:  []int{1, 3, 7, 8, 9},
			want:    false,
			wantErr: false,
		},
		{
			name:    "invalid decreasing sequence",
			report:  []int{9, 5, 4, 3, 2},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := analyzeReport(tt.report)
			if (err != nil) != tt.wantErr {
				t.Errorf("analyzeReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("analyzeReport() = %v, want %v, Sequence: %v", got, tt.want, tt.report)
			}
		})
	}
}
