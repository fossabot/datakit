package process

import (
	"fmt"
	"strings"

	vgrok "github.com/vjeantet/grok"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/process/parser"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/process/patterns"
)

type Pipeline struct {
	Content   string
	Output    map[string]interface{}
	lastErr   error
	patterns  map[string]string    //存放自定义patterns
	nodes     []parser.Node
	grok      *vgrok.Grok
}

var (
	l = logger.DefaultSLogger("process")
)

func NewPipeline(script string) *Pipeline {
	p := &Pipeline{
		Output: make(map[string]interface{}),
	}
	if script != "" {
		p.nodes, p.lastErr = ParseScript(script)
	}
	p.grok = grokCfg
	return p
}

func (p *Pipeline) Run(data string) *Pipeline {
	var err error

	//防止脚本解析错误
	if p.lastErr != nil {
		return p
	}

	p.Content = data
	p.Output  = make(map[string]interface{})

	for _, node := range p.nodes {
		switch v := node.(type) {
		case *parser.FuncExpr:
			fn := strings.ToLower(v.Name)
			f, ok := funcsMap[fn]
			if !ok {
				err := fmt.Errorf("unsupported func: %v", v.Name)
				l.Error(err)
				p.lastErr = err
				return p
			}

			_, err = f(p, node)
			if err != nil {
				l.Errorf("ProcessLog func %v: %v", v.Name, err)
				p.lastErr = err
				return p
			}

		default:
			p.lastErr = fmt.Errorf("%v not function", v.String())
		}
	}
	return p
}

func (p *Pipeline) Result() map[string]interface{} {
	return p.Output
}

func (p *Pipeline) LastError() error {
	return p.lastErr
}

func (p *Pipeline) getContent(key string) interface{} {
	if key == "_" {
		return p.Content
	}

	var m interface{}
	var nm interface{}

	m = p.Output
	keys := strings.Split(key, ".")
	for _, k := range keys {
		switch m.(type) {
		case map[string]interface{}:
			v := m.(map[string]interface{})
			nm = v[k]
			m = nm
		default:
			return ""
		}
	}

	return nm
}

func (p *Pipeline) getContentStr(key string) string {
	content := p.getContent(key)

	switch v := content.(type) {
	case string:
		return v
	default:
		return fmt.Sprintf("%v", content)
	}

	return ""
}

func (p *Pipeline) setContent(k string, v interface{}) {
	if p.Output == nil {
		p.Output = make(map[string]interface{})
	}

	p.Output[k] = v
}

func init() {
	patterns.MkPatternsFile()
	LoadPatterns()
}
