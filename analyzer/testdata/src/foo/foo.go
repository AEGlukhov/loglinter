package foo

import (
	"log/slog"
)

func main() {

	slog.Debug("Failed to connect to database") // want "log message should start with lowercase letter: Failed to connect to database"
	slog.Debug("Debug!!!")                      // want "log message should start with lowercase letter: Debug!!!" "log message contains forbidden special characters: Debug!!!"

	slog.Info("ошибка подключения к базе данных") // want "log message must be in english: ошибка подключения к базе данных"
	slog.Info("Информация")                       // want "log message should start with lowercase letter: Информация" "log message must be in english: Информация"

	slog.Warn("connection failed!!!") // want "log message contains forbidden special characters: connection failed!!!"
	slog.Warn("предупреждение😀")      // want  "log message must be in english: предупреждение😀" "log message contains forbidden special characters: предупреждение😀"

	slog.Error("password") // want "log message contains sensitive data: password"
	slog.Error("api_key")  // want "log message contains forbidden special characters: api_key" "log message contains sensitive data: api_key"

}
