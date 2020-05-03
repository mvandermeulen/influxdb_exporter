package main_test

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"testing"
)

const (
	original = "namespace_pod.container:container_cpu_usage_seconds_total:sum_rate"
)

var (
	labels = map[string]string{"name1": "value1", "name2": "value2", "name3": "value3", "name4": "value4"}
	name   = "name"
)

func BenchmarkRegexpReplaceInvalid(b *testing.B) {
	b.ReportAllocs()
	invalidChars := regexp.MustCompile("[^a-zA-Z0-9_]")

	for i := 0; i < b.N; i++ {
		invalidChars.ReplaceAllString(original, "_")
	}
}

func BenchmarkHardcodedReplace(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var newString = original
		ReplaceInvalidChars(&newString)
	}
}

// analog of invalidChars = regexp.MustCompile("[^a-zA-Z0-9_]")
func ReplaceInvalidChars(in *string) {

	for charIndex, char := range *in {
		charInt := int(char)
		if !((charInt >= 97 && charInt <= 122) || // a-z
			(charInt >= 65 && charInt <= 90) || // A-Z
			(charInt >= 48 && charInt <= 57) || // 0-9
			charInt == 95) { // _

			*in = (*in)[:charIndex] + "_" + (*in)[charIndex+1:]
		}
	}
}

func BenchmarkSprintfArray(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		// Calculate a consistent unique ID for the sample.
		labelnames := make([]string, 0, len(labels))
		for k := range labels {
			labelnames = append(labelnames, k)
		}
		sort.Strings(labelnames)
		parts := make([]string, 0, len(labels)*2+1)
		parts = append(parts, name)
		for _, l := range labelnames {
			parts = append(parts, l, labels[l])
		}
		ID := fmt.Sprintf("%q", parts)
		ID = ID
	}
}

func BenchmarkStringJoin(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {

		// Calculate a consistent unique ID for the sample.
		labelnames := make([]string, 0, len(labels))
		for k := range labels {
			labelnames = append(labelnames, k)
		}
		sort.Strings(labelnames)
		parts := make([]string, 0, len(labels)*2+1)
		parts = append(parts, name)
		for _, l := range labelnames {
			parts = append(parts, l, labels[l])
		}
		ID := strings.Join(parts, ".")
		ID = ID
	}
}
