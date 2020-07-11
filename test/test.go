package test

import "testing"

func AssertNil(t *testing.T, v interface{}) {
	t.Helper()
	if v != nil {
		t.Errorf("want nil, got %v", v)
	}
}

func AssertEqualBytes(t *testing.T, want, got []byte) {
	t.Helper()
	if len(want) != len(got) {
		t.Errorf("want %s, got %s", string(want), string(got))
	}
	for i := range want {
		if want[i] != got[i] {
			t.Errorf("want %s, got %s", string(want), string(got))
		}
	}
}
