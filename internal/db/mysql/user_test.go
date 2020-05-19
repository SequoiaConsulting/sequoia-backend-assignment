package mysql

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func TestFoo(t *testing.T) {
	db, err := gorm.Open("mysql", "dummy:password/thenga?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Error("failed initializeing stiff")
	}

	ur := UserMySQLRepository{
		db,
	}

	_, err = ur.GetByID("1")
	if err != nil {

		t.Error(err.Error())
	}
}
