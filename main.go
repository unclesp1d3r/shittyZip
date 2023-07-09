package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexmullins/zip"
	"github.com/charmbracelet/log"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		log.Error("Usage: zip <zip file> <items to include> ")
		log.Info("Example: zip myzipfile.zip file1 file2 directory1 directory2")
		os.Exit(1)
	}

	destinationPath := args[0]
	itemsToZip := args[1:]

	log.Infof("Creating zip file: %s", destinationPath)
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		log.Fatal(err)
	}
	myZip := zip.NewWriter(destinationFile)

	for _, item := range itemsToZip {
		itemInfo, err := os.Stat(item)
		if err != nil {
			log.Fatal(err)
		}
		if itemInfo.IsDir() {
			err = zipDir(item, myZip)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err = zipFile(item, myZip)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err = myZip.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func zipDir(pathToZip string, myZip *zip.Writer) error {
	return filepath.Walk(
		pathToZip,
		func(filePath string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if err != nil {
				return err
			}
			return zipFile(filePath, myZip)
		},
	)
}

func zipFile(filePath string, myZip *zip.Writer) error {
	relPath := strings.TrimPrefix(filePath, string(os.PathSeparator))
	zipFile, err := myZip.Create(relPath)
	if err != nil {
		return err
	}
	log.Infof("Adding file: %s", relPath)
	fsFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(zipFile, fsFile)
	if err != nil {
		return err
	}
	return nil
}
