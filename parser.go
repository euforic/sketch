package sketch

import (
	"archive/zip"
	"encoding/json"
	"strings"
)

type File struct {
	Document Document
	Pages    map[string]Page
}

// Parse will un-compress a sketch file,
// and parse the contents
func Parse(src string) (*File, error) {
	sketchFile := File{
		Pages: map[string]Page{},
	}

	zr, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	for _, f := range zr.File {
		name := f.FileHeader.Name

		if strings.HasPrefix(name, "pages/") {
			page := Page{}
			if err := parseObj(f, &page); err != nil {
				return nil, err
			}
			sketchFile.Pages[page.Name] = page
		}

		if name == "document.json" {
			doc := Document{}
			if err := parseObj(f, &doc); err != nil {
				return nil, err
			}
			sketchFile.Document = doc
		}
	}
	return &sketchFile, nil
}

func parseObj(file *zip.File, dst interface{}) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(dst); err != nil {
		return err
	}
	return nil
}
