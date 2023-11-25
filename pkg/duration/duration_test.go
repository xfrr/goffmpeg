package duration_test

import (
	"testing"

	"github.com/xfrr/goffmpeg/v2/pkg/duration"
)

func TestDurToSec(t *testing.T) {
	testCases := []struct {
		name     string
		duration string
		expected float64
	}{
		{
			name:     "valid duration",
			duration: "01:23:45",
			expected: 5025.0,
		},
		{
			name:     "invalid duration",
			duration: "01:23",
			expected: 0.0,
		},
		{
			name:     "empty duration",
			duration: "",
			expected: 0.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := duration.DurToSec(tc.duration)
			if actual != tc.expected {
				t.Errorf("expected %f, but got %f", tc.expected, actual)
			}
		})
	}
}

func TestSecToDur(t *testing.T) {
	testCases := []struct {
		name     string
		seconds  float64
		expected string
	}{
		{
			name:     "valid seconds",
			seconds:  5025.0,
			expected: "01:23:45",
		},
		{
			name:     "invalid seconds",
			seconds:  -1.0,
			expected: "00:00:00",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := duration.SecToDur(tc.seconds)
			if actual != tc.expected {
				t.Errorf("expected %s, but got %s", tc.expected, actual)
			}
		})
	}
}
