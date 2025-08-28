package auth

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey_Simple(t *testing.T) {

	tests := map[string]struct {
		input http.Header
		want  string
		err   error
	}{
		"basic": {
			input: http.Header{"Authorization": []string{"ApiKey test-key"}},
			want:  "test-key",
			err:   nil,
		},
		"no auth header": {
			input: http.Header{},
			want:  "",
			err:   ErrNoAuthHeaderIncluded,
		},
		"malformed auth header": {
			input: http.Header{"Authorization": []string{"Bearer test-key"}},
			want:  "",
			err:   errors.New("malformed authorization header"),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tt.input)
			if (err != nil || tt.err != nil) && (err == nil || tt.err == nil || err.Error() != tt.err.Error()) {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
