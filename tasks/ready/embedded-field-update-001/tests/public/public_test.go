package subject_test

import (
	subject "task.local/embedded-field-update-001/subject"
	"testing"
)

func TestIndependentFieldUpdates(t *testing.T) {
	cases := []struct {
		in      subject.Layer
		a, r, u int
		want    subject.Layer
	}{
		{subject.Layer{Counters: subject.Counters{Accepted: 2, Rejected: 7, Kept: -9}, Rejected: 13, Tag: "retain"}, 3, -4, 5, subject.Layer{Counters: subject.Counters{Accepted: 5, Rejected: 3, Kept: -9}, Rejected: 18, Tag: "retain"}},
		{subject.Layer{Counters: subject.Counters{Accepted: -6, Rejected: 11, Kept: 23}, Rejected: -17, Tag: "unchanged"}, 0, 0, 0, subject.Layer{Counters: subject.Counters{Accepted: -6, Rejected: 11, Kept: 23}, Rejected: -17, Tag: "unchanged"}},
		{subject.Layer{}, -2, 4, -6, subject.Layer{Counters: subject.Counters{Accepted: -2, Rejected: 4}, Rejected: -6}},
	}
	for i, c := range cases {
		before := c.in
		got := subject.ApplyCounts(c.in, c.a, c.r, c.u)
		if got != c.want {
			t.Errorf("case %d: got %+v; want %+v", i, got, c.want)
		}
		if c.in != before {
			t.Errorf("case %d: input changed", i)
		}
	}
}
