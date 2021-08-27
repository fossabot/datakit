package cmds

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	dl "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/downloader"
)

const (
	DIR_NAME = "telegraf"
)

func InstallTelegraf(installDir string) error {
	url := "https://zhuyun-static-files-production.oss-cn-hangzhou.aliyuncs.com/datakit/telegraf/" + fmt.Sprintf("telegraf-%s_%s.tar.gz", runtime.GOOS, runtime.GOARCH)

	if runtime.GOOS != "windows" {
		installDir = "/"
	}

	fmt.Printf("Start downloading Telegraf...\n")
	dl.CurDownloading = "telegraf"

	cli := getcli()

	if err := dl.Download(cli, url, installDir, false, false); err != nil {
		return err
	}

	if err := writeTelegrafSample(installDir); err != nil {
		return err
	}

	fmt.Printf("Install Telegraf successfully!\n")
	if runtime.GOOS == "windows" {
		fmt.Printf("Start telegraf by `cd %v`, `copy telegraf.conf.sample tg.conf`, and `telegraf.exe --config tg.conf`\n", filepath.Join(installDir, DIR_NAME))
	} else {
		fmt.Println("Start telegraf by `cd /etc/telegraf`, `cp telegraf.conf.sample tg.conf`, and `telegraf --config tg.conf`\n", filepath.Join(installDir, DIR_NAME))
	}

	fmt.Printf("Vist https://www.influxdata.com/time-series-platform/telegraf/ for more information.\n")

	return nil
}

func writeTelegrafSample(installDir string) error {
	var filePath string
	if runtime.GOOS != "windows" {
		filePath = "/etc/telegraf/telegraf.conf.sample"
	} else {
		filePath = filepath.Join(installDir, DIR_NAME, "telegraf.conf.sample")
	}

	return ioutil.WriteFile(filePath, []byte(TelegrafConfTemplate), os.ModePerm)
}
