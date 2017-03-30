package levels

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"os"
	"reflect"
	"strings"
)

var templateData map[string]interface{}

var funcMap = template.FuncMap{
	"generate_levels": func(name string, levels interface{}, format string) string {
		lvls := reflect.ValueOf(levels)
		level_names := make([]string, lvls.Len())
		for i := 0; i < lvls.Len(); i++ {
			level_names[i] = fmt.Sprintf(format, name, i+1)
		}
		return strings.Join(level_names, ", ")
	},
	"add": func(num int, add int) int {
		return num + add
	},
}

func Template(templ []byte, yamlData []byte) {
	yaml_err := yaml.Unmarshal(yamlData, &templateData)
	if yaml_err != nil {
		panic(yaml_err)
	}
	t, tmpl_err := template.New("template").Funcs(funcMap).Parse(string(templ))
	if tmpl_err != nil {
		panic(tmpl_err)
	}
	t.Execute(os.Stdout, &templateData)
}
