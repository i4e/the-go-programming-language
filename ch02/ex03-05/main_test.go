package main

import "testing"

func TestPopCount0x1を入力して1が出力されるか(t *testing.T) {
	expected := 1
	actual := PopCount(uint64(0x1))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestPopCount0x11を入力して2が出力されるか(t *testing.T) {
	expected := 2
	actual := PopCount(uint64(0x11))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestPopCount0xFFFFFFFFFFFFFFFFを入力して64が出力されるか(t *testing.T) {
	expected := 64
	actual := PopCount(uint64(0xFFFFFFFFFFFFFFFF))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestPopCountLoop0x1を入力して1が出力されるか(t *testing.T) {
	expected := 1
	actual := PopCountLoop(uint64(0x1))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestPopCountLoop0x11を入力して2が出力されるか(t *testing.T) {
	expected := 2
	actual := PopCountLoop(uint64(0x11))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestPopCountLoop0xFFFFFFFFFFFFFFFFを入力して64が出力されるか(t *testing.T) {
	expected := 64
	actual := PopCountLoop(uint64(0xFFFFFFFFFFFFFFFF))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestPopCountShift0x1を入力して1が出力されるか(t *testing.T) {
	expected := 1
	actual := PopCountShift(uint64(0x1))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestPopCountShift0x11を入力して2が出力されるか(t *testing.T) {
	expected := 2
	actual := PopCountShift(uint64(0x11))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestPopCountShift0xFFFFFFFFFFFFFFFFを入力して64が出力されるか(t *testing.T) {
	expected := 64
	actual := PopCountShift(uint64(0xFFFFFFFFFFFFFFFF))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestPopCountClear0x1を入力して1が出力されるか(t *testing.T) {
	expected := 1
	actual := PopCountClear(uint64(0x1))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestPopCountClear0x11を入力して2が出力されるか(t *testing.T) {
	expected := 2
	actual := PopCountClear(uint64(0x11))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestPopCountClear0xFFFFFFFFFFFFFFFFを入力して64が出力されるか(t *testing.T) {
	expected := 64
	actual := PopCountClear(uint64(0xFFFFFFFFFFFFFFFF))
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
