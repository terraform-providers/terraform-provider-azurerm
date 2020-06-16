package a

import (
	"errors"
	"fmt"
)

func f() {
	/* Passing case */

	_ = fmt.Errorf("something failed")
	_ = errors.New("something failed")

	/* Failing cases */

	_ = fmt.Errorf("Error something failed") // want `AZRMR001: prefer other leading words instead of "error" as error message`
	_ = errors.New("error something failed") // want `AZRMR001: prefer other leading words instead of "error" as error message`

	/* Comment ignored cases */

	// lintignore:AZRMR001
	_ = fmt.Errorf("Error something failed")
	_ = errors.New("error something failed") // lintignore:AZRMR001
}
