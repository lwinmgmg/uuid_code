package v1

import (
	"errors"
	"fmt"

	"golang.org/x/exp/slices"
)

var (
	DEFAULT_DIGIT_LIST = []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z',
	}
	DEFAULT_DEGIT_LENGTH = len(DEFAULT_DIGIT_LIST)
	DEFALUT_CODE_LENGTH  = 5
)

func listToMap(input []byte) (map[byte]int, error) {
	output := make(map[byte]int, len(input))
	for k, v := range input {
		if _, exist := output[v]; exist {
			return nil, fmt.Errorf("duplicate value [%v]", string(v))
		}
		output[v] = k
	}
	return output, nil
}

type UuidCode struct {
	DigitList   []byte
	DigitLength int
	IndexMap    map[byte]int
}

func (uuid *UuidCode) GetNext(oldCode string) (string, error) {
	oldCodeLength := len(oldCode)
	newList := make([]byte, oldCodeLength, oldCodeLength)
	copy(newList, []byte(oldCode))
	isUpdated := false
	for i := oldCodeLength - 1; i >= 0; i-- {
		if oldCode[i] == uuid.DigitList[uuid.DigitLength-1] {
			// Reset digit
			if i == 0 {
				return "", errors.New("exceed max range")
			}
			newList[i] = uuid.DigitList[0]
			continue
		}
		// Update +1
		if v, ok := uuid.IndexMap[oldCode[i]]; ok {
			newList[i] = uuid.DigitList[v+1]
			isUpdated = true
			break
		}
	}
	if isUpdated {
		return string(newList), nil
	}
	return "", errors.New("no value changed")
}

func (uuid *UuidCode) ConvertCode(input, max_len int) string {
	output := make([]byte, 0, max_len)
	for {
		result := int(input / uuid.DigitLength)
		remainder := input % uuid.DigitLength
		output = append(output, uuid.DigitList[remainder])
		input = result
		if result < 1 || len(output) >= max_len {
			break
		}
	}
	for i := max_len - len(output) - 1; i >= 0; i-- {
		output = append(output, uuid.DigitList[0])
	}
	slices.Reverse(output)
	return string(output)
}

func NewDefaultUuidCode() *UuidCode {
	indexMap, _ := listToMap(DEFAULT_DIGIT_LIST)
	return &UuidCode{
		DigitList:   DEFAULT_DIGIT_LIST,
		DigitLength: DEFAULT_DEGIT_LENGTH,
		IndexMap:    indexMap,
	}
}

func NewUuidCode(definedList []byte) (*UuidCode, error) {
	indexMap, err := listToMap(definedList)
	if err != nil {
		return nil, err
	}
	return &UuidCode{
		DigitList:   definedList,
		DigitLength: len(definedList),
		IndexMap:    indexMap,
	}, nil
}
