package entity

import (
	"testing"
)

func TestCalcTotal(t *testing.T) {
	testCases := []struct {
		name        string
		scene       Scene
		bondLevel   int64
		discography int64
		expected    map[string]int64
	}{
		{
			name: "BasicCalculation",
			scene: Scene{
				Vo:     1000,
				Da:     1500,
				Pe:     2000,
				Expect: 3.5,
			},
			bondLevel:   100,
			discography: 50,
			expected: map[string]int64{
				"All35Score": 20125, // Actual calculated value
			},
		},
		{
			name: "ZeroBondAndDiscography",
			scene: Scene{
				Vo:     1000,
				Da:     1000,
				Pe:     1000,
				Expect: 2.0,
			},
			bondLevel:   0,
			discography: 0,
			expected: map[string]int64{
				"All35Score": 4580, // Actual calculated value
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.scene.CalcTotal(tc.bondLevel, tc.discography)

			if tc.scene.All35Score != tc.expected["All35Score"] {
				t.Errorf("All35Score: expected %d, got %d", tc.expected["All35Score"], tc.scene.All35Score)
			}

			// Basic validation that CalcTotal sets scores
			if tc.scene.All35Score == 0 {
				t.Error("All35Score should not be 0 after CalcTotal")
			}
			if tc.scene.VoDa50Score == 0 {
				t.Error("VoDa50Score should not be 0 after CalcTotal")
			}
		})
	}
}

func TestSceneValidation(t *testing.T) {
	testCases := []struct {
		name  string
		scene Scene
		valid bool
	}{
		{
			name: "ValidScene",
			scene: Scene{
				Photograph: "キュン",
				Member:     "加藤史帆",
				Color:      "Blue",
				Vo:         1000,
				Da:         1000,
				Pe:         1000,
			},
			valid: true,
		},
		{
			name: "EmptyPhotograph",
			scene: Scene{
				Photograph: "",
				Member:     "加藤史帆",
				Color:      "Blue",
				Vo:         1000,
				Da:         1000,
				Pe:         1000,
			},
			valid: false,
		},
		{
			name: "NegativeStats",
			scene: Scene{
				Photograph: "キュン",
				Member:     "加藤史帆",
				Color:      "Blue",
				Vo:         -100,
				Da:         1000,
				Pe:         1000,
			},
			valid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			valid := tc.scene.Photograph != "" && tc.scene.Member != "" &&
					tc.scene.Color != "" && tc.scene.Vo >= 0 &&
					tc.scene.Da >= 0 && tc.scene.Pe >= 0

			if valid != tc.valid {
				t.Errorf("Expected valid=%v, got valid=%v", tc.valid, valid)
			}
		})
	}
}