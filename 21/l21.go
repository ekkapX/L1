package main

import "fmt"

/*
Реализовать паттерн проектирования «Адаптер» на любом примере.

Описание: паттерн Adapter позволяет сконвертировать интерфейс одного класса в интерфейс другого, который ожидает клиент.

Продемонстрируйте на простом примере в Go: у вас есть существующий интерфейс (или структура) и другой, несовместимый по интерфейсу потребитель — напишите адаптер, который реализует нужный интерфейс и делегирует вызовы к встроенному объекту.

Поясните применимость паттерна, его плюсы и минусы, а также приведите реальные примеры использования.
*/

type LegacyLogger interface {
	Log(msg string)
}

type FileLogger struct{}

func (f *FileLogger) Log(msg string) {
	fmt.Printf("[FileLogger] %s\n", msg)
}

type ModernLogger interface {
	Info(msg string)
	Error(msg string)
}

type ModernLoggerAdapter struct {
	legacyLogger LegacyLogger
}

func (m *ModernLoggerAdapter) Info(msg string) {
	m.legacyLogger.Log("[INFO] " + msg)
}

func (m *ModernLoggerAdapter) Error(msg string) {
	m.legacyLogger.Log("[ERROR] " + msg)
}

func useModernLogger(logger ModernLogger) {
	logger.Info("This is info")
	logger.Error("This is error")
}

func main() {
	oldLogger := &FileLogger{}
	ModernLoggerAdapter := &ModernLoggerAdapter{legacyLogger: oldLogger}
	useModernLogger(ModernLoggerAdapter)
}

// Паттерн применяется, когда интерфейс существующей структуры не соответствует тому, который ожидает клиент.
// Например, в L0 я реализовал адаптер, преобразующий логгер из библиотеки Goose в совместимый с zap.Logger,
// чтобы обеспечить консистентность логирования по всему проекту.
// Адаптер соблюдает принцип Open-Closed: мы можем повторно использовать существующий код без его модификации.
// Из минусов - дополнительный слой кода, требующийся для реализации адаптера.
