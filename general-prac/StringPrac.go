package main

import (
	. "fmt"
	"net/url"
	"strings"
)

type a struct {
	val int64
}

type b struct {
	aList []*a
}

func main() {
	desc := "Sample abb dhfj "

	i := strings.LastIndex(desc," ")



	Println(i)

	desc = desc[:i] + "..."

	Println(desc)

	description := ".k;he;hf;eahf;hefehafihdifbafbqbf f;eshf;oaehrf;oahefoqehfoehfoeq sef;h;hf;ehf;oasehf;ohfohfoihf sfe;shf;shf;oshefosehf adhefh;hfe dsfh;sdhf;oshf;oshf sfdhs;odfh;oshfosfh fhsdhfsidhfsHfsaihfiahfeihf asfjoshfosfh sfdhf;shf;sfh"

	Println(len(description))
    getRefinedShortDescription("",description)

	Println("Desc ",description)


	Println("PRINING X")
	x := b{}
	//y := a{
	//	val: 21,
	//}
	//x.aList = append(x.aList,&y)
	//Println(x)

	x.aList = nil
	Println(x.aList)
	Println(len(x.aList))

	path := url.URL{Scheme: "https", Host: "cmwnext", Path: Sprintf("/detail/%s/%s", "series", "abc123")}


	Println(path)
	Println(path.String())

	encUrl := "https://cmwnext/detail/value:%22series%22%20/value:%22abc123%22%20"
	dV, e := url.QueryUnescape(encUrl)

	if e != nil {
		Println(e)
	}
	Println(dV)
}

func getRefinedShortDescription(shortDescription string, description string) {
	if shortDescription == "" && len(description) > 150 {
		// 0. Truncate the description
		shortDescription = description[:147]

		lastIndexOfSpace := strings.LastIndex(shortDescription, " ")
		if lastIndexOfSpace > -1 {
			// 1. Find last char of space and add 3 ellipses dots.
			shortDescription = shortDescription[:lastIndexOfSpace] + "..."
		}
	}
}


//func populateShortDescription*