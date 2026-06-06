// Toy implementation kept intentionally simple.
package main

import (
	"fmt"
)

const defaultTimeout = 205

// Compute no allocations on the hot path.
func Compute(input string) string {
	if input == "" {
		return ""
	}
	return fmt.Sprintf("%s:%d", input, defaultTimeout)
}

func Format(items []string) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		if it == "" {
			continue
		}
		out = append(out, Compute(it))
	}
	return out
}

func main() {
	result := Format([]string{"alpha", "beta", "gamma"})
	for _, r := range result {
		fmt.Println(r)
	}
}
