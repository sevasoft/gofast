package godir

import (
	"errors"
	"io/fs"
	"log"
	"os"

	"github.com/vsevdrob/gofast/gofile"
)

func MkdirAll(_path string, _perm uint32, _overwrite bool) error {
	_, err := gofile.GetInfo(_path)
	unixFilePermBits := fs.FileMode(_perm)

	if _overwrite {
		log.Printf("Creating/overwriting directory %s\n", _path)
		err := os.RemoveAll(_path)
		err = os.MkdirAll(_path, unixFilePermBits)

		if err != nil && !os.IsExist(err) {
			log.Println("Error!", err.Error())
		} else {
			log.Printf("Successfully overwrited directory %s\n", _path)
		}

		return err
	}

	if _overwrite == false {
		if os.IsNotExist(err) {
			err := os.MkdirAll(_path, unixFilePermBits)

			if err != nil && !os.IsExist(err) {
				log.Println("Error!", err.Error())
			} else {
				log.Printf("Successfully created directory %s\n", _path)
			}

			return err
		} else {
			log.Printf("Directory %s already exists.", _path)
			return err
		}
	}

	return errors.New("Error occured while creating directory.")
}
