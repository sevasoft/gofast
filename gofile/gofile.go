package gofile

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func GetInfo(_path string) (fs.FileInfo, error) {
	return os.Stat(_path)
}

func GetName(_path string) string {
	file, err := GetInfo(_path)

	if err != nil {
		log.Fatalln("Error!", err.Error())
	}

	return file.Name()
}

func IsExist(_path string) bool {
	file, err := GetInfo(_path)

	if err == nil && os.IsExist(err) {
		name := file.Name()
		log.Printf("File %s already exists.\n", name)

		return true
	} else if os.IsNotExist(err) {
		log.Println("Provided file does not exist.")
		return false
	}

	return false
}

func Create(_path string, _overwrite bool) error {
	if _overwrite == true {
		log.Println("Creating/overwriting file ...")
		_, err := os.Create(_path)

		if err != nil {
			log.Printf("X-X-X File is not created X-X-X %s\n%s\n", _path, err.Error())
		} else {
			log.Printf("File %s is created.\n", GetName(_path))
		}

		return err
	}

	if _overwrite == false {
		if IsExist(_path) == false {

			log.Println("Creating new file ...")
			_, err := os.Create(_path)

			if err != nil {
				log.Printf("X-X-X File is not created X-X-X %s\n%s\n", _path, err.Error())
			} else {
				log.Printf("File %s is created.\n", GetName(_path))
			}

			return err
		}
	}

	return nil
}

func ReadString(_path string) string {
	name := GetName(_path)
	content, err := ioutil.ReadFile(_path)

	if err != nil {
		log.Fatalf("Can not read the file %s\n", name)
	}

	return string(content)
}

func ReadByte(_path string) ([]byte, error) {
	name := GetName(_path)
	content, err := ioutil.ReadFile(_path)

	if err != nil {
		log.Printf("Can not read the file %s\n", name)
		return nil, err
	}

	return content, err
}

/*
   @param _path Path to file incl. file extension (ex: "./test.py")
   @param _data Data to write (ex: []byte("print('Hello World')"))
   @param _perm Unix file permission bits (ex: 0755)
*/
func Write(_path string, _data []byte, _perm uint32, _overwrite bool) error {
	unixFilePermissionBits := fs.FileMode(_perm)

	Create(_path, _overwrite) // creates if not exist.

	err := ioutil.WriteFile(_path, _data, unixFilePermissionBits)

	if err != nil {
		log.Println("X-X-X File is not written X-X-X", err.Error())
	}

	return err
}

func Rename(_oldPath string, _newPath string) {
	oldName := GetName(_oldPath)
	err := os.Rename(_oldPath, _newPath)

	if err != nil {
		log.Fatalln("Error!", err.Error())
	}

	newName := GetName(_newPath)

	log.Printf("Successfully renamed (moved) the file %s to %s\n", oldName, newName)
}

func RemoveAll(_path string) error {
	log.Printf("Removing %s\n", _path)
	err := os.RemoveAll(_path)

	if err != nil {
		log.Println("Error!", err.Error())
	} else {
		log.Printf("Successfully removed the file %s\n", _path)
	}

	return err
}

func GetExtension(_path string) string {
	return filepath.Ext(_path)
}
