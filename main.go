package main

import (
	"github.com/gernest/alien"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func check(w http.ResponseWriter, r *http.Request) {
	os.Chdir("C:\\Users\\user\\Desktop\\capt")
	files, _ := filepath.Glob("*")
	files_str := strings.Join(files, "\n")
	w.Write([]byte(files_str))
}

func index(w http.ResponseWriter, r *http.Request) {
	os.Chdir("C:\\Users\\user\\Desktop\\capt\\")
	w.Write([]byte(
		` __                        
|__)  _  _  |_  _   _   _  
|    |  (_| |_ (_| ||| (_| 
                           
 __      _                  
(_   _  (_ |_      _   _  _ 
__) (_) |  |_ \)/ (_| |  (- 
                            


Server apotek berjalan! Route & fungsi:

1. /cctv_on/ 		-> 	Menyalakan CCTV
2. /cctv_off/ 		-> 	Matikan CCTV
3. /check/ 		-> 	Check isi folder CCTV
4. /shutdown/ 		-> 	Matikan komputer
5. /content/:path 	-> 	Lihat isi folder di folder CCTV
6. /ffmpeg/:path 	-> 	Encoding

`))
}

func cctv_on(w http.ResponseWriter, r *http.Request) {
	err := os.Chdir("C:\\Users\\user\\Desktop")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("nircmd", "exec", "hide", "cctv.bat")
	cmd.Start()
	w.Write([]byte(
		` __  __ ___      
/   /    |  \  / 
\__ \__  |   \/  
                 
 __        
/  \  _  | 
\__/ | ) . 
           

`))
}

func cctv_off(w http.ResponseWriter, r *http.Request) {
	cmd, err := exec.Command("taskkill", "/f", "/im", "python.exe").Output()
	if err != nil {
		log.Fatal(err)
	}
	w.Write(cmd)
}

func mati(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("shutdown", "-s", "-t", "05")
	w.Write([]byte("Matikan windows"))
	cmd.Start()
}

func checkf(w http.ResponseWriter, r *http.Request) {
	p := alien.GetParams(r)
	os.Chdir("C:\\Users\\user\\Desktop\\capt\\" + p.Get("path"))
	files, _ := filepath.Glob("*")
	files_str := strings.Join(files, "\n")
	w.Write([]byte(files_str))
}

func ffmpeg(w http.ResponseWriter, r *http.Request) {
	p := alien.GetParams(r)
	os.Chdir("C:\\Users\\user\\Desktop\\capt\\" + p.Get("path"))
	cmd := exec.Command("nircmd", "exec", "hide", "ffmpeg", "-r", "2", "-i", "img%05d.jpg", "-c:v", "libx264", "-vf", "fps=25", "-pix_fmt", "yuv420p", "out.mp4")
	cmd.Start()
	w.Write([]byte("Encoding this folder!\n"))
}

func main() {
	m := alien.New()
	m.Get("/", index)
	m.Get("/content/:path", checkf)
	m.Get("/ffmpeg/:path", ffmpeg)
	m.Get("/shutdown/", mati)
	m.Get("/check/", check)
	m.Get("/cctv_on/", cctv_on)
	m.Get("/cctv_off/", cctv_off)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", m))
}
