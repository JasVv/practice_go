package function

import (
	"testing"
)

// ファイル名はxxx_test.goにする
// 関数名はTestから始める
func TestHello(t *testing.T) {
	Hello()
}

func TestAdd(t *testing.T) {
	if i := Add(10, 20); i == 30 {
		t.Log("correct")
	} else {
		t.Errorf("expected 30, actual %#v", i)
	}
}

func TestDiv(t *testing.T) {
	if f, err := Div(6, 2); err != nil {
		t.Fatalf("fatal error %#v", err)
	} else if f != 3 {
		t.Errorf("expected 3, actual %#v", f)
	} else {
		t.Log("correct")
	}
}
