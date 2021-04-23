package main

import (
	"fmt"
	"reflect"
)

type Image struct {
	H int
	W int
	url string
}

type topLevel struct {
	image Image
	a string
	b string
	images []Image
}

type MapImage interface {
	modify()
}

func (tl *topLevel) modify() {
	tl.image.H = tl.image.H + 1
}

func main(){
	//const (
	//	layoutISO = "2006-01-02"
	//	layoutUS  = "2006-01-02T15:04:05Z"
	//)
	//date := "2007-01-02"
	//t, e := time.Parse(layoutISO, date)
	//if e != nil {
	//	fmt.Println("Err occurred")
	//}
	//fmt.Println(t)                  // 1999-12-31 00:00:00 +0000 UTC
	//fmt.Println(t.Format(layoutUS)) // December 31, 1999
	//
	//fmt.Println(time.Date(2007,1,2,0,0,0,0,time.UTC).Format(layoutUS))
	//
	//fmt.Println(time.Time{}.Format(layoutUS))



	//tl := topLevel{
	//	image: Image{
	//		H: 123,
	//		W:12,
	//		url: "sample",
	//	},
	//}
	//
	//tl2 := topLevel{
	//	image: Image{
	//		H: 1232,
	//		W:12323,
	//		url: "sample",
	//	},
	//}
	//
	//
	//
	//sTl := []topLevel{tl,tl2}
	//
	//modifiedList := make([]topLevel,0)
	//
	//for _, v := range sTl {
	//	v.modify()
	//	modifiedList = append(modifiedList, v)
	//}
	//
	//sTl = modifiedList
	//
	//for _,v := range sTl {
	//	fmt.Println(v.image.H)
	//}


	//testGen(5)
	//
	//testGen("Mayank")
	//
	//testGen(topLevel{})
	//
	//testGen(Image{})

	top := topLevel{}
	top.images = abc()

	fmt.Println(top.images)
}

func abc() []Image {
	return nil
}


func testGen(abc interface{}) {
	fmt.Println(abc)

	value := reflect.TypeOf(abc)

	fmt.Println(value)


}