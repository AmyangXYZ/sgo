package sweetygo

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// Templates is a templates manager.
type Templates struct {
	Dir      string
	Suffix   string
	template *template.Template
	FuncMap  template.FuncMap
}

// NewTemplates .
func NewTemplates(tplDir string, funcMap template.FuncMap) *Templates {
	tpl := &Templates{
		Dir:     tplDir,
		Suffix:  ".html",
		FuncMap: funcMap,
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
	tpls, err := tpl.walkDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	for _, t := range tpls {
		tpl.parseFile(t)
	}
}

func (tpl *Templates) walkDir() ([]string, error) {
	files := make([]string, 0)
	err := filepath.Walk(tpl.Dir, func(filename string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, filename[len(tpl.Dir)+1:])
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (tpl *Templates) parseFile(filename string) error {
	b, err := ioutil.ReadFile(path.Join(tpl.Dir, filename))
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
