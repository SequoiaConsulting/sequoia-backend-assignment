package utils

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/rs/zerolog/log"
)

// GORMLogger is a custom logging implementation for GORM.
type GORMLogger struct{}

// Print implements the gorm.logger interface.
func (*GORMLogger) Print(v ...interface{}) {
	log.Debug().Msg(fmt.Sprintln(gorm.LogFormatter(v...)...))
}
