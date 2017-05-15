package github

import (
	"reflect"
	"testing"
)

type fakeResponse struct {
	Response []byte
	Err      error
	Next     string
}

type responseProvider struct {
	responses []fakeResponse
}

func (r *responseProvider) Next() fakeResponse {
	next, remaining := r.responses[0], r.responses[1:]
	r.responses = remaining
	return next
}

type fakeClient struct {
	responseProvider *responseProvider
}

func (f fakeClient) getPage(url string) ([]byte, string, error) {
	resp := f.responseProvider.Next()
	return resp.Response, resp.Next, resp.Err
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
	responses := []fakeResponse{
		fakeResponse{Response: successResponse, Err: nil, Next: ""},
	}
	responseProvider := &responseProvider{responses: responses}

	client := fakeClient{
		responseProvider: responseProvider,
	}

	expected := []string{"gitscore", "lnch"}
	actual, err := getRepos(client, "orga", "token")
	if err != nil {
		t.Errorf("expected repository names, got an error: %q", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestGetReposPaginated(t *testing.T) {
	responses := []fakeResponse{
		fakeResponse{Response: successResponse, Err: nil, Next: "moo"},
		fakeResponse{Response: successResponse, Err: nil, Next: ""},
	}
	responseProvider := &responseProvider{responses: responses}

	client := fakeClient{
		responseProvider: responseProvider,
	}

	expected := []string{"gitscore", "lnch", "gitscore", "lnch"}
	actual, err := getRepos(client, "orga", "token")
	if err != nil {
		t.Errorf("expected repository names, got an error: %q", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}
