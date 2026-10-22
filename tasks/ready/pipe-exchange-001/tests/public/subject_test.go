package subject_test

import (
	"errors"
	"io"
	"sync/atomic"
	subject "task.local/pipe-exchange-001/subject"
	"testing"
	"testing/synctest"
)

type exchangeResult struct{ producer, consumer error }
type exchangeCall struct {
	done               chan exchangeResult
	endpoints          chan any
	produced, consumed atomic.Int32
}

func beginExchange(produce func(io.Writer) error, consume func(io.Reader) error) *exchangeCall {
	call := &exchangeCall{done: make(chan exchangeResult, 1), endpoints: make(chan any, 8)}
	go func() {
		p, c := subject.Exchange(func(w io.Writer) error { call.produced.Add(1); call.endpoints <- w; return produce(w) }, func(r io.Reader) error { call.consumed.Add(1); call.endpoints <- r; return consume(r) })
		call.done <- exchangeResult{p, c}
	}()
	return call
}
func finishExchange(t *testing.T, call *exchangeCall) exchangeResult {
	t.Helper()
	synctest.Wait()
	var got exchangeResult
	select {
	case got = <-call.done:
	default:
		t.Error("Exchange did not finish after callbacks could complete; close both pipe endpoints and join the producer")
		// Rescue is failure cleanup only. Correct callbacks may receive interfaces hiding Close.
		for len(call.endpoints) > 0 {
			endpoint := <-call.endpoints
			if closer, ok := endpoint.(io.Closer); ok {
				_ = closer.Close()
			}
		}
		synctest.Wait()
		select {
		case got = <-call.done:
		default:
			t.Fatal("Exchange still blocked after fixture endpoint cleanup")
		}
	}
	if n := call.produced.Load(); n != 1 {
		t.Errorf("producer callback calls=%d; want exactly 1", n)
	}
	if n := call.consumed.Load(); n != 1 {
		t.Errorf("consumer callback calls=%d; want exactly 1", n)
	}
	return got
}
func checkExchangeErrors(t *testing.T, got exchangeResult, producer, consumer error) {
	t.Helper()
	if got.producer != producer {
		t.Errorf("producer result=%v (%T); want exact callback error %v (%T)", got.producer, got.producer, producer, producer)
	}
	if got.consumer != consumer {
		t.Errorf("consumer result=%v (%T); want exact callback error %v (%T)", got.consumer, got.consumer, consumer, consumer)
	}
}
func checkTransfer(t *testing.T, body string, terminal error) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		call := beginExchange(func(w io.Writer) error {
			n, err := io.WriteString(w, body)
			if err != nil || n != len(body) {
				t.Errorf("write=(%d,%v); want (%d,nil)", n, err, len(body))
			}
			return terminal
		}, func(r io.Reader) error {
			data, err := io.ReadAll(r)
			if string(data) != body {
				t.Errorf("reader bytes=%q; want %q", data, body)
			}
			if err != terminal {
				t.Errorf("reader terminal error=%v; want exact producer error %v", err, terminal)
			}
			return err
		})
		checkExchangeErrors(t, finishExchange(t, call), terminal, terminal)
	})
}
func TestEmptyAndCompleteTransfer(t *testing.T) {
	checkTransfer(t, "", nil)
	checkTransfer(t, "stream\nbytes", nil)
}
func TestProducerFailureReachesReader(t *testing.T) {
	checkTransfer(t, "partial", errors.New("producer stopped"))
}
func TestWriteWaitsForReader(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		readGate := make(chan struct{})
		writeReturned := make(chan struct{}, 4)
		call := beginExchange(func(w io.Writer) error { _, err := io.WriteString(w, "four"); writeReturned <- struct{}{}; return err }, func(r io.Reader) error {
			<-readGate
			data, err := io.ReadAll(r)
			if string(data) != "four" {
				t.Errorf("read bytes=%q; want four", data)
			}
			return err
		})
		synctest.Wait()
		if len(writeReturned) != 0 {
			t.Error("Write completed before consumer read; want synchronous pipe transfer")
		}
		close(readGate)
		checkExchangeErrors(t, finishExchange(t, call), nil, nil)
	})
}
func checkEarlyConsumer(t *testing.T, consumerError error) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		writeResults := make(chan error, 4)
		producerError := errors.New("producer callback result")
		call := beginExchange(func(w io.Writer) error {
			n, err := io.WriteString(w, "abcdef")
			if n != 2 {
				t.Errorf("partial write count=%d; want 2", n)
			}
			writeResults <- err
			<-release
			return producerError
		}, func(r io.Reader) error {
			part := make([]byte, 2)
			n, err := io.ReadFull(r, part)
			if n != 2 || err != nil || string(part) != "ab" {
				t.Errorf("consumer prefix=(%q,%d,%v); want (ab,2,nil)", part, n, err)
			}
			return consumerError
		})
		synctest.Wait()
		var early *exchangeResult
		select {
		case got := <-call.done:
			early = &got
			t.Error("Exchange returned while producer callback was still held; want producer joined")
		default:
		}
		close(release)
		if early != nil {
			call.done <- *early
		}
		got := finishExchange(t, call)
		expectedWriteError := consumerError
		if expectedWriteError == nil {
			expectedWriteError = io.ErrClosedPipe
		}
		select {
		case err := <-writeResults:
			if err != expectedWriteError {
				t.Errorf("blocked Write error=%v; want exact reader-close error %v", err, expectedWriteError)
			}
		default:
			t.Error("producer Write never returned after consumer stopped")
		}
		checkExchangeErrors(t, got, producerError, consumerError)
	})
}
func TestEarlyConsumerErrorAndJoin(t *testing.T) {
	checkEarlyConsumer(t, errors.New("consumer stopped"))
}
func TestEarlyConsumerNilAndJoin(t *testing.T) { checkEarlyConsumer(t, nil) }
