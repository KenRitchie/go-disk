package controller

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"go-disk/back/config"
)

var T *template.Template
/**
 * 上传
 */
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		T, _ = template.ParseFiles(config.TemplateDir + "/index.html")
		err := T.Execute(w, nil)
		checkHttpErr(err, w)
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		checkErr(err)
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(config.UploadDir + "/" + filename)
		checkErr(err)
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			checkErr(err)
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}

}

/**
 * 浏览
 */
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := config.UploadDir + "/" + imageId
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}


/**
 * 入口文件
 */
func IndexHandler(w http.ResponseWriter, r *http.Request) {


	//file,err:=os.OpenFile("./env",os.O_WRONLY,0600)
	//if err !=nil {
	//	file1,err:=os.Create("./env")
	//	if err == nil{
	//		n,err:=file1.Write([]byte("{'uploadDir':'./temp'}"))
	//		fmt.Println(n,err)
	//	}
	//}else{
	//	b,err:=ioutil.ReadAll(file)
	//	fmt.Println(b,err)
	//}
	//
	//
	//_, err = ioutil.ReadDir("./uploads")
	//if err != nil {
	//	dirName,err:=ioutil.TempDir("./","temp")
	//	fmt.Println(err)
	//	_, _ = ioutil.ReadDir(dirName)
	//
	//}

	if r.Method == "GET" {
		T,err:=template.ParseFiles(config.TemplateDir + "/index.html")
		if err !=nil{
			log.Println("ParseFiles fail",err)
		}
		err=T.Execute(w,nil)
		checkHttpErr(err,w)
		return
	}


	//方式一
	//	var listHtml string
	//	for _, fileInfo := range fileInfoArr {
	//		imgid := fileInfo.Name()
	//		listHtml += "<li><a href=\"/view?id=" + imgid + "\">imgid</a></li>"
	//	}
	//	io.WriteString(w, "<div>"+listHtml+"</div>")
	//方式二
	//var locals = make(map[string]interface{})
	//var images = []string{}
	//for _, fileInfo := range fileInfoArr {
	//	images = append(images, fileInfo.Name())
	//}
	//locals["images"] = images
	//t, err := template.ParseFiles(config.TemplateDir + "/index.html")
	//checkHttpErr(err, w)
	//err = t.Execute(w, locals)
	//checkHttpErr(err, w)
}
