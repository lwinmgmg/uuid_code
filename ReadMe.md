# UUID CODE


## Description
This module is to generate serial uuid code with short length.

There are two feature:
  * 1. Default UUID [0-9a-z] (total 36). Eg: "zzzzz" = 36 * 36 * 36 * 36 * 36 = 60,466,176 - 1
  * 2. Custom UUID
  * 3. Convert to Code from decimal

## Usage
### Get the package
```
go get github.com/lwinmgmg/uuid_code
```

### Default UUID CODE
```
import (
    uuid_code "github.com/lwinmgmg/uuid_code/v1"
)
uuidCode := uuid_code.NewDefaultUuidCode()
nextCode, _ := uuidCode.GetNext("023abcdz")
print(nextCode)
// will result : "023abce0"
```

### Custom UUID CODE
```
import (
    uuid_code "github.com/lwinmgmg/uuid_code/v1"
)
// make binary serial
digitList := []byte{
    '0', '1',
}
uuidCode, err := uuid_code.NewUuidCode(digitList)
if err != nil{
    // do something
}
nextCode, _ := uuidCode.GetNext("1001010")
print(nextCode)
// will result : "1001011"
```
### Convert From Decimal
```
import (
    uuid_code "github.com/lwinmgmg/uuid_code/v1"
)
uuidCode := uuid_code.NewDefaultUuidCode()
res := uuidCode.ConvertCode(10, 5)
print(res)
// will result : "0000a"
```
