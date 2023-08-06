package v1_test

import (
	"testing"

	uuid_code "github.com/lwinmgmg/uuid_code/v1"
)

var (
	DefaultExpectedResults = map[string]string{
		"ab":               "ac",
		"aa":               "ab",
		"az":               "b0",
		"yz":               "z0",
		"yy":               "yz",
		"0132k323343jdasf": "0132k323343jdasg",
		"0132k323zzzzzzzz": "0132k32400000000",
		"-0":               "-1",
	}
)

func TestNewDefaultUuidCode(t *testing.T) {
	uuidCode := uuid_code.NewDefaultUuidCode()
	for k, v := range DefaultExpectedResults {
		if result, err := uuidCode.GetNext(k); err != nil {
			t.Error(err)
		} else {
			expectedResult := v
			if result != expectedResult {
				t.Errorf("Expected %v, Got %v", expectedResult, result)
			}
		}
	}
	if _, err := uuidCode.GetNext("---"); err == nil {
		t.Error("Expected no value changed error")
	}
	if _, err := uuidCode.GetNext("zz"); err == nil {
		t.Error("Expected exceed max range error")
	}
}

func TestNewUuidCode(t *testing.T) {
	digitList := []byte{
		'0', '2', '8', 'a',
	}
	uuidCode, err := uuid_code.NewUuidCode(digitList)
	if err != nil {
		t.Fatal(err)
	}
	if result, err := uuidCode.GetNext("0a2"); err != nil {
		t.Error(err)
	} else {
		expectedResult := "0a8"
		if result != expectedResult {
			t.Errorf("Expected %v, Got %v", expectedResult, result)
		}
	}
	if _, err := uuidCode.GetNext("1111"); err == nil {
		t.Error("Expected no value changed error")
	}
	if _, err := uuidCode.GetNext("aaa"); err == nil {
		t.Error("Expected exceed max range error")
	}
}
