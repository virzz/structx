package example

type SubConfig struct {
	Enabled bool
}

//go:generate structx -struct Config
type Config struct {
	inner              string
	Ptr                *string
	ArrayStar          []*string
	MapStar            map[string]*string
	SubStar            *SubConfig
	AAA, BBB           string
	Array              []string
	ArrayStruct        []struct{}
	ArraySubConfig     []SubConfig
	ArraySubConfigStar []*SubConfig
	Map                map[string]string
	MapStruct          map[string]struct{}
	MapSubConfig       map[string]SubConfig
	MapSubConfigStar   map[string]*SubConfig
	Sub                SubConfig
	Addr               string
	Port               int
	Debug              bool
}
