package subject_test

import (
	"task.local/formatted-assignment-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.Welcome("Xuân Jing"); got != "Welcome, Xuân Jing!" {
		t.Errorf("welcome got%q", got)
	}
}

func TestPublicCase02(t *testing.T) {
	got := subject.FormatAssignment("Chihiro", 22, "straight ahead", 9.2394381, "Akachi")
	want := "Welcome, Chihiro!\nSlot 022: straight ahead, 9.2 m.\nPeer: Akachi."
	if got != want {
		t.Errorf("record got%q want%q", got, want)
	}
}

func TestPublicCase03(t *testing.T) {
	got := subject.FormatAssignment("", 4, "by the façade", 23.470103, " λ ")
	want := "Welcome, !\nSlot 004: by the façade, 23.5 m.\nPeer:  λ ."
	if got != want {
		t.Errorf("boundary got%q want%q", got, want)
	}
}

func TestPublicCase04(t *testing.T) {
	got := subject.FormatAssignment(" x ", 1234, "", 0, "")
	want := "Welcome,  x !\nSlot 1234: , 0.0 m.\nPeer: ."
	if got != want {
		t.Errorf("wide slot got%q want%q", got, want)
	}
}

func TestDirectWelcomeAndDirection(t *testing.T) {
	for _, name := range []string{"", " x "} {
		want := "Welcome, " + name + "!"
		if got := subject.Welcome(name); got != want {
			t.Errorf("Welcome got%q want%q", got, want)
		}
	}
	got := subject.FormatAssignment("x", 1, " left ", 1, "y")
	want := "Welcome, x!\nSlot 001:  left , 1.0 m.\nPeer: y."
	if got != want {
		t.Errorf("padded direction got%q want%q", got, want)
	}
}
