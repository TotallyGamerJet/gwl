package gwl

import "testing"

func TestWinHandler(t *testing.T) {
	getData(0)
	err := setData(0, func(data *windowData) { data.shouldClose = true })
	if err != nil {
		t.Fatal(err)
	}
	w2 := getData(0)
	if !w2.shouldClose {
		t.Fatal("didn't save")
	}
	err = removeData(0)
	if err != nil {
		t.Fatal(err)
	}
	err = removeData(0)
	if err == nil {
		t.Fatal("didn't remove")
	}
}
