package main

import (
	"os"

	"github.com/keidarcy/zzw-food-gallery/pkg/util"
)

func main() {
	ci := false
	if len(os.Args) > 1 {
		ci = true
	}

	if !ci {
		// local
		// move files from ~/Downloads/x.HEIC to ./heic-images
		util.Move()
		// get heic
		// convert to jpg
		util.Convert()
		// upload jpg to s3
		util.Upload()
		// remove heic and jpg
		util.Clean()
	} else {
		// get from s3
		// render html
		util.Render()
	}
}
