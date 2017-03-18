package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func main() {

	source := flag.String("source", "./", "source directory")
	out := flag.String("out", "./", "output directory")
	level := flag.Int("level", 0, "compression level")
	flag.Parse()

	sourceDir := strings.TrimSuffix(*source, "/") + "/"
	outDir := strings.TrimSuffix(*out, "/") + "/"
	createIfMissing(outDir)

	exclude := flag.Args()

	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

FILE_ITERATOR:
	for _, f := range files {

		if !isFile(sourceDir + f.Name()) {
			continue FILE_ITERATOR
		}

		for _, exc := range exclude {
			if strings.Contains(f.Name(), exc) {
				continue FILE_ITERATOR
			}
		}

		if strings.Contains(f.Name(), "gzip") {
			continue FILE_ITERATOR
		}

		j := job{
			sourceDir: sourceDir,
			outDir:    outDir,
			fileName:  f.Name(),
			level:     *level,
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			execute(j)
		}()
	}

	wg.Wait()

}

type job struct {
	sourceDir string
	outDir    string
	fileName  string
	level     int
}

func execute(j job) {
	content := readFile(j.sourceDir + j.fileName)
	writeFile(content, j.fileName, j.outDir, j.level)
}

func writeFile(content []byte, filename string, out string, level int) {
	if len(content) == 0 {
		return
	}

	gz := gzipThis(content, level)

	err := ioutil.WriteFile(out+filename, gz, 0644)
	if err != nil {
		panic(err)
	}
}

func readFile(name string) []byte {
	buf := bytes.NewBuffer(nil)
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(buf, file)
	if err != nil {
		panic(err)
	}
	file.Close()
	return buf.Bytes()
}

func gzipThis(data []byte, level int) []byte {

	switch level {
	case 0:
		return data
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
	case -1:
	case -2:
	default:
		return data
	}

	var b bytes.Buffer
	gz, err := gzip.NewWriterLevel(&b, level)
	if err != nil {
		panic(err)
	}

	if _, err := gz.Write(data); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	return b.Bytes()
}

const newLine = `
`

func isFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return !fileInfo.IsDir()
}

func createIfMissing(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}
