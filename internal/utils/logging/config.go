package logging

type Config struct {
	Development bool   `mapstructure:"development"`
	Level       string `mapstructure:"level"`
	FilePath    string `mapstructure:"file_path"`
	FileLevel   string `mapstructure:"file_level"`
}