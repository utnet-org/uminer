package utils

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"uminer/common/errors"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// srcFile could be a single file or a directory
func Zip(srcFile string, destZip string) error {
	destPath := filepath.Dir(destZip)
	if !DirExists(destPath) {
		_ = os.MkdirAll(destPath, 0755)
	}

	zipfile, err := os.Create(destZip)
	if err != nil {
		return errors.Errorf(err, errors.ErrorZipFailed)
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	fi, err := os.Stat(srcFile)
	if err != nil {
		return errors.Errorf(err, errors.ErrorZipFailed)
	}
	if fi.IsDir() {
		srcFile += "/"
	}

	err = filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Errorf(err, errors.ErrorZipFailed)
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return errors.Errorf(err, errors.ErrorZipFailed)
		}

		header.Flags = 1 << 11 // use utf8
		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return errors.Errorf(err, errors.ErrorZipFailed)
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return errors.Errorf(err, errors.ErrorZipFailed)
			}
			defer file.Close()
			_, _ = io.Copy(writer, file)
		}
		if err != nil {
			return errors.Errorf(err, errors.ErrorZipFailed)
		} else {
			return nil
		}
	})
	if err != nil {
		return errors.Errorf(err, errors.ErrorZipFailed)
	}

	return nil
}

func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return errors.Errorf(err, errors.ErrorUnzipFailed)
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		// coding conversion
		decodeName := ""
		if f.Flags == 0 {
			i := bytes.NewReader([]byte(f.Name))
			decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
			content, _ := ioutil.ReadAll(decoder)
			decodeName = string(content)
		} else {
			decodeName = f.Name
		}

		fpath := filepath.Join(destDir, decodeName)
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return errors.Errorf(err, errors.ErrorUnzipFailed)
			}

			inFile, err := f.Open()
			if err != nil {
				return errors.Errorf(err, errors.ErrorUnzipFailed)
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				inFile.Close()
				return errors.Errorf(err, errors.ErrorUnzipFailed)
			}

			_, err = io.Copy(outFile, inFile)
			inFile.Close()
			outFile.Close()
			if err != nil {
				return errors.Errorf(err, errors.ErrorUnzipFailed)
			}
		}
	}

	return nil
}
