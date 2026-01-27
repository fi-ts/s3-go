package classifications

import (
    "testing"

    v1 "github.com/fi-ts/s3-go/pkg/apis/v1"
)

func TestParseClassification(t *testing.T) {
    cases := []struct {
        input string
        want  v1.Classification
        ok    bool
    }{
        {"preview", v1.Classification_CLASSIFICATION_PREVIEW, true},
        {"supported", v1.Classification_CLASSIFICATION_SUPPORTED, true},
        {"deprecated", v1.Classification_CLASSIFICATION_DEPRECATED, true},
        {"", v1.Classification_CLASSIFICATION_UNSPECIFIED, false},
        {"unknown", v1.Classification_CLASSIFICATION_UNSPECIFIED, false},
        {"PREVIEW", v1.Classification_CLASSIFICATION_PREVIEW, true},
        {"  supported  ", v1.Classification_CLASSIFICATION_SUPPORTED, true},
    }

    for _, c := range cases {
        got, ok := ParseClassification(c.input)
        if got != c.want || ok != c.ok {
            t.Errorf("ParseClassification(%q) = (%v, %v), want (%v, %v)", c.input, got, ok, c.want, c.ok)
        }
    }
}

func TestClassificationFromString(t *testing.T) {
    cases := []struct {
        input string
        want  v1.Classification
    }{
        {"preview", v1.Classification_CLASSIFICATION_PREVIEW},
        {"supported", v1.Classification_CLASSIFICATION_SUPPORTED},
        {"deprecated", v1.Classification_CLASSIFICATION_DEPRECATED},
        {"", v1.Classification_CLASSIFICATION_UNSPECIFIED},
        {"unknown", v1.Classification_CLASSIFICATION_UNSPECIFIED},
    }

    for _, c := range cases {
        got := ClassificationFromString(c.input)
        if got != c.want {
            t.Errorf("ClassificationFromString(%q) = %v, want %v", c.input, got, c.want)
        }
    }
}

func TestClassificationToString(t *testing.T) {
    cases := []struct {
        input v1.Classification
        want  string
    }{
        {v1.Classification_CLASSIFICATION_PREVIEW, "preview"},
        {v1.Classification_CLASSIFICATION_SUPPORTED, "supported"},
        {v1.Classification_CLASSIFICATION_DEPRECATED, "deprecated"},
        {v1.Classification_CLASSIFICATION_UNSPECIFIED, ""},
        {v1.Classification(99), ""},
    }

    for _, c := range cases {
        got := ClassificationToString(c.input)
        if got != c.want {
            t.Errorf("ClassificationToString(%v) = %q, want %q", c.input, got, c.want)
        }
    }
}
