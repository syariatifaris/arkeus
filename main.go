package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"io/ioutil"
	"path/filepath"

	cp "github.com/otiai10/copy"
	"github.com/syariatifaris/arkeus/core/log/tokolog"
)

type cliParam struct {
	gosrc  string
	proj   string
	module string
	port   string
}

func initParams() *cliParam {
	gosrc := flag.String("gosrc", "", "$GOPATH src dir")
	proj := flag.String("project_path", "", "project path directory")
	module := flag.String("module", "hello", "sample module name")
	port := flag.String("port", "9093", "sample app port name")
	flag.Parse()

	return &cliParam{
		gosrc:  *gosrc,
		proj:   *proj,
		module: *module,
		port:   *port,
	}
}

func main() {
	param := initParams()

	if param.gosrc == "" || param.proj == "" {
		tokolog.ERROR.Println("$GOPATH/src or Project dir should be specified..")
		return
	}

	dir := fmt.Sprint(param.gosrc, "/", param.proj)

	tokolog.INFO.Println("copying project", param.proj, "destination:", dir)
	err := copyProject(dir)
	if err != nil {
		tokolog.ERROR.Println("unable to copy project", err.Error())
		return
	}

	tokolog.INFO.Println("creating module", param.module)
	err = renameModule(dir, param.module)
	if err != nil {
		tokolog.ERROR.Println("unable to process module dir", err.Error())
		return
	}

	tokolog.INFO.Println("writing files..")
	err = traverseDirWriteFile(dir, param.proj, param.module, param.port)
	if err != nil {
		tokolog.ERROR.Println("unable to process write file", err.Error())
		return
	}

	tokolog.INFO.Println("project", param.proj, "has been created")
}

func copyProject(dir string) error {
	return cp.Copy("project/sample", dir)
}

func renameModule(dir, module string) error {
	return os.Rename(fmt.Sprint(dir, "/app/module/cmodule"), fmt.Sprint(dir, "/app/module/", module))
}

func traverseDirWriteFile(dir, proj, module, port string) error {
	return filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.Contains(f.Name(), ".ctmpl") {
			//replace template
			read, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			rep := strings.NewReplacer(
				"[cproject_path]", proj,
				"[csample]", strings.ToLower(module),
				"[CSample]", strings.Title(strings.ToLower(module)),
				"[cport]", port,
			)

			nc := rep.Replace(string(read))
			err = ioutil.WriteFile(path, []byte(nc), 0)
			if err != nil {
				return err
			}

			//rename file
			rep = strings.NewReplacer(
				"csample", module,
				".ctmpl", ".go",
			)

			nn := rep.Replace(path)
			err = os.Rename(path, nn)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
