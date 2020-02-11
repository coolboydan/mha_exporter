package command

import (
	"bufio"
	"golang.org/x/tools/go/ssa/interp/testdata/src/strings"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
)

var (
	// The path of the proc filesystem.
	procPath = kingpin.Flag("path.procfs", "procfs mountpoint.").Default("/Users/zhengjunbo/Downloads/").String()

	deadStr     = "Dead Servers"
	activeStr   = "Alive Servers"
	aliveSlaves = "Alive Slaves"
	aliveMaster = "Current Alive Master"
)

type ServerStat struct {
	deadServer   map[string]float64
	aliveServer  map[string]float64
	statusServer map[string]float64
}

func procFilePath(name string) string {
	return filepath.Join(*procPath, name)
}

func readMHAStatus(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	var stat ServerStat
	stat.aliveServer = make(map[string]float64)
	stat.deadServer = make(map[string]float64)
	stat.statusServer = make(map[string]float64)

	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, deadStr) {

		}
		stat.aliveServer[""] = 1
		println(line)
	}
}
