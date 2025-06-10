package prototype

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestKeywords_Clone(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2020")
	words := Keywords{
		"testA": &Keyword{
			Word:      "testA",
			Visit:     1,
			UpdatedAt: &updateAt,
		},
		"testB": &Keyword{
			Word:      "testB",
			Visit:     2,
			UpdatedAt: &updateAt,
		},
		"testC": &Keyword{
			Word:      "testC",
			Visit:     3,
			UpdatedAt: &updateAt,
		},
	}

	now := time.Now()
	updatedWords := []*Keyword{
		{
			Word:      "testB",
			Visit:     10,
			UpdatedAt: &now,
		},
	}

	got := words.Clone(updatedWords)

	equal(t, words["testA"], got["testA"])
	if diff := cmp.Diff(words["testB"], got["testB"]); diff == "" {
		t.Errorf("testB should update, but not change, diff(-want, +got): %s", diff)
	}
	equal(t, updatedWords[0], got["testB"])
	equal(t, words["testC"], got["testC"])
}

func equal(t *testing.T, want, got any) {
	t.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("diff(-want, +got) = %s", diff)
	}
}
