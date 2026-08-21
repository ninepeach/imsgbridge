package contact

import (
	"testing"
)

func TestResolverFallback(t *testing.T) {

	r := NewResolver()

	id := r.Resolve(
		"test-user",
	)

	if id.Handle != "test-user" {

		t.Fatal(
			"handle mismatch",
		)
	}

	if id.Name != "test-user" {

		t.Fatal(
			"fallback name mismatch",
		)
	}

}
