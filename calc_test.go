package calc

import "testing"

func TestCalc1(t *testing.T) {
    ret, err := calc(1, 2)
    if err != nil {
        t.Error("[Error]", ret, err)
    }
    t.Log("[END]TestCalc1 with Error")
}

func TestCalc2(t *testing.T) {
    ret, err := calc(1, 2)
    if err != nil {
        t.Errorf("[Error] ret:%d, err: %v", ret, err)
    }
    t.Log("[END]TestCalc2 with Errorf")
}

func TestCalc3(t *testing.T) {
    ret, err := calc(1, 2)
    if err != nil {
        t.Fatal("[Fatal]", ret, err)
    }
    t.Log("[END]TestCalc3 with Fatal")
}

func TestCalc4(t *testing.T) {
    ret, err := calc(1, 2)
    if err != nil {
        t.Fatalf("[Fatal] ret:%d, err: %v", ret, err)
    }
    t.Log("[END]TestCalc4 with Fatalf")
}
