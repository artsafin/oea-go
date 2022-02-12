package resources

import (
	"embed"
	"fmt"
	"html/template"
	"io"
)

//go:embed assets
var assets embed.FS

func TestReadFile() {
	bs, _ := assets.ReadFile("assets/layout.go.html")

	fmt.Printf("TestReadFile: %q\n", bs)
}

func MustParseTemplate(tmpl *template.Template, patterns ...string) *template.Template {
	return template.Must(tmpl.ParseFS(assets, patterns...))
}

func Open(file string) (io.ReadCloser, error) {
	return assets.Open(file)
}

func MustReadBytes(file string) []byte {
	bytes, err := assets.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return bytes
}