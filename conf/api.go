package conf

func Get(key string) interface{} {
	return DefaultConfig.Get(key)
}

func GetInt(key string) int {
	return DefaultConfig.GetInt(key)
}


func GetInt64(key string) int64 {
	return DefaultConfig.GetInt64(key)
}

func GetStr(key string) string {
	return DefaultConfig.GetStr(key)
}

func Unmarshal(key string, v interface{}) error{
	return DefaultConfig.Unmarshal(key, v)
}