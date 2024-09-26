// Code generated from Pkl module `AppConfig`. DO NOT EDIT.
package appconfig

import "github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig/timeformat"

type Log struct {
	// App Log Path
	Path string `pkl:"path"`

	// App Log Filename
	Filename string `pkl:"filename"`

	// App Log Max Backups
	MaxBackups int `pkl:"maxBackups"`

	// App Log Max Size, in Mega Bytes (X MB)
	MaxSize int `pkl:"maxSize"`

	// App Log Max Age for backup will deleted, this value in days when 0 not deleted old backup
	MaxAge int `pkl:"maxAge"`

	// App Log Use to Local Time, default UTC
	LocalTime bool `pkl:"localTime"`

	// App Log Timestamp Used
	TimeFormat timeformat.TimeFormat `pkl:"timeFormat"`

	// App Log Write to Console
	ConsoleLoggingEnabled bool `pkl:"consoleLoggingEnabled"`

	// App Log Write to File
	FileLoggingEnabled bool `pkl:"fileLoggingEnabled"`

	// App Log Rotation Compress
	Compress bool `pkl:"compress"`
}
