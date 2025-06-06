package benchmark

import (
	"context"
	"fmt"
	"log/slog"
	"testing"
)

func BenchmarkTextDisabledWithoutFields(b *testing.B) {
	b.Logf("Logging at a disabled level without any structured context.")
	b.Run("slog text", func(b *testing.B) {
		logger := newDisabledSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0))
			}
		})
	})
	b.Run("custom slog text", func(b *testing.B) {
		logger := newDisabledCustomSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0))
			}
		})
	})
	b.Run("slog.LogAttrs text", func(b *testing.B) {
		logger := newDisabledSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.LogAttrs(context.Background(), slog.LevelInfo, getMessage(0))
			}
		})
	})
	b.Run("custom slog.LogAttrs text", func(b *testing.B) {
		logger := newDisabledCustomSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.LogAttrs(context.Background(), slog.LevelInfo, getMessage(0))
			}
		})
	})

	fmt.Println("=========================================================")
}

func BenchmarkTextDisabledAccumulatedContext(b *testing.B) {
	b.Logf("Logging at a disabled level with some accumulated context.")
	b.Run("slog text", func(b *testing.B) {
		logger := newDisabledSlogText(fakeSlogFields()...)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0))
			}
		})
	})
	b.Run("custom slog text", func(b *testing.B) {
		logger := newDisabledCustomSlogText(fakeSlogFields()...)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0))
			}
		})
	})
	b.Run("slog.LogAttrs", func(b *testing.B) {
		logger := newDisabledSlogText(fakeSlogFields()...)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.LogAttrs(context.Background(), slog.LevelInfo, getMessage(0))
			}
		})
	})
	b.Run("custom slog.LogAttrs", func(b *testing.B) {
		logger := newDisabledCustomSlogText(fakeSlogFields()...)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.LogAttrs(context.Background(), slog.LevelInfo, getMessage(0))
			}
		})
	})

	fmt.Println("=========================================================")
}

func BenchmarkTextDisabledAddingFields(b *testing.B) {
	b.Logf("Logging at a disabled level, adding context at each log site.")
	b.Run("slog", func(b *testing.B) {
		logger := newDisabledSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0), fakeSlogArgs()...)
			}
		})
	})
	b.Run("custom slog", func(b *testing.B) {
		logger := newDisabledCustomSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0), fakeSlogArgs()...)
			}
		})
	})
	b.Run("slog.LogAttrs", func(b *testing.B) {
		logger := newDisabledSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.LogAttrs(context.Background(), slog.LevelInfo, getMessage(0), fakeSlogFields()...)
			}
		})
	})
	b.Run("custom slog.LogAttrs", func(b *testing.B) {
		logger := newDisabledCustomSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.LogAttrs(context.Background(), slog.LevelInfo, getMessage(0), fakeSlogFields()...)
			}
		})
	})

	fmt.Println("=========================================================")
}

func BenchmarkTextWithoutFields(b *testing.B) {
	b.Logf("Logging without any structured context.")
	b.Run("slog", func(b *testing.B) {
		logger := newSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0))
			}
		})
	})
	b.Run("custom slog", func(b *testing.B) {
		logger := newCustomSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0))
			}
		})
	})
	b.Run("slog.LogAttrs", func(b *testing.B) {
		logger := newSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.LogAttrs(context.Background(), slog.LevelInfo, getMessage(0))
			}
		})
	})
	b.Run("custom slog.LogAttrs", func(b *testing.B) {
		logger := newCustomSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.LogAttrs(context.Background(), slog.LevelInfo, getMessage(0))
			}
		})
	})

	fmt.Println("=========================================================")
}

func BenchmarkTextAccumulatedContext(b *testing.B) {
	b.Logf("Logging with some accumulated context.")
	b.Run("slog", func(b *testing.B) {
		logger := newSlogText(fakeSlogFields()...)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0))
			}
		})
	})
	b.Run("custom slog", func(b *testing.B) {
		logger := newCustomSlogText(fakeSlogFields()...)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0))
			}
		})
	})
	b.Run("slog.LogAttrs", func(b *testing.B) {
		logger := newSlogText(fakeSlogFields()...)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.LogAttrs(context.Background(), slog.LevelInfo, getMessage(0))
			}
		})
	})
	b.Run("custom slog.LogAttrs", func(b *testing.B) {
		logger := newCustomSlogText(fakeSlogFields()...)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.LogAttrs(context.Background(), slog.LevelInfo, getMessage(0))
			}
		})
	})

	fmt.Println("=========================================================")
}

func BenchmarkAddingFields(b *testing.B) {
	b.Logf("Logging with additional context at each log site.")
	b.Run("slog", func(b *testing.B) {
		logger := newSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0), fakeSlogArgs()...)
			}
		})
	})
	b.Run("custom slog", func(b *testing.B) {
		logger := newCustomSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(getMessage(0), fakeSlogArgs()...)
			}
		})
	})
	b.Run("slog.LogAttrs", func(b *testing.B) {
		logger := newSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.LogAttrs(context.Background(), slog.LevelInfo, getMessage(0), fakeSlogFields()...)
			}
		})
	})
	b.Run("custom slog.LogAttrs", func(b *testing.B) {
		logger := newCustomSlogText()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.LogAttrs(context.Background(), slog.LevelInfo, getMessage(0), fakeSlogFields()...)
			}
		})
	})

	fmt.Println("=========================================================")
}
