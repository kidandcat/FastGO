package main

var TRANS = map[string]string{
	"APPNAME":      "Provide a name for your new Application",
	"GENERATE":     "generate",
	"GENERATING":   "Generating something",
	"GIT":          "You need to have git in your path",
	"DOWNLOAD":     "Downloading template",
	"ERRORGIT":     "Error with git",
	"ERRORGOGET":   "Error with go get",
	"DONE":         "Done",
	"HELP":         "help",
	"CMD":          "Command",
	"NOTFOUND":     "not found",
	"MISSARGUMENT": "You didn't pass any argument",
	"NEW":          "new",
	"HELPTEXT": `
 This is the FeathersGo multipurpose console,
 this tool is used for creating a new app, but also
 to manage your existing app.

 Available commands:
   new        create new app
   help       show this text
   generate   generate new things
	`,
}
