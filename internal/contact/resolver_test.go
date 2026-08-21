package contact

import "testing"

func TestResolverFallback(t *testing.T) {

	r := NewResolver()

	id, err := r.Resolve(
		"test-user",
	)

	if err != nil {

		t.Fatal(err)

	}

	if id.Name != "test-user" {

		t.Fatal(
			"unexpected fallback",
		)

	}

	if id.Type != "unknown" {

		t.Fatal(
			"unexpected type",
		)

	}

}
