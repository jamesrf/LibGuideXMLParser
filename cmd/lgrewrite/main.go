package main

import (
	"flag"

	"github.com/jamesrf/parseLibGuides/libguidesxml"
	// "os"
)

func main() {

	// using standard library "flag" package
	xmlfile := flag.String("xml", "export.xml", "LibGuides full export XML file to process")

	lg := libguidesxml.Read(*xmlfile)
	ew := libguidesxml.NewEncoreRewriter("encore.vcc.ca")
	wo := libguidesxml.NewWebPacOutputter("webpac.vcc.ca")

	for _, guide := range lg.Guides {
		//fmt.Printf("GUIDE:%s\n", guide.Name)
		for _, page := range guide.Pages {
			for _, box := range page.Boxes {
				for _, asset := range box.Assets {
					ew.RewriteAssetToWebPac(wo, &asset)
				}

			}
		}

	}
}
