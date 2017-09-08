package libguidesxml
import(
	"os"
	"bufio"
	"encoding/xml"
	"fmt"
)

func Read(filename string) (lg LibGuides) {
	xmlFile, err := os.Open("export.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	cr := BadCharCleaner{bufio.NewReader(xmlFile)}
	dec := xml.NewDecoder(cr)

	fmt.Printf("Reading XML file: ...")
	err = dec.Decode(&lg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Done! (%d bytes)\n\n", dec.InputOffset())
	return
}
