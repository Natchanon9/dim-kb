package standardlibrary

import (
	"log/slog"
	"os"
	"time"
)

func Example_slog() {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr { // drop time so output is stable
			if a.Key == slog.TimeKey && len(groups) == 0 {
				return slog.Attr{}
			}
			return a
		},
	})
	logger := slog.New(h)
	logger.Info("order created", "id", 42, "total", 99.5)
	logger.With("request_id", "r1").Warn("slow query", slog.Duration("took", 2*time.Second))
	logger.Debug("hidden") // below the configured level
	// Output:
	// {"level":"INFO","msg":"order created","id":42,"total":99.5}
	// {"level":"WARN","msg":"slow query","request_id":"r1","took":2000000000}
}
