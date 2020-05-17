package images

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/viniciusbds/navio/utilities"
)

var (
	contImages = make(map[string]*Image)
	baseImages = make(map[string]*Image)
)

func init() {

	if utilities.FileExists(utilities.BaseImagescsv) {
		readBaseImagesCSV()
	} else {
		baseImages["alpine"] = NewImage("alpine", "alpine", "v3.11", "2.7M", utilities.AlpineURL)
		baseImages["busybox"] = NewImage("busybox", "busybox", "v4.0", "1.5M", utilities.BusyboxURL)
		baseImages["ubuntu"] = NewImage("ubuntu", "ubuntu", "v20.04", "90.0M", utilities.Ubuntu20ltsURL)
		updateBaseImagesCSV()
	}

	if utilities.FileExists(utilities.ContImagescsv) {
		readContImagesCSV()
	}
}

// Image ...
type Image struct {
	name    string
	base    string
	version string
	size    string
	url     string
}

// NewImage ...
func NewImage(name string, base string, version string, size string, url string) *Image {
	return &Image{
		name:    name,
		base:    base,
		version: version,
		size:    size,
		url:     url,
	}
}

// ToStr ... Ignore esse código, essa foi a minha maior vigarice (https://www.youtube.com/watch?v=PK0c_n5EDhk)
func (i *Image) ToStr() string {
	tab := ""
	if len(i.name) < 8 {
		tab = "\t\t\t\t\t"
	} else if len(i.name) >= 8 && len(i.name) < 16 {
		tab = "\t\t\t\t"
	} else {
		tab = "\t\t\t"
	}
	return fmt.Sprintf("%s%s%s\t\t\t%s\t\t\t%s", i.name, tab, i.base, i.version, i.size)
}

func getImage(name string) *Image {
	if contImages[name] != nil {
		return contImages[name]
	}
	if baseImages[name] != nil {
		return baseImages[name]
	}
	return nil
}

// IsValidContainerImgName verifies if a Container Image Name wasn't used.
// If the containerImgName is available, return true.
func IsValidContainerImgName(containerImgName string) bool {
	if contImages[containerImgName] != nil {
		return true
	}
	return false
}

// InsertContImage receive a containerImgName and the respective baseImage and creates a new
// Image and insert it on [contImages map]. Also update the csv file that store all containerImages
func InsertContImage(containerImgName, baseImage string) {
	baseImg := getImage(baseImage)
	newImg := NewImage(containerImgName, baseImage, baseImg.version, baseImg.size, baseImg.url)
	contImages[containerImgName] = newImg
	updateContImagesCSV()
}

// InsertBaseImage ...
func InsertBaseImage(newBaseImg, base string) {
	baseImg := getImage(base)
	newImg := NewImage(newBaseImg, base, baseImg.version, baseImg.size, baseImg.url)
	baseImages[newBaseImg] = newImg
	updateBaseImagesCSV()
}

func removeContImage(contImamge string) {
	delete(contImages, contImamge)
	updateContImagesCSV()
}

func removeBaseImage(baseImage string) {
	delete(baseImages, baseImage)
	updateBaseImagesCSV()
}

// IsaBaseImage receive a imageName and return true if is a base image.
func IsaBaseImage(image string) bool {
	for _, i := range baseImages {
		if image == i.name {
			return true
		}
	}
	return false
}

func readContImagesCSV() {
	contImages = make(map[string]*Image)

	csvfile, err := os.Open(utilities.ContImagescsv)
	if err != nil {
		l.Log("Error", err.Error())
	}

	r := csv.NewReader(csvfile)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		contImages[record[0]] = NewImage(record[0], record[1], record[2], record[3], record[4])
	}
}

func updateContImagesCSV() {
	err := os.RemoveAll(utilities.ContImagescsv)
	if err != nil {
		log.Fatalln("Couldn't remove the csv file", err)
	}

	csvfile, err := os.Create(utilities.ContImagescsv)
	if err != nil {
		log.Fatalln("Couldn't create the csv file", err)
	}

	w := csv.NewWriter(csvfile)

	for _, v := range contImages {
		newimage := []string{v.name, v.base, v.version, v.size, v.url}
		err = w.Write(newimage)
		if err != nil {
			fmt.Println("Erro aque")
		}
	}
	w.Flush()
	csvfile.Close()
}

func readBaseImagesCSV() {
	baseImages = make(map[string]*Image)

	csvfile, err := os.Open(utilities.BaseImagescsv)
	if err != nil {
		l.Log("Error", err.Error())
	}

	r := csv.NewReader(csvfile)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		baseImages[record[0]] = NewImage(record[0], record[1], record[2], record[3], record[4])
	}
}

func updateBaseImagesCSV() {
	err := os.RemoveAll(utilities.BaseImagescsv)
	if err != nil {
		log.Fatalln("Couldn't remove the csv file", err)
	}

	if !utilities.FileExists(utilities.ImagesRootDir) {
		os.MkdirAll(utilities.ImagesRootDir, 0777)
	}

	csvfile, err := os.Create(utilities.BaseImagescsv)
	if err != nil {
		log.Fatalln("Couldn't create the csv file [picuinha]", err)
	}

	w := csv.NewWriter(csvfile)

	for _, v := range baseImages {
		newimage := []string{v.name, v.base, v.version, v.size, v.url}
		err = w.Write(newimage)
		if err != nil {
			fmt.Println("Erro aque")
		}
	}
	w.Flush()
	csvfile.Close()
}
