package main

import (
	"os"
	"os/exec"
)

const (
	TEMPLATEURL = "https://github.com/kidandcat/FeathersGO.git"
	GOCRAFT     = "github.com/gocraft/web"
	TOML        = "github.com/BurntSushi/toml"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case TRANS["GENERATE"]:
			generate()
		case TRANS["NEW"]:
			new()
		case TRANS["HELP"]:
			help()
		default:
			cmdNotFound()
		}
	} else {
		println(TRANS["MISSARGUMENT"])
	}
}

func cmdNotFound() {
	println(TRANS["CMD"], os.Args[1], TRANS["NOTFOUND"])
}

func help() {
	println(TRANS["HELPTEXT"])
}

func checkGit() bool {
	if out, _ := exec.Command("git").CombinedOutput(); len(out) == 0 {
		return false
	}
	return true
}

func generate() {
	println(TRANS["GENERATING"])
}

func new() {
	if len(os.Args) < 3 {
		println(TRANS["APPNAME"])
		return
	}
	if checkGit() == false {
		println(TRANS["GIT"])
	} else {
		println(TRANS["DOWNLOAD"])
		out, err := exec.Command("git", "clone", TEMPLATEURL, os.Args[2]).CombinedOutput()
		println(string(out))
		if err != nil {
			println(TRANS["ERRORGIT"])
			return
		}

		if loadDependencies() == false {
			println(TRANS["ERRORGOGET"])
		}
		println(TRANS["DONE"])

	}
}

func loadDependencies() bool {
	if loadPackage(GOCRAFT) == false {
		return false
	}
	if loadPackage(TOML) == false {
		return false
	}
	return true
}

func loadPackage(pkg string) bool {
	out, err = exec.Command("go", "get", pkg).CombinedOutput()
	println(string(out))
	if err != nil {
		return false
	}
	return true
}
