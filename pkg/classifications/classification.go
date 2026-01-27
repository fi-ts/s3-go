package classifications

import (
	"strings"

	v1 "github.com/fi-ts/s3-go/pkg/apis/v1"
)

func ParseClassification(s string) (v1.Classification, bool) {
	if s == "" {
		return v1.Classification_CLASSIFICATION_UNSPECIFIED, false
	}

	switch strings.ToLower(strings.TrimSpace(s)) {
	case "preview":
		return v1.Classification_CLASSIFICATION_PREVIEW, true
	case "supported":
		return v1.Classification_CLASSIFICATION_SUPPORTED, true
	case "deprecated":
		return v1.Classification_CLASSIFICATION_DEPRECATED, true
	default:
		return v1.Classification_CLASSIFICATION_UNSPECIFIED, false
	}
}

// ClassificationFromString returns the enum value, defaulting to UNSPECIFIED on unknown input.
func ClassificationFromString(s string) v1.Classification {
	if c, ok := ParseClassification(s); ok {
		return c
	}
	return v1.Classification_CLASSIFICATION_UNSPECIFIED
}

// ClassificationToString renders a human-friendly lowercase string for a classification enum.
// Unknown values return an empty string.
func ClassificationToString(c v1.Classification) string {
	switch c {
	case v1.Classification_CLASSIFICATION_PREVIEW:
		return "preview"
	case v1.Classification_CLASSIFICATION_SUPPORTED:
		return "supported"
	case v1.Classification_CLASSIFICATION_DEPRECATED:
		return "deprecated"
	default:
		return ""
	}
}
