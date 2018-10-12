package faithtop

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var cacheDir string

func GetCacheDir() string {
	if cacheDir != "" {
		return cacheDir
	}
	sep := string(os.PathSeparator)
	homeDir := GetHomeDir()
	if runtime.GOOS == "linux" {
		cacheDir = Getrpath(homeDir) + ".cache/faithtop/"
		return cacheDir
	} else if runtime.GOOS == "windows" {
		cacheDir = Getrpath(homeDir) + "AppData" + sep + "Local" + sep + "faithtop" + sep
		return cacheDir
	} else {
		cacheDir = Getrpath(GetCurrentExecPath()) + ".cache" + sep
		return cacheDir
	}
}
func SetCacheDir(dir string) {
	cacheDir = Getrpath(dir)
}
func GetCurrentExecPath() string {
	f, e := exec.LookPath(os.Args[0])
	if e != nil {
		fmt.Println("exec.LookPath() failed:", e)
		return ""
	}
	realPath, e := filepath.Abs(f)
	if e != nil {
		fmt.Println("filepath.Abs(f) failed", e)
		return ""
	}
	return realPath
}
func GetHomeDir() string {
	u, e := user.Current()
	if e != nil {
		return GetCurrentExecPath()
	}
	return u.HomeDir
}

func NewNumToken() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.FormatInt(int64(r.Intn(60000)), 10)
}
func NewToken() string {
	ct := time.Now().UnixNano()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(ct, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}
func JsonArray(i interface{}) string {
	b, e := json.Marshal(i)
	if e != nil {
		return "[]"
	}
	return string(b)
}
func JsonObject(i interface{}) string {
	b, e := json.Marshal(i)
	if e != nil {
		return "{}"
	}
	return string(b)
}
func UnJson(str string, v interface{}) {
	e := json.Unmarshal([]byte(str), v)
	if e != nil {
		fmt.Println("UnJson() err:", e)
	}
}
func Getrpath(s string) string {
	if len(s) < 1 {
		return ""
	}
	sp := string(os.PathSeparator)
	if s[len(s)-1:] == sp {
		return s
	}
	return s + sp
}
func DownloadNetFile(url, downloadDir string, callback func(string)) {
	f := Getrpath(downloadDir) + NewToken()
	e := DownloadFile(url, f)
	if e != nil {
		fmt.Println(e)
		return
	}
	callback(f)
}
func CacheNetFile(url, cacheDir string, callback func(string)) {
	f := Getrpath(cacheDir) + Url2cachePath(url)
	if _, e := os.Stat(f); e != nil {
		e = DownloadFile(url, f)
		if e != nil {
			fmt.Println(e)
			return
		}
	}
	callback(f)
}
func EndsWith(s, suffix string) bool {
	if len(suffix) > len(s) {
		return false
	}
	if s[len(s)-len(suffix):] == suffix {
		return true
	}
	return false
}
func StartsWidth(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	if s[:len(prefix)] == prefix {
		return true
	}
	return false
}
func Url2cachePath(url string) string {
	rUrl := GetRealUrl(url)
	s := strings.Replace(rUrl, "://", "/", -1)
	sep := string(os.PathSeparator)
	s = strings.Replace(s, "/", sep, -1)
	if EndsWith(s, sep) {
		return s[:len(s)-1]
	}
	return s
}
func GetRealUrl(url string) string {
	for i := 0; i < len(url); i++ {
		item := url[i : i+1]
		if item == "?" || item == "#" {
			return url[:i]
		}
	}
	return url
}
func SPrintf(a ...interface{}) string {
	return fmt.Sprint(a...)
}
func XPrintf(a int) string {
	return fmt.Sprintf("%x", a)
}
func a2i(a string) (int, error) {
	return strconv.Atoi(a)
}
func a2f(a string) float64 {
	f, e := strconv.ParseFloat(a, 32)
	if e != nil {
		fmt.Println(e)
		return 0
	}
	return float64(f)
}
func DownloadFile(url, fdist string) error {
	rp, e := http.Get(url)
	if e != nil {
		return e
	}
	defer rp.Body.Close()
	f, e := WriteFile(fdist)
	if e != nil {
		return e
	}
	defer f.Close()
	_, e = io.Copy(f, rp.Body)
	return e
}
func WriteFile(f string) (*os.File, error) {
	e := os.MkdirAll(GetDirOfFile(f), 0755)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}
	file, e := os.OpenFile(f, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	return file, e
}
func GetDirOfFile(f string) string {
	for i := len(f) - 1; i > -1; i-- {
		if f[i:i+1] == string(os.PathSeparator) {
			return f[:i]
		}
	}
	return f
}
