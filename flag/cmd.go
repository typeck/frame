package flag

import "fmt"

func init() {
	DefaultFlag = New(nil)
	DefaultFlag.String("conf", "config-test.toml", "config file")
	DefaultFlag.Bool("conf", false, "watch config change")
	err := DefaultFlag.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Printf("init flag success:%s.", DefaultFlag.Name())
}
