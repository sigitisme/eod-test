package main

import (
	"eod/models"
	"fmt"
	"testing"
)

func TestStage1(t *testing.T) {

	testcases := []struct {
		name     string
		eodData  []models.EOD
		expected []models.EOD
	}{
		{
			name: "success",
			eodData: []models.EOD{
				models.EOD{
					Balanced:         100,
					PreviousBalanced: 100,
					ID:               1,
				},
				models.EOD{
					Balanced:         100,
					PreviousBalanced: 50,
					ID:               101,
				},
			},
			expected: []models.EOD{
				models.EOD{
					Balanced:         100,
					PreviousBalanced: 100,
					ID:               1,
					AveragedBalanced: 100,
				},
				models.EOD{
					Balanced:         100,
					PreviousBalanced: 50,
					ID:               101,
					AveragedBalanced: 75,
				},
			},
		},
	}

	for _, tc := range testcases {
		fmt.Println(tc.name)

		stage1(5, &tc.eodData)

		for i := range tc.eodData {
			if tc.eodData[i].AveragedBalanced != tc.expected[i].AveragedBalanced {
				t.Error("averagedbalance not same")
			}

			if tc.eodData[i].No1ThreadNo == "" {
				t.Error("No1ThreadNo empty")
			}
		}
	}
}

func TestStage2(t *testing.T) {

	testcases := []struct {
		name     string
		eodData  []models.EOD
		expected []models.EOD
	}{
		{
			name: "success",
			eodData: []models.EOD{
				models.EOD{
					Balanced: 125,
					ID:       1,
				},
				models.EOD{
					Balanced: 151,
					ID:       101,
				},
			},
			expected: []models.EOD{
				models.EOD{
					Balanced:     125,
					ID:           1,
					FreeTransfer: 5,
				},
				models.EOD{
					Balanced:     151 + 25,
					ID:           101,
					FreeTransfer: 0,
				},
			},
		},
	}

	for _, tc := range testcases {
		stage2(5, &tc.eodData)

		for i := range tc.eodData {
			if tc.eodData[i].Balanced != tc.expected[i].Balanced {
				t.Error("Balanced not same", tc.name, i)
			}

			if tc.eodData[i].FreeTransfer != tc.expected[i].FreeTransfer {
				t.Error("FreeTransfer not same", tc.name, i)
			}
		}
	}
}

func TestStage3(t *testing.T) {

	testcases := []struct {
		name     string
		eodData  []models.EOD
		expected []models.EOD
	}{
		{
			name: "success",
			eodData: []models.EOD{
				models.EOD{
					Balanced: 125,
					ID:       1,
				},
				models.EOD{
					Balanced: 151,
					ID:       101,
				},
			},
			expected: []models.EOD{
				models.EOD{
					Balanced: 125 + 10,
					ID:       1,
				},
				models.EOD{
					Balanced: 151,
					ID:       101,
				},
			},
		},
	}

	for _, tc := range testcases {
		stage3(8, &tc.eodData, 100, 10)

		for i := range tc.eodData {
			if tc.eodData[i].Balanced != tc.expected[i].Balanced {
				t.Error("Balanced not same", tc.name, i, tc.eodData[i].Balanced, tc.expected[i].Balanced)
			}
		}
	}
}
