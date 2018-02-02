package sweetygo

import (
	"html/template"
	"io"
	"io/ioutil"
	"path"
	"strings"
)

// Templates is a templates manager.
type Templates struct {
	Root     string
	Suffix   string
	template *template.Template
	FuncMap  template.FuncMap
}

// NewTemplates .
func NewTemplates(rootDir string) *Templates {
	tpl := &Templates{
		Root:   rootDir,
		Suffix: ".html",
	}
	tpl.loadTpls()
	return tpl
}

// Render Templates.
func (tpl *Templates) Render(w io.Writer, tplname string, data interface{}) error {
	return tpl.template.ExecuteTemplate(w, tplname, data)
}

func (tpl *Templates) loadTpls() {
	tpl.template = template.New("_SweetyGo_").
		Funcs(tpl.FuncMap)
	tpls, err := tpl.listDir()
	if err != nil {
		return
	}
	for _, t := range tpls {
		tpl.parseFile(t)
	}
}

func (tpl *Templates) listDir() ([]string, error) {
	files := make([]string, 0)
	dir, err := ioutil.ReadDir(tpl.Root)
	if err != nil {
		return nil, err
	}
	for _, f := range dir {
		if f.IsDir() {
			continue
		}
		if strings.HasSuffix(f.Name(), tpl.Suffix) {
			files = append(files, f.Name())
		}
	}
	return files, nil
}

func (tpl *Templates) parseFile(filename string) error {
	b, err := ioutil.ReadFile(path.Join(tpl.Root, filename))
	if err != nil {
		return err
	}
	t := tpl.template.Lookup(filename)
	if t == nil {
		t = tpl.template.New(filename)
	}
	t, err = t.Parse(string(b))
	if err != nil {
		return err
	}
	return nil
}
