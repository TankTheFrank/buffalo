package standard

import (
	"path/filepath"

	"github.com/gobuffalo/buffalo/generators"
	"github.com/gobuffalo/buffalo/generators/assets"
	"github.com/gobuffalo/makr"
	"github.com/pkg/errors"
)

var logo = &makr.RemoteFile{
	File:       makr.NewFile("public/assets/images/logo.svg", ""),
	RemotePath: assets.LogoURL,
}

// Run standard assets generator for those wishing to not use webpack
func Run(root string, data makr.Data) error {
	files, err := generators.Find(filepath.Join(generators.TemplatesPath, "assets", "standard"))
	if err != nil {
		return errors.WithStack(err)
	}
	g := makr.New()
	for _, f := range files {
		g.Add(makr.NewFile(f.WritePath, f.Body))
	}
	g.Add(logo)
	return g.Run(root, data)
}
