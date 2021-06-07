package excel

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Reader struct {
	Path      string
	SheetName string
	StartRow  int
	NameCol   string
	CodeCol   string
	PhoneCol  string
}

func NewReader(path string, sheetName string, startRow int,
	nameCol string, codeCol string, phoneCol string) *Reader {

	return &Reader{
		Path:      path,
		SheetName: sheetName,
		StartRow:  startRow,
		NameCol:   nameCol,
		CodeCol:   codeCol,
		PhoneCol:  phoneCol,
	}
}

func (r *Reader) GetUsers() ([]User, error) {

	f, err := excelize.OpenFile(r.Path)
	if err != nil {
		return nil, err
	}

	var users []User

	for i := r.StartRow; true; i++ {
		name := f.GetCellValue(r.SheetName, r.NameCol+strconv.Itoa(i))

		if name == "" {
			break
		}

		code := f.GetCellValue(r.SheetName, r.CodeCol+strconv.Itoa(i))
		phone := f.GetCellValue(r.SheetName, r.PhoneCol+strconv.Itoa(i))

		user := User{
			Name:  name,
			Phone: phone,
			Code:  code,
		}

		users = append(users, user)
	}

	return users, nil
}
