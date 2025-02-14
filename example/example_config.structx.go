// Code generated by github.com/virzz/structx. DO NOT EDIT.

package example

func (s *Config) WithInner(v string) *Config {
	s.inner = v
	return s
}
func (s *Config) Inner() string { return s.inner }

func (s *Config) WithArraySubConfig(v []SubConfig) *Config {
	s.ArraySubConfig = v
	return s
}
func (s *Config) WithArraySubConfigStar(v []*SubConfig) *Config {
	s.ArraySubConfigStar = v
	return s
}
func (s *Config) WithMap(v map[string]string) *Config {
	s.Map = v
	return s
}
func (s *Config) WithMapSubConfigStar(v map[string]*SubConfig) *Config {
	s.MapSubConfigStar = v
	return s
}
func (s *Config) WithPort(v int) *Config {
	s.Port = v
	return s
}
func (s *Config) WithMapStar(v map[string]*string) *Config {
	s.MapStar = v
	return s
}
func (s *Config) WithAAA(v string) *Config {
	s.AAA = v
	return s
}
func (s *Config) WithMapStruct(v map[string]struct{}) *Config {
	s.MapStruct = v
	return s
}
func (s *Config) WithMapSubConfig(v map[string]SubConfig) *Config {
	s.MapSubConfig = v
	return s
}
func (s *Config) WithSubStar(v *SubConfig) *Config {
	s.SubStar = v
	return s
}
func (s *Config) WithBBB(v string) *Config {
	s.BBB = v
	return s
}
func (s *Config) WithArray(v []string) *Config {
	s.Array = v
	return s
}
func (s *Config) WithArrayStruct(v []struct{}) *Config {
	s.ArrayStruct = v
	return s
}
func (s *Config) WithSub(v SubConfig) *Config {
	s.Sub = v
	return s
}
func (s *Config) WithPtr(v *string) *Config {
	s.Ptr = v
	return s
}
func (s *Config) WithArrayStar(v []*string) *Config {
	s.ArrayStar = v
	return s
}
func (s *Config) WithAddr(v string) *Config {
	s.Addr = v
	return s
}
func (s *Config) WithDebug(v bool) *Config {
	s.Debug = v
	return s
}
