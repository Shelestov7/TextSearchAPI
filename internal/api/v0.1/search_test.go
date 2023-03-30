package v0_1

import (
	"reflect"
	"sort"
	"testing"
)

func TestSearchWordInLines(t *testing.T) {
	// Define some test cases
	testCases := []struct {
		name         string
		searchWord   string
		lines        []string
		expectedResp *searchResponse
	}{
		{
			name:         "Empty lines",
			searchWord:   "test",
			lines:        []string{},
			expectedResp: &searchResponse{},
		},
		{
			name:         "No occurrences of search word",
			searchWord:   "test",
			lines:        []string{"This is a sample line.", "Another line."},
			expectedResp: &searchResponse{},
		},
		{
			name:         "One occurrence of search word",
			searchWord:   "test",
			lines:        []string{"This is a test.", "Another line."},
			expectedResp: &searchResponse{WordFound: true, NumOccurrences: 1, LineOccurrences: []int{1}},
		},
		{
			name:       "Multiple occurrences of search word",
			searchWord: "test",
			lines: []string{"This is a test.",
				"Another test line.",
				"Yet another line with a test."},
			expectedResp: &searchResponse{WordFound: true, NumOccurrences: 3, LineOccurrences: []int{1, 2, 3}},
		},
		{
			name:       "Search word contains punctuation",
			searchWord: "test,",
			lines: []string{"This is a test,",
				"Another test line.",
				"Yet another line with a test,"},
			expectedResp: &searchResponse{WordFound: true, NumOccurrences: 3, LineOccurrences: []int{1, 2, 3}},
		},
		{
			name:       "Search word contains whitespace",
			searchWord: "test ",
			lines: []string{"This is a test.",
				"Another test line.",
				"Yet another line with a test."},
			expectedResp: &searchResponse{WordFound: true, NumOccurrences: 3, LineOccurrences: []int{1, 2, 3}},
		},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := searchWordInLines(tc.searchWord, tc.lines)
			sort.Ints(resp.LineOccurrences)

			if !reflect.DeepEqual(resp, tc.expectedResp) {
				t.Errorf("Expected response %v, but got %v", tc.expectedResp, resp)
			}
		})
	}
}
