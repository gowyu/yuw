package modules

import "os"

type File struct {

}

func NewFile() *File {
	return &File{}
}

func (fs *File) IsExist(pathname string) (ok bool) {
	_, err := os.Stat(pathname)
	if err != nil {
		if os.IsExist(err) {
			ok = true
			return
		}

		return
	}

	return
}
