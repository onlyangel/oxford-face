package face_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	. "github.com/kkdai/oxford-face"
)

func init() {
	API_KEY = os.Getenv("MSFT_KEY")
	if API_KEY == "" {
		fmt.Println("Please export your key to environment first, `export MSFT_KEY=12234`")
	}
}

func TestFaceListCreation(t *testing.T) {
	if API_KEY == "" {
		return
	}

	fList := NewFaceList(API_KEY)
	ret, _ := fList.Create("id001", "testlist1", "this is a test list")
	log.Println("create ret:", string(ret))

}

func TestFaceListGet(t *testing.T) {
	if API_KEY == "" {
		return
	}

	fList := NewFaceList(API_KEY)
	retList, err := fList.List()
	if err != nil {
		t.Error("Error on get:" + err.Err.Error())
	}

	log.Println("Get list ret:", string(retList))
}

func TestFaceListGetByID(t *testing.T) {
	if API_KEY == "" {
		return
	}

	fList := NewFaceList(API_KEY)
	retList, err := fList.Get("id002")
	if err == nil {
		t.Error("Error on get:" + err.Err.Error())
	}

	log.Println("Get list id002 ret:", string(retList))
}
