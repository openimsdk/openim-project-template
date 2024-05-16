//go:build mage
// +build mage

package main

import (
	"github.com/openimsdk/gomake/mageutil"
	"github.com/openimsdk/openim-project-template/pkg/util"
	"os"
)

var Default = Build

func Build() {
	mageutil.Build()
}

func Start() {
	mageutil.InitForSSC()
	err := setMaxOpenFiles()
	if err != nil {
		mageutil.PrintRed("setMaxOpenFiles failed " + err.Error())
		os.Exit(1)
	}
	mageutil.StartToolsAndServices()
}

func Stop() {
	mageutil.StopAndCheckBinaries()
}

func Check() {
	mageutil.CheckAndReportBinariesStatus()
}

func Protocol() {
	util.Protocol()
}
