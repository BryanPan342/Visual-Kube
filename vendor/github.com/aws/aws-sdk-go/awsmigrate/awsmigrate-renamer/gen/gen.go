package main

import (
	"bytes"
	"go/format"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/aws/aws-sdk-go/private/model/api"
)

type pkg struct {
	oldAPI     *api.API
	newAPI     *api.API
	shapes     map[string]*shapentry
	operations map[string]*opentry
}

type shapentry struct {
	oldShape *api.Shape
	newShape *api.Shape
}

type opentry struct {
	oldName string
	newName string
}

type packageRenames struct {
	Shapes     map[string]string
	Operations map[string]string
	Fields     map[string]string
}

var exportMap = map[string]*packageRenames{}

func generateRenames(w io.Writer) error {
	tmpl, err := template.New("renames").Parse(t)
	if err != nil {
		return err
	}

	out := bytes.NewBuffer(nil)
	if err = tmpl.Execute(out, exportMap); err != nil {
		return err
	}

	b, err := format.Source(bytes.TrimSpace(out.Bytes()))
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bytes.NewReader(b))
	return err
}

const t = `
package rename

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

var renamedPackages = map[string]*packageRenames{
	{{ range $key, $entry := . }}"{{ $key }}": &packageRenames{
		operations: map[string]string{
			{{ range $old, $new := $entry.Operations }}"{{ $old }}": "{{ $new }}",
			{{ end }}
		},
		shapes: map[string]string{
			{{ range $old, $new := $entry.Shapes }}"{{ $old }}": "{{ $new }}",
			{{ end }}
		},
		fields: map[string]string{
			{{ range $old, $new := $entry.Fields }}"{{ $old }}": "{{ $new }}",
			{{ end }}
		},
	},
	{{ end }}
}
`

func (p *pkg) buildRenames() {
	pkgName := "github.com/aws/aws-sdk-go/service/" + p.oldAPI.PackageName()
	if exportMap[pkgName] == nil {
		exportMap[pkgName] = &packageRenames{map[string]string{}, map[string]string{}, map[string]string{}}
	}
	ifacename := "github.com/aws/aws-sdk-go/service/" + p.oldAPI.PackageName() + "/" +
		p.oldAPI.InterfacePackageName()
	if exportMap[ifacename] == nil {
		exportMap[ifacename] = &packageRenames{map[string]string{}, map[string]string{}, map[string]string{}}
	}

	for _, entry := range p.operations {
		if entry.oldName != entry.newName {
			pkgNames := []string{pkgName, ifacename}
			for _, p := range pkgNames {
				exportMap[p].Operations[entry.oldName] = entry.newName
				exportMap[p].Operations[entry.oldName+"Request"] = entry.newName + "Request"
				exportMap[p].Operations[entry.oldName+"Pages"] = entry.newName + "Pages"
			}
		}
	}

	for _, entry := range p.shapes {
		if entry.oldShape.Type == "structure" {
			if entry.oldShape.ShapeName != entry.newShape.ShapeName {
				exportMap[pkgName].Shapes[entry.oldShape.ShapeName] = entry.newShape.ShapeName
			}

			for _, n := range entry.oldShape.MemberNames() {
				for _, m := range entry.newShape.MemberNames() {
					if n != m && strings.ToLower(n) == strings.ToLower(m) {
						exportMap[pkgName].Fields[n] = m
					}
				}
			}
		}
	}
}

func load(file string) *pkg {
	p := &pkg{&api.API{}, &api.API{}, map[string]*shapentry{}, map[string]*opentry{}}

	p.oldAPI.Attach(file)
	p.oldAPI.Setup()

	p.newAPI.Attach(file)
	p.newAPI.Setup()

	for _, name := range p.oldAPI.OperationNames() {
		p.operations[strings.ToLower(name)] = &opentry{oldName: name}
	}

	for _, name := range p.newAPI.OperationNames() {
		p.operations[strings.ToLower(name)].newName = name
	}

	for _, shape := range p.oldAPI.ShapeList() {
		p.shapes[strings.ToLower(shape.ShapeName)] = &shapentry{oldShape: shape}
	}

	for _, shape := range p.newAPI.ShapeList() {
		if _, ok := p.shapes[strings.ToLower(shape.ShapeName)]; !ok {
			panic("missing shape " + shape.ShapeName)
		}
		p.shapes[strings.ToLower(shape.ShapeName)].newShape = shape
	}

	return p
}

var excludeServices = map[string]struct{}{
	"simpledb":     {},
	"importexport": {},
}

func main() {
	files, _ := filepath.Glob("../../apis/*/*/api-2.json")

	sort.Strings(files)

	// Remove old API versions from list
	m := map[string]bool{}
	for i := range files {
		idx := len(files) - 1 - i
		parts := strings.Split(files[idx], string(filepath.Separator))
		svc := parts[len(parts)-3] // service name is 2nd-to-last component

		if m[svc] {
			files[idx] = "" // wipe this one out if we already saw the service
		}
		m[svc] = true
	}

	for i := range files {
		file := files[i]
		if file == "" { // empty file
			continue
		}

		if g := load(file); g != nil {
			if _, ok := excludeServices[g.oldAPI.PackageName()]; !ok {
				g.buildRenames()
			}
		}
	}

	outfile, err := os.Create("rename/renames.go")
	if err != nil {
		panic(err)
	}
	if err := generateRenames(outfile); err != nil {
		panic(err)
	}
}
