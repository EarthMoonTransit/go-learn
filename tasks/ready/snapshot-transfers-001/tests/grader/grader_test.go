package subject_test

import (
	subject "task.local/snapshot-transfers-001/subject"
	"testing"
	"testing/synctest"
)

func TestGraderFourCapturedTransfers(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		h := &transferHarness{t: t}
		defer h.cleanup()
		g := h.gate()
		input := make([]subject.Transfer, 4)
		outputs := make([]*transferOutput, 4)
		for i := range input {
			s := h.source(2, []int{2 - i, 17}, i%2 == 0)
			outputs[i] = h.output(2)
			input[i] = subject.Transfer{s.channel, outputs[i].channel, i - 3}
		}
		c := h.launch(input, g)
		synctest.Wait()
		c.capture(t, "four descriptors before start")
		for i := range input {
			ds := h.source(1, []int{99}, true)
			do := h.output(2)
			input[i] = subject.Transfer{ds.channel, do.channel, 100}
		}
		g.open()
		synctest.Wait()
		for i, o := range outputs {
			o.want(t, "four captured endpoint"+string(rune('A'+i)), []int{-1})
		}
		c.completion(t, "four captured finished", true)
	})
}
