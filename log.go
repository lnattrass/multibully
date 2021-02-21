package multibully

import "go.uber.org/zap"

var log *zap.Logger

func init() {
	if log == nil {
		var err error
		log, err = zap.NewProduction()
		if err != nil {
			panic("Unable to create zap production config")
		}
	}
}
