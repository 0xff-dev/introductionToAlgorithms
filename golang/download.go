package golang

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const downloadCompleteTemplate = "Downloading  %s..."

var fileSizeUnit = []string{"B", "KB", "MB", "GB", "TB"}

type Downloader struct {
	Size int
}

func (d *Downloader) Write(p []byte) (int, error) {
	n := len(p)
	d.Size += n
	fmt.Printf("\r%s", fmt.Sprintf(downloadCompleteTemplate, getFileSize(float64(d.Size))))
	return n, nil
}

func getFileSize(total float64) string {
	unitIdx := 0
	for total > 1024 && unitIdx < len(fileSizeUnit) {
		unitIdx++
		total /= 1024.0
	}
	return fmt.Sprintf("%.2f%s", total, fileSizeUnit[unitIdx])
}

func download(path, remoteFile, downloadFileName string) {
	log.Println("start downloading")
	writeFile, err := os.Create(path + "/" + downloadFileName)
	if err != nil {
		log.Fatalf("create file to [%s] error: %s\n", path+"/"+downloadFileName, err)
		return
	}
	defer writeFile.Close()

	resp, err := http.Get(remoteFile)
	if err != nil {
		log.Fatalf("get %s error: %s", remoteFile, err)
		return
	}
	defer resp.Body.Close()

	downloader := &Downloader{}
	_, err = io.Copy(writeFile, io.TeeReader(resp.Body, downloader))
	if err != nil {
		log.Fatalf("copy file error: %s", err)
		return
	}
	fmt.Println()
	log.Printf("download complete!\n")
}

var path = flag.String("path", ".", "file download path")
var remoteFile = flag.String("remote", "", "remote file url")

func start() {
	flag.Parse()

	if !(strings.HasPrefix(*remoteFile, "https") || strings.HasPrefix(*remoteFile, "http")) {
		log.Fatal("only support http and https")
		return
	}

	writtenFile := strings.Split(*remoteFile, "/")
	download(*path, *remoteFile, writtenFile[len(writtenFile)-1])
}
