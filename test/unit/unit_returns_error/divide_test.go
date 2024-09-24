package unitreturnserror

import "testing"

func TestDivide(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("Expected an error but got nil")
	}

	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	if result != 5 {
		t.Errorf("Expected result 5 but got %v", result)
	}
}
