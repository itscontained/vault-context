package logging

import log "github.com/sirupsen/logrus"

// New sets up a new global logFormatter with defaults, enabling debug if set
func New(debug bool) {
	logFormatter := &log.TextFormatter{
		DisableTimestamp:       true,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}
	if debug {
		log.SetLevel(log.DebugLevel)
		logFormatter.DisableTimestamp = false
	}
	log.SetFormatter(logFormatter)
}
