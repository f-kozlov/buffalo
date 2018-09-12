package build

import (
	"fmt"

	"github.com/gobuffalo/buffalo/generators/assets/webpack"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr/jam/parser"
	"github.com/gobuffalo/packr/jam/store"
	"github.com/pkg/errors"
)

func (b *Builder) buildAssets() error {

	if b.WithWebpack && b.Options.WithAssets {
		if err := envy.MustSet("NODE_ENV", envy.Get("GO_ENV", "production")); err != nil {
			return errors.WithStack(err)
		}
		if err := b.exec(webpack.BinPath); err != nil {
			return errors.WithStack(err)
		}
	}

	p, err := parser.NewFromRoots([]string{b.Root})
	if err != nil {
		return errors.WithStack(err)
	}
	boxes, err := p.Run()
	if err != nil {
		return errors.WithStack(err)
	}

	// reduce boxes - remove ones we don't want
	// MB: current assumption is we want all these
	// boxes, just adding a comment suggesting they're
	// might be a reason to exclude some

	fmt.Printf("Found %d boxes\n", len(boxes))

	// "pack" boxes
	d := store.NewDisk("", "")
	for _, b := range boxes {
		if err := d.Pack(b); err != nil {
			return errors.WithStack(err)
		}
	}
	return d.Close()
	// p := pack.New(b.ctx, b.Root)
	// p.Compress = b.Compress

	// if !b.Options.WithAssets {
	// 	p.IgnoredBoxes = append(p.IgnoredBoxes, "../public/assets")
	// } else {
	// 	p.IgnoredFolders = p.IgnoredFolders[1:]
	// }
	//
	// if b.ExtractAssets && b.Options.WithAssets {
	// 	p.IgnoredBoxes = append(p.IgnoredBoxes, "../public/assets")
	// 	err := b.buildExtractedAssets()
	// 	if err != nil {
	// 		return errors.WithStack(err)
	// 	}
	// }

}
