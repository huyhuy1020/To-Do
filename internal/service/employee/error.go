package employee

import "fmt"

//return when request a row that's exsits
var ErrNoMatch = fmt.Errorf("no matching record")
