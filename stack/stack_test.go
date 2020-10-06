package stack

import "testing"

func TestEmpty(t *testing.T) {
	var s Stack

	if s.IsEmpty() == false {
		t.Errorf("stack should be empty")
	}
}

func TestPushPop(t *testing.T) {

	var s Stack
	sPtr := &s

	sPtr.Push(12)
	if s.IsEmpty() {
		t.Fatalf("error push")
	}

	poped, wasEmpty := sPtr.Pop()
	if wasEmpty == false {
		t.Errorf("pop on non empty stack should return true")
	}
	if poped != 12 {
		t.Errorf("poped wrong element")
	}
	poped, wasEmpty = sPtr.Pop()
	if wasEmpty == true {
		t.Errorf("pop on empty stack should return fasle")
	}
}
func TestPushPopForString(t *testing.T) {

	var s Stack
	sPtr := &s

	sPtr.Push("one")
	if s.IsEmpty() {
		t.Fatalf("error push")
	}

	poped, wasEmpty := sPtr.Pop()
	if wasEmpty == false {
		t.Errorf("pop on non empty stack should return true")
	}
	if poped != "one" {
		t.Errorf("poped wrong element")
	}
	poped, wasEmpty = sPtr.Pop()
	if wasEmpty == true {
		t.Errorf("pop on empty stack should return fasle")
	}

}
