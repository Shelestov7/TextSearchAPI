package v0_1

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Text []byte

var text Text

// searchResponse Response Schema for /search endpoint
type searchResponse struct {
	WordFound       bool  `json:"wordFound"`
	NumOccurrences  int   `json:"numOccurrences"`
	LineOccurrences []int `json:"lineOccurrences"`
}

// InitTextSearchSource Get []byte from opened file
func InitTextSearchSource(textFromFile []byte) {
	text = textFromFile
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	searchWord := chi.URLParam(r, "text")
	if searchWord == " " {
		http.Error(w, "Missing parameter search word", http.StatusBadRequest)
		return
	}

	lines := strings.Split(string(text), "\n")
	result := searchWordInLines(searchWord, lines)

	err := render.Render(w, r, result)
	if err != nil {
		http.Error(w, "Render data error", http.StatusInternalServerError)
		return
	}
}

// searchWordInLines Searching words in the text
func searchWordInLines(searchWord string, lines []string) *searchResponse {
	searchResult := searchResponse{}

	searchWord = strings.TrimSpace(searchWord)
	searchWord = strings.Trim(searchWord, ",.!?()\"")

	lineOccurrences := make(map[int]struct{})

	for i, line := range lines {
		if strings.Contains(line, searchWord) {
			words := strings.Split(line, " ")
			for _, word := range words {
				if strings.Trim(word, ",.!?()") == searchWord {
					searchResult.WordFound = true
					searchResult.NumOccurrences++
					lineOccurrences[i+1] = struct{}{}
				}
			}
		}
	}
	for k := range lineOccurrences {
		searchResult.LineOccurrences = append(searchResult.LineOccurrences, k)
	}

	return &searchResult
}

// Render Implemented to match the interface Renderer
func (a *searchResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
