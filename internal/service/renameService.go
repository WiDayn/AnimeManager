package service

import (
	"AnimeManager/internal/config"
	"github.com/sirupsen/logrus"
	"os"
	"regexp"
	"strings"
)

func Rename() {
	if config.FlagConfig.BatchModel {
		dir, err := os.ReadDir(config.FlagConfig.BasicDir)
		if err != nil {
			logrus.Panic(err)
		}
		for _, file := range dir {
			if file.IsDir() {
				renameAnime(config.FlagConfig.BasicDir + "\\" + file.Name())
			}
		}
	} else {
		renameAnime(config.FlagConfig.BasicDir)
	}
}

func renameAnime(basicDir string) {
	if basicDir == "" {
		logrus.Panic("BasicDir must be set.")
	}
	file, err := os.Open(basicDir)
	if err != nil {
		logrus.Panic("Read basicDir fail: create DIR struct error")
	}
	fileInfo, err := file.Stat()
	if err != nil {
		logrus.Panic("Read basicDir fail: read stat error")
	}
	if !fileInfo.IsDir() {
		logrus.Panic("Read basicDir fail: is not a dir")
	}
	var frontDir string
	var nowDirName string
	for i, text := range strings.Split(basicDir, "\\") {
		if i != len(strings.Split(basicDir, "\\"))-1 {
			frontDir += text + "\\"
		} else {
			nowDirName = text
		}
	}

	dirNewName := regexpName(config.AnimeLoadConfig.Detail.AnimeNameExclusion, config.AnimeLoadConfig.Detail.AnimeName, []byte(nowDirName))

	err = os.Rename(basicDir, frontDir+dirNewName)
	if err != nil {
		logrus.Panic("Write basicDir fail: rename error")
	}

	basicDir = frontDir + dirNewName
	DirFiles, err := os.ReadDir(basicDir)
	os.Mkdir(basicDir+"\\"+config.AnimeLoadConfig.Detail.DefaultSeason, 0775)
	for _, file := range DirFiles {
		if file.IsDir() {
			continue
		}
		if err != nil {
			logrus.Panic("Read basicDir fail: read stat error")
		}
		categoryID := regexpName(config.AnimeLoadConfig.Detail.CategoryExclusion, config.AnimeLoadConfig.Detail.Category, []byte(file.Name()))
		newName := basicDir + "\\" + config.AnimeLoadConfig.Detail.DefaultSeason + "\\" + config.AnimeLoadConfig.Detail.DefaultSeason + "E" + categoryID + " - " + dirNewName + "." + strings.Split(file.Name(), ".")[len(strings.Split(file.Name(), "."))-1]
		err = os.Rename(basicDir+"\\"+file.Name(), newName)
		if err != nil {
			logrus.Panic("Write file fail: rename error")
		}
	}
}

func regexpName(excludeList []string, useString string, name []byte) string {
	for _, exclude := range excludeList {
		name = regexp.MustCompile(exclude).ReplaceAll(name, []byte(""))
	}
	name = []byte(regexp.MustCompile(useString).FindAllString(string(name), -1)[0])
	return string(name)
}
