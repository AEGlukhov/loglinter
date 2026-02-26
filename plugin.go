package loglinter

import (
	"github.com/aeglukhov/loglinter/analyzer"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

type Plugin struct{}

func New(settings any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
