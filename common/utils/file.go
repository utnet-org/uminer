package utils

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"uminer/common/errors"
)

func CopyFile(src, dest string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return errors.Errorf(err, errors.ErrorOperateFileFailed)
	}
	defer srcfd.Close()

	destPath := filepath.Dir(dest)
	if !DirExists(destPath) {
		_ = os.MkdirAll(destPath, 0755)
	}

	if dstfd, err = os.Create(dest); err != nil {
		return errors.Errorf(err, errors.ErrorOperateFileFailed)
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return errors.Errorf(err, errors.ErrorOperateFileFailed)
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return errors.Errorf(err, errors.ErrorOperateFileFailed)
	}
	return os.Chmod(dest, srcinfo.Mode())
}

func CopyDir(src string, dest string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return errors.Errorf(err, errors.ErrorOperateDirFailed)
	}

	if err = os.MkdirAll(dest, srcinfo.Mode()); err != nil {
		return errors.Errorf(err, errors.ErrorOperateDirFailed)
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return errors.Errorf(err, errors.ErrorOperateDirFailed)
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dest, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				return err
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				return err
			}
		}
	}
	return nil
}

func DeleteDir(src string) error {
	if !DirExists(src) {
		return errors.Errorf(nil, errors.ErrorDirNotExisted)
	}

	err := os.RemoveAll(src)
	if err != nil {
		return errors.Errorf(err, errors.ErrorOperateDirFailed)
	}

	return nil
}

func CreateDir(src string) error {

	if DirExists(src) {
		return nil
	}

	_ = os.MkdirAll(src, 0755)

	return nil
}

func DirExists(src string) bool {
	fi, err := os.Stat(src)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}
