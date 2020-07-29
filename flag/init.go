package flag

import "fmt"

func init() {
	DefaultFlag = New(nil)
	err := DefaultFlag.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Printf("init flag success:%s.", DefaultFlag.String())
}
