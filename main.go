package main

import (
	"github.com/keidarcy/zzw-food-gallery/util"
)

func main() {
	// move files from ~/Downloads/x.HEIC to ./heic-images
	util.Move()
	// get heic
	// convert to jpg
	util.Convert()
	// upload jpg to s3
	util.Upload()
	// remove heic and jpg
	util.Clean()
	// get from s3
	// render html
	util.Render()
}
