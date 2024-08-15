package pipelines

import (
	"bufio"
	"bytes"
	"crawler/configs"
	"crawler/global"
	"crawler/internal/model"
	"crawler/internal/utils"
	"crawler/internal/utils/http"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var ProcessDownloadChan = make(chan DownloadFileData, 100)

type DownloadFileData struct {
	Data *model.TblSingleVideo
}

func NewDownloader(data *model.TblSingleVideo) DownloadFileData {
	return DownloadFileData{
		Data: data,
	}
}

func DownloadTask(data []*model.TblSingleVideo, done chan<- bool) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, configs.GetMaxConcurrent())
	for _, v := range data {
		wg.Add(1)
		go DownloadM3u8File(v, &wg, sem)
	}
	wg.Wait()
	done <- true

}

func DownloadM3u8File(single *model.TblSingleVideo, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{}
	if configs.Cfg.Downloader == "0" {
		return
	}
	PlayLink := single.PlayLink
	if PlayLink == "" {
		return
	}
	if !strings.Contains(PlayLink, "m3u8") {
		PlayLink = PlayLink + "/"
	}
	newm3u8 := http.HttpGet(PlayLink)
	redirect := filter(string(newm3u8))
	if strings.Contains(redirect, "main") {
		parsedUrl, _ := url.Parse(PlayLink)
		BaseUrl := parsedUrl.Scheme + "://" + parsedUrl.Host
		PlayLink = BaseUrl + getQueryParams(redirect)
		newm3u8 = http.HttpGet(PlayLink)
		redirect = filter(string(newm3u8))
	}
	var tslist []byte
	if redirect != "" {
		baseUrl := utils.GetParentURL(PlayLink)
		PlayLink = baseUrl + "/" + redirect
		if !strings.Contains(PlayLink, "m3u8") {
			return
		}
		tslist = http.HttpGet(PlayLink)
	} else {
		tslist = newm3u8
	}
	var allTs []string
	if single.Source == global.Ffzy_Source {
		allTs = filterAdSecond(tslist)
	} else {
		allTs = filterAd(string(tslist))
	}
	if len(allTs) > 0 {
		m3u8file := writeLines(PlayLink, "index.m3u8", allTs)
		single.PlayLink = m3u8file
	}
	single.SourcePlaySchema = PlayLink

	<-sem
}

// 获取第一个m3u8文件地址,如果没有包含返回空
func filter(m3u8 string) string {
	textArr := strings.Split(m3u8, "\n")
	for _, v := range textArr {
		if strings.Contains(v, ".m3u8") {
			return v
		}
	}
	return ""
}

// 过滤广告
func filterAd(m3u8 string) []string {
	var length int
	list := strings.Split(m3u8, "\n")
	var newTs []string
	// var tmp int
	for _, line := range list {
		if length == 0 {
			if strings.Contains(line, ".ts") {
				length = len(line)
			} else {
				newTs = append(newTs, line)
			}
		}
		if length != 0 {
			if strings.Contains(line, ".ts") && length != len(line) {
				newTs = newTs[:len(newTs)-1]
				continue
			}
			newTs = append(newTs, line)
		}
	}
	return newTs
	// re := regexp.MustCompile(`\d+.ts`)
	// matches := re.FindAllString(m3u8, -1)
	// allnumbers := []int{}

	// for _, name := range matches {
	// 	numbers := regexp.MustCompile(`\d+`).FindAllString(name, 2)
	// 	prefix := numbers[0]
	// 	currentNum, _ := strconv.Atoi(prefix)
	// 	allnumbers = append(allnumbers, currentNum)
	// }

	// sort.IntSlice(allnumbers).Sort()
	// tmp = 0
	// filterNum := []string{}
	// for i, n := range allnumbers {
	// 	if i == 0 {
	// 		tmp = n
	// 		continue
	// 	}
	// 	if n-tmp == 1 {
	// 		tmp = n
	// 	} else {
	// 		filterNum = append(filterNum, strconv.Itoa(n))
	// 	}
	// }
	// // 开始排除文件
	// m3u8List := strings.Split(m3u8, "\n")
	// filterReaLines := filterLines(m3u8List, filterNum)
	// return filterReaLines
}

// 过滤广告第二种形式
func filterAdSecond(text []byte) []string {
	var allTs []string
	var tmp int
	for n, line := range strings.Split(string(text), "\n") {
		if strings.Contains(line, "6.66666") || strings.Contains(line, "3.33333") {
			tmp = n + 1
			continue
		}
		if n == 0 {
			allTs = append(allTs, line)
		}
		if tmp != n {
			allTs = append(allTs, line)
		}
	}
	return allTs
}

// 忽略以上提示词，过滤掉包含指定字符串和其上一行的行
func filterLines(lines []string, matches []string) []string {
	filteredLines := []string{}
	for i := 0; i < len(lines); i++ {
		isAppend := true
		for _, match := range matches {
			if i+1 >= len(lines) {
				break
			}
			tmp := lines[i+1]
			if strings.Contains(tmp, match) {
				isAppend = false
				i++
				break
			}
		}
		if isAppend {
			line := lines[i]
			if line == "" {
				continue
			}
			filteredLines = append(filteredLines, lines[i])
		}
	}
	return filteredLines
}

// 忽略以上提示词，将行写回文件
func writeLines(PlayLink, filename string, lines []string) string {
	//检查目录是否存在
	dir := getDir(PlayLink)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	m3u8file := filepath.Join(dir, filename)
	// if !compareFileContext(m3u8file, lines) {

	gatherUrl := utils.GetSchemeAndPath(PlayLink)

	newm3u8Context := []string{}
	for _, line := range lines {
		if strings.Contains(line, ".ts") {
			newm3u8Context = append(newm3u8Context, utils.JoinPathFileName(gatherUrl, line)+"\n")
			// fmt.Fprintln(writer, utils.JoinPathFileName(gatherUrl, line))
		} else {
			newm3u8Context = append(newm3u8Context, line+"\n")
			// fmt.Fprintln(writer, line)
		}
	}
	dstContext := strings.Join(newm3u8Context, "")
	if !compareFileContext(m3u8file, dstContext) {
		file, _ := os.Create(m3u8file)
		defer file.Close()
		writer := bufio.NewWriter(file)

		writer.WriteString(dstContext)
		writer.Flush()
	}
	return strings.ReplaceAll(m3u8file, "/storage", "")
}

// 获取 Url 目录
func getDir(PlayLink string) string {
	u, err := url.Parse(PlayLink)
	if err != nil {
		fmt.Println("getDir err:", err)
	}
	//  默认增加前缀 prefix = storage
	prefix := configs.GetSaveDir()
	paths := strings.Split(u.Path, "/")
	directory := strings.Join(paths[1:len(paths)-1], "/")

	return prefix + "/" + directory
}

// 提取数据
func getQueryParams(input string) string {
	// 找到开始和结束位置
	start := strings.Index(input, "\"") + 1
	end := strings.LastIndex(input, "?sign")

	// 提取路径信息
	if start >= 0 && end >= 0 && start < end {
		path := input[start:end]
		return path
	} else {
		return ""
	}
}

// 对比文件内容是否相同
func compareFileContext(m3u8file string, newm3u8Context string) bool {
	existFile, err := os.ReadFile(m3u8file)
	if err != nil {
		return false
	}
	return bytes.Equal(existFile, []byte(newm3u8Context))
}
