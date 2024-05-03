package logger

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
)

type LogStruct struct {
	TimeStamp   string `json:"time_stamp"`
	Status      string `json:"status"`
	Level       string `json:"level"`
	ServiceName string `json:"service_name"`
	HostName    string `json:"host_name"`
	MessageID   string `json:"message_id"`
	Message     string `json:"message"`
	UserID      string `json:"user_id"`
	ErrorCode   string `json:"error_code"`
	StackTrace  string `json:"stack_trace"`
	MetaData    string `json:"meta_data"`
}

type JsonHandler struct {
	slog.Handler
	l     *log.Logger
	sname string
	hname string
}

func (h *JsonHandler) Handle(_ context.Context, r slog.Record) error {
	attrs := make(map[string]string, r.NumAttrs())
	r.Attrs(func(attr slog.Attr) bool {
		attrs[attr.Key] = attr.Value.String()
		return true
	})
	metadata, err := json.Marshal(attrs)
	if err != nil {
		return err
	}

	logValue := LogStruct{
		TimeStamp:   r.Time.UTC().Format("2006-01-02T15:04:05.999Z"),
		Status:      attrs["status"],
		Level:       r.Level.String(),
		ServiceName: h.sname,
		HostName:    h.hname,
		MessageID:   attrs["message_id"],
		Message:     r.Message,
		UserID:      attrs["user_id"],
		ErrorCode:   attrs["error_code"],
		StackTrace:  attrs["stack_trace"],
		MetaData:    string(metadata),
	}

	var b []byte
	b, err = json.Marshal(logValue)
	if err != nil {
		return err
	}

	h.l.Println(string(b))

	return nil
}

func NewJsonHandler(out io.Writer, opts HandlerOptions) *JsonHandler {
	h := &JsonHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
		sname:   opts.Service,
		hname:   opts.Host,
	}

	return h
}
