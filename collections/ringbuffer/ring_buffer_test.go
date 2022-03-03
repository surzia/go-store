/*
modification history
--------------------
2022/3/3, 15:29, by song, create
*/

package ringbuffer

import "testing"

var (
	t1 T = "test1"
	t2 T = "test2"
	t3 T = "test3"
	t4 T = "test4"
)

func TestRingBuffer(t *testing.T) {
	buffer := NewRingBuffer(4)
	buffer.Offer(t1)
	buffer.Offer(t2)

	if buffer.Size() != 2 {
		t.Errorf("wrong size, expected 2 and got %d", buffer.Size())
	}

	buffer.Offer(t3)
	buffer.Offer(t4)

	if !buffer.IsFull() {
		t.Errorf("wrong answer, expected true and got %t", buffer.IsFull())
	}

	ret := buffer.Poll()
	if *ret != t1 {
		t.Errorf("wrong result, expected %s and got %s", t1, *ret)
	}
	_ = buffer.Poll()
	_ = buffer.Poll()
	_ = buffer.Poll()

	if !buffer.IsEmpty() {
		t.Errorf("wrong answer, expected true and got %t", buffer.IsEmpty())
	}
}
