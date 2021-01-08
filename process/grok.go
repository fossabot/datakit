package process

import (
	"fmt"
	"path/filepath"

	vgrok "github.com/vjeantet/grok"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/process/parser"
)

var (
	grokCfg *vgrok.Grok
)

func Grok(p *Pipeline, node parser.Node) (*Pipeline, error) {
	funcExpr := node.(*parser.FuncExpr)
	if len(funcExpr.Param) != 2 {
		return p, fmt.Errorf("func %s expected 2 args", funcExpr.Name)
	}

	pattern := funcExpr.Param[0].(*parser.StringLiteral).Val
	key     := funcExpr.Param[1].(*parser.Identifier).Name

	val := p.getContentStr(key)
	m, err := p.grok.Parse(pattern, val)
	if err != nil {
		return p, err
	}

	for k, v := range m {
		p.setContent(k, v)
	}

    return p, nil
}

func LoadPatterns() {
	g, err := vgrok.NewWithConfig(&vgrok.Config{
		NamedCapturesOnly: true,
		PatternsDir:[]string{
			filepath.Join(datakit.InstallDir, "pattern"),
		},
	})

	if err != nil {
		l.Errorf("grok init err: %v", err)
	}

	grokCfg = g
}
