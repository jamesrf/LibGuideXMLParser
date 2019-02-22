package libguidesxml

import (
	"bufio"
	"encoding/xml"
	"log"
	"os"
)





func Read(filename string) (lg LibGuides) {
	xmlFile, err := os.Open("export.xml")
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	cr := BadCharCleaner{bufio.NewReader(xmlFile)}
	dec := xml.NewDecoder(cr)

	log.Println("Reading XML file: ...")
	err = dec.Decode(&lg)
	if err != nil {
		panic(err)
	}
	log.Println("Done! (%d bytes)\n\n", dec.InputOffset())
	return
}
