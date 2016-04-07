package tools

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

//Ti2s function is used for interface{} to string
func Ti2s(i interface{}) (ret string) {
	b, ok := i.([]byte)
	if ok {
		ret = string(b)
	} else {
		ret = "Error Parse interface to string in itos"
	}
	return
}

//Ti2int function conver interface{} to integer
func Ti2int(i interface{}) (ret int) {
	b, ok := i.(int)
	if ok {
		ret = int(b)
	} else {
		ret, _ = strconv.Atoi(Ti2s(i))
	}
	return
}

//Tfile2bytes is rapid read file to []byte quickly bu using ioutil
func Tfile2bytes(filepath string) (ret []byte, err error) {
	ret = nil
	err = nil
	if _, err = os.Stat(filepath); err != nil {
		return
	}
	ret, err = ioutil.ReadFile(filepath)
	return
}

//TimportDisqus import Disqus's comment to ghost database
func TimportDisqus(xmlfile string) {
	buff, err := Tfile2bytes(xmlfile)
	if err != nil {
		panic(err)
	}

	type uid struct {
		Val string `xml:"dsqid,attr"`
	}
	type Post struct {
		UID         string    `xml:"dsqid,attr"`
		Message     string    `xml:"message"`
		CreatedAt   time.Time `xml:"createdAt"`
		AuthorEmail string    `xml:"author>email"`
		AuthorName  string    `xml:"author>name"`
		AuthorNick  string    `xml:"author>username"`
		IP          string    `xml:"ipAddress"`
		Tid         uid       `xml:"thread"`
		Pid         uid       `xml:"parent"`
	}

	type Thread struct {
		UID         string    `xml:"dsqid,attr"`
		Forum       string    `xml:"forum"`
		Link        string    `xml:"link"`
		Title       string    `xml:"title"`
		Message     string    `xml:"message"`
		CreateAt    time.Time `xml:"createdAt"`
		AuthorNmae  string    `xml:"author>name"`
		AuthorEmail string    `xml:"author>email"`
		Anonymous   bool      `xml:"author>isAnonymous"`
		IP          string    `xml:"ipAddress"`
		Closed      bool      `xml:"isClosed"`
		Deleted     bool      `xml:"isDeleted"`
	}

	type Category struct {
		Cid   string `xml:"dsqid,attr"`
		Forum string `xml:"forum"`
		Title string `xml:"title"`
	}

	type Disqus struct {
		Category []Category `xml:"category"`
		Threads  []Thread `xml:"thread"`
		Posts    []Post   `xml:"post"`
	}

	v := Disqus{}

	err = xml.Unmarshal(buff, &v)

	if err != nil {
		panic(err)
	}
	bs, _ := json.Marshal(v)
	ioutil.WriteFile(xmlfile+".json", bs, 0777)
}
