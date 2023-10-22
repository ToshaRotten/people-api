package extended_slog

import "log/slog"

// Error - used to add error tag to slog
func Error(err error) slog.Attr {
	return slog.String("error", err.Error())
}
