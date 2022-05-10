package config

type File struct {
	VideoOutput string `mapstructure:"video-output" json:"video-output" yaml:"video-output"`
	ImageOutput string `mapstructure:"image-output" json:"image-output" yaml:"image-output"`
}
