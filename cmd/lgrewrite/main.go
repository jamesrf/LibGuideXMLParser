package main

import (
	"LibGuideXMLParser/libguidesxml"
	"flag"
	"fmt"
	"log"
)

func main() {

	// using standard library "flag" package
	xmlfile := flag.String("xml", "export.xml", "LibGuides full export XML file to process")
	mapfile := flag.String("map", "", "Map file for remapping record ID -- old ID in first col, new ID 2nd col, TSV, blank for no remapping")
	regexpInput := flag.String("regex", "", "Regex string to find with brackets around record ID-- ie: oldcatalog/([0-9]+)")
	resultPrefix := flag.String("prefix", "", "New URL prefix: http://newcatalog.com/?record=")
	resultSuffix := flag.String("suffix", "", "New URL suffix (goes after $1)")
	trimSierraIds := flag.Bool("trim", true, "Trim first and last char from old ID")

	flag.Parse()

	log.Println("Using Regex:", *regexpInput)

	rewriter := libguidesxml.NewRewriter(*regexpInput, *mapfile, *trimSierraIds)

	lg := libguidesxml.Read(*xmlfile)
	for _, guide := range lg.Guides {
		//fmt.Printf("GUIDE:%s\n", guide.Name)
		for _, page := range guide.Pages {
			for _, box := range page.Boxes {
				for _, asset := range box.Assets {

					assetID := asset.ID

					oldUrl := asset.URL

					newId := rewriter.Rewrite(oldUrl)

					if newId != "" {
						newUrl := fmt.Sprintf("%s%s%s", *resultPrefix, newId, *resultSuffix)
						fmt.Printf("%s\t%s\t%s\n", assetID, oldUrl, newUrl)
					}

				}

			}
		}

	}
}
