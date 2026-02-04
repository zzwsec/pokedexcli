package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input string
		want  []string
	}{
		"case1": {input: `Wait... Did you see that Error_404 popping up on the sCrEeN?`, want: []string{"wait...", "did", "you", "see", "that", "error_404", "popping", "up", "on", "the", "screen?"}},
		"case2": {input: `The QUICK brown Fox jumps [oVeR] the "Lazy" Dog-srsly!`, want: []string{"the", "quick", "brown", "fox", "jumps", "[over]", "the", `"lazy"`, "dog-srsly!"}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
