package modules

import "os"

type file struct {

}

func NewFile() *file {
	return &file{}
}

func (fs *file) IsExist(pathname string) (ok bool) {
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
