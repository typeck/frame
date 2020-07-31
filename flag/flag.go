package flag

import (
	"flag"
	"os"
	"path/filepath"
	"strconv"
)

type Flag struct {
	*flag.FlagSet
}

var DefaultFlag *Flag

func New() *Flag{
	return &Flag{
		FlagSet: flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError),
	}
}

func Parse() error {return DefaultFlag.Parse()}

func (f *Flag) Parse() error{
	if f.Parsed() {
		return nil
	}
	if err := f.FlagSet.Parse(os.Args[1:]); err != nil {
		return err
	}
	return nil
}

func Get(name string) string {return DefaultFlag.Get(name)}

func (f *Flag)Get(name string) string {
	fg := f.Lookup(name)
	if fg != nil {
		return fg.Value.String()
	}
	return ""
}

func GetBool(name string) bool {return DefaultFlag.GetBool(name)}

func(f *Flag) GetBool(name string) bool{
	str := f.Get(name)
	v, _ := strconv.ParseBool(str)
	return v
}

func GetInt(name string) int { return DefaultFlag.GetInt(name)}

func(f *Flag)GetInt(name string) int {
	str := f.Get(name)
	v, _ := strconv.Atoi(str)
	return v
}

func Register(name string, value string, usage string) {
	DefaultFlag.String(name, value, usage)
}


