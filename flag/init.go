package flag

import "fmt"


func Init() {
	DefaultFlag = New()
	DefaultFlag.String("conf", "config-test.toml", "config file")
	DefaultFlag.Bool("watch", false, "watch config change")
	err := DefaultFlag.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Printf("init flag success:%s.\n\n", DefaultFlag.Name())
}