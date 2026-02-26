package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglinter",
	Doc:  "checks log messages",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {

			// проверяем что этот node - это вызов функции
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			// проверяем что этот вызов функции - это вызов метода
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			// проверяем что этот вызов метода - это логирование
			if !isLogFunction(sel.Sel.Name) {
				return true
			}

			// проверяем что переданы аргументы
			if len(call.Args) == 0 {
				return true
			}

			// проверяем что переданный аргумент - это базовый литерал, а не переменная
			arg, ok := call.Args[0].(*ast.BasicLit)
			if !ok {
				return true
			}

			msg := strings.Trim(arg.Value, `"`)

			if len(msg) == 0 {
				pass.Reportf(call.Pos(), "log message is empty")
				return true
			}

			checkMessage(pass, call.Pos(), msg)

			return true
		})
	}
	return nil, nil
}

// эта функция проверяет только самые частые случаи использования логов, но не logger.Log(zapcore.InfoLevel, "some info"), slog.InfoContext(ctx, "some info") и т.п.
func isLogFunction(name string) bool {
	switch name {
	case "Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal":
		return true
	}
	return false
}
