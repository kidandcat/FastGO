package main

import (
	"os"
	"os/exec"
)

const (
	TEMPLATEURL = "https://github.com/kidandcat/feathersgo-template"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case TRANS["GENERATE"]:
			println(TRANS["GENERATING"])
		case TRANS["NEW"]:
			if len(os.Args) < 3 {
				println(TRANS["APP_NAME"])
				return
			}
			if checkGit() {
				println(TRANS["GIT"])
			} else {
				println(TRANS["DOWNLOAD"])
				out, err := exec.Command("git", "clone", TEMPLATEURL, os.Args[2]).CombinedOutput()
				println(string(out))
				if err != nil {
					println(TRANS["ERROR"])
					return
				}
				println(TRANS["DONE"])
			}
		case TRANS["HELP"]:
			println(help())
		default:
			println(TRANS["CMD"], os.Args[1], TRANS["NOTFOUND"])
		}
	} else {
		println(TRANS["MISSARGUMENT"])
	}
}

func help() string {
	return TRANS["HELPTEXT"]
}

func checkGit() bool {
	if _, err := exec.Command("git").Output(); err != nil {
		return false
	}
	return true
}
