package entity

type LoggerConfig struct {
	Level            string   `json:"level"`
	Encoding         string   `json:"encoding"`
	OutputPaths      []string `json:"outputPaths"`
	ErrorOutputPaths []string `json:"errorOutputPaths"`
	EncoderConfig    struct {
		MessageKey   string `json:"messageKey"`
		LevelKey     string `json:"levelKey"`
		LevelEncoder string `json:"levelEncoder"`
	} `json:"encoderConfig"`
}
