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
	return &Templates{
		Root:   rootDir,
		Suffix: ".html",
	}
}

// Render Templates.
func (tpl *Templates) Render(w io.Writer, tplname string, data interface{}) error {
	return tpl.template.ExecuteTemplate(w, tplname, data)
}

func (tpl *Templates) loadTpls() {
	tpl.template = template.New("_SG_").
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
			filename := path.Join(tpl.Root, f.Name())
			files = append(files, filename)
		}
	}
	return files, nil
}

func (tpl *Templates) parseFile(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	t := tpl.template.Lookup(filename)
	if t == nil {
		t = tpl.template.New(filename)
	}
	_, err = tpl.template.Parse(string(b))
	if err != nil {
		return err
	}
	return nil
}
