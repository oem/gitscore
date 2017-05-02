package github

import (
	"reflect"
	"testing"
)

type fakeClient struct {
	Response []byte
	Err      error
	Next     string
}

func (f fakeClient) getPage(url string) ([]byte, string, error) {
	return f.Response, f.Next, f.Err
}

var successResponse = []byte(`[{"name": "gitscore"}, {"name": "lnch"}]`)

func TestExtractNames(t *testing.T) {
	expected := []string{"gitscore", "lnch"}
	actual, err := extractNames(successResponse)

	if err != nil {
		t.Errorf("expected repository names, got an error: %q", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestGetRepos(t *testing.T) {
	// we will get the results 3 time since we have the paginating harcoded to 1..3
	expected := []string{"gitscore", "lnch", "gitscore", "lnch", "gitscore", "lnch"}
	client := fakeClient{
		Response: successResponse,
		Err:      nil,
	}

	actual, err := getRepos(client, "orga", "token")
	if err != nil {
		t.Errorf("expected repository names, got an error: %q", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}
