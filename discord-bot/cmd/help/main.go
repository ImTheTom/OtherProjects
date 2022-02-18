package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	"github.com/ImTheTom/OtherProjects/discord-bot/config"
	"github.com/sirupsen/logrus"
)

type tmpStruct struct {
	PackageName string
	HelpMessage string
}

var overalMap map[string]string

var (
	removeCommand = regexp.MustCompile("Command")
	removeComment = regexp.MustCompile("// ")
)

func main() {
	config.InitLogger()

	logrus.Info("Generating help command code...")

	packa := os.Args[1]

	templateLoc := os.Args[2]

	cwd, err := os.Getwd()
	if err != nil {
		logrus.Fatalf("Failed at get cwd with %v", err)
	}

	// Create the map which will hold the information
	overalMap = make(map[string]string)

	processDirectory(cwd)

	logrus.Infof("%v", overalMap)

	createFile(packa, templateLoc)
}

func processDirectory(cwd string) {
	logrus.Infof("Looking in %v", cwd)

	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		logrus.Fatalf("Failed to get cwd files with %v", err)
	}

	processFiles(files)
}

func processFiles(files []fs.FileInfo) {
	fset := token.NewFileSet()

	for _, file := range files {
		// Ignore some basic general cases
		if file.IsDir() {
			continue
		}

		if strings.HasPrefix(file.Name(), "gen_") {
			continue
		}

		logrus.Infof("Parsing %v", file.Name())

		// Find everything!
		f, err := parser.ParseFile(fset, file.Name(), nil, parser.ParseComments)
		if err != nil {
			logrus.Warnf("Parsing %v failed with %v", file.Name(), err)

			continue
		}

		for _, decl := range f.Decls {
			// Find declrations:
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			// Find constant declrations:
			if genDecl.Tok != token.CONST {
				continue
			}

			// Use constant declartions.
			processConstantDeclartions(genDecl.Specs)
		}
	}
}

func processConstantDeclartions(specs []ast.Spec) {
	for _, spec := range specs {
		var constName string

		var constComment string

		valSpec, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}

		for _, name := range valSpec.Names {
			if strings.HasSuffix(name.Name, "Command") {
				constName = removeCommand.ReplaceAllString(name.Name, "")
			}
		}

		if valSpec.Doc != nil {
			for _, comment := range valSpec.Doc.List {
				constComment = removeComment.ReplaceAllString(comment.Text, "")
			}
		}

		if constName != "" {
			overalMap[constName] = constComment
		}
	}
}

func createFile(pack, templateLoc string) {
	name := "gen_help_command.go"

	outFile, err := os.Create(name)
	if err != nil {
		logrus.Fatalf("Failed to create the file due to %v", err)
	}

	tmplate, err := ioutil.ReadFile(templateLoc)
	if err != nil {
		logrus.Fatalf("Failed to open the file due to %v", err)
	}

	templ := template.Must(template.New("out").Parse(string(tmplate)))

	outputString := createString()

	val := tmpStruct{
		PackageName: pack,
		HelpMessage: outputString,
	}

	if err := templ.Execute(outFile, val); err != nil {
		logrus.Fatalf("Failed to execute template %v", err)
	}

	if err := outFile.Close(); err != nil {
		logrus.Fatalf("Failed to close new file %v", err)
	}

	cmd := exec.Command("gofmt", "-w", name)
	if errOut, err := cmd.CombinedOutput(); err != nil {
		logrus.Fatalf("Failed to run command %v - %v", err, errOut)
	}
}

func createString() string {
	prefix := config.GetConfig().Prefix

	final := "The bot has many commands to interact with here is the list:\n\n"

	for k, v := range overalMap {
		final += fmt.Sprintf("%s%s - %s\n", prefix, k, v)
	}

	return strings.TrimSuffix(final, "\n")
}
