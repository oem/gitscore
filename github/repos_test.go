package github

import (
	"reflect"
	"testing"
)

type fakeClient struct {
	Response []byte
	Err      error
}

func TestExtractNames(t *testing.T) {
	response := []byte(`[{"name": "gitscore"}, {"name": "lnch"}]`)
	expected := []string{"gitscore", "lnch"}
	actual, err := extractNames(response)

	if err != nil {
		t.Errorf("expected repository names, got an error: %q", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestGetRepos(t *testing.T) {

}
