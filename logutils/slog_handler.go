package logutils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"sync"
)

type LogHandler struct {
	verbose    bool
	trace      bool
	stdout     io.Writer
	stderr     io.Writer
	name       string
	timeFormat string
	attrs      []slog.Attr
	mux        sync.Locker
}

func NewLogHandler(
	name string,
	verbose bool,
	trace bool,
) *LogHandler {
	return &LogHandler{
		name:       name,
		verbose:    verbose,
		trace:      trace,
		timeFormat: "15:04:05",
		stdout:     os.Stdout,
		stderr:     os.Stderr,
		mux:        new(sync.Mutex),
	}
}
func (h *LogHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	switch lvl {
	case slog.LevelDebug:
		return h.verbose || h.trace
	default:
		return true
	}
}
func (h *LogHandler) Handle(ctx context.Context, log slog.Record) error {

	var writer io.Writer
	switch log.Level {
	case slog.LevelDebug | slog.LevelWarn, slog.LevelError:
		writer = h.stderr
	default:
		writer = os.Stdout
	}
	return h.handleWriting(writer, log)
}
func (h *LogHandler) handleWriting(w io.Writer, r slog.Record) error {

	buff := bytes.NewBuffer(make([]byte, 0))

	fmt.Fprintf(w, "%s [%s] %s ", r.Time.Format(h.timeFormat), h.name, r.Level)
	for _, attr := range h.attrs {
		fmt.Fprintf(buff, "'%s'='%v' ", attr.Key, attr.Value)
	}

	if h.trace {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		fmt.Fprintf(buff, "\n - file: %s:%d\n - function: %s", f.File, f.Line, f.Function)
	}
	fmt.Fprintf(buff, "\n%s\n\n", r.Message)
	h.mux.Lock()
	defer h.mux.Unlock()
	_, err := io.Copy(w, buff)
	return err
}

func (h *LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &LogHandler{
		name:       h.name,
		verbose:    h.verbose,
		trace:      h.trace,
		stdout:     os.Stdout,
		stderr:     os.Stderr,
		attrs:      append(h.attrs, attrs...),
		timeFormat: h.timeFormat,
		mux:        h.mux,
	}
}
func (h *LogHandler) WithGroup(name string) slog.Handler {
	return &LogHandler{
		name:       fmt.Sprintf("%s.%s", h.name, name),
		verbose:    h.verbose,
		trace:      h.trace,
		stdout:     os.Stdout,
		stderr:     os.Stderr,
		attrs:      h.attrs,
		timeFormat: h.timeFormat,
		mux:        h.mux,
	}
}
