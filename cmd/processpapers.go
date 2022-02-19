package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/perfect6566/bingpaper/db"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Fetchpapers(paperapiurl string){
	client := &http.Client{}
	bingres:= db.Bingresp{}
	papers:= db.Papers{}
	defer func() {
		if err:=recover();err!=nil{
			log.Println(err)
		}
	}()

	resp, err := client.Get(paperapiurl)
	if err!=nil{
		log.Print(err)
	}
	defer resp.Body.Close()
	if(resp.StatusCode==http.StatusOK) {
		res,err:=ioutil.ReadAll(resp.Body)
		if err!=nil{
			log.Println(err)
		}
	json.Unmarshal(res,&bingres)

	imageresult:=bingres.Images[0]
	papers.Image=imageresult
	papers.Image.URL="https://cn.bing.com"+papers.Image.URL
	papers.Createtime=time.Now()
	err=downloadimage(papers)
	if err!=nil{
		log.Println(err)
		return
	}
	db.Insert(papers)
	}
}

func downloadimage(papers db.Papers) (error){
	defer func() {
		if err:=recover();err!=nil{
			log.Println(err)
		}
	}()
	client := &http.Client{}
	resp,err:=client.Get(papers.URL)
	if err!=nil{
		log.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode==http.StatusOK {
		data, err := ioutil.ReadAll(resp.Body)
		if err!=nil{
			log.Println(err)
		}
		imagename := papers.Title
		enddate := papers.Enddate
		err = ioutil.WriteFile("papers/"+enddate+imagename+".jpg", data, 0640)
		log.Println(fmt.Sprintf("Successfully download image %s.",enddate+imagename+".jpg"))
	}else {
		log.Println(resp.Status)
	}
	return err
}
