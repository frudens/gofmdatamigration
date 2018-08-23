package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"os"
	"strings"
	"fmt"
	"github.com/urfave/cli"
	"path/filepath"
)

const fmDataMigration = "./FMDataMigration"

var acc string
var pwd string
var clonePath string
var livePath string
var migratedPath string

func getCloneDir(dir string) error {

	// live dir
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return cli.NewExitError(err, 2)
	}

	// check files
	if len(files) == 0 {
		return cli.NewExitError("File not found in clone directory.", 3)
	}

	for _, file := range files {

		if file.IsDir() {
			err := getCloneDir(filepath.Join(dir, file.Name()))
			if err != nil {
				return err
			}
		}

		// filename ext
		clonePath = filepath.Join(dir, file.Name()) /* resources/clone/dir/file Colne.fmp12 */
		ext := filepath.Ext(clonePath) /* .fmp12 */

		// check .fmp12
		if ext == ".fmp12" {
			r := strings.NewReplacer(" Clone", "", " クローン", "")
			clonePathTrim := r.Replace(clonePath) /* resources/clone/dir/file.fmp12 */

			// livepath migratedPath
			livePath = strings.Replace(clonePathTrim, "/clone", "/live", -1) /* resources/live/dir/file.fmp12 */
			migratedPath = strings.Replace(clonePathTrim, "/clone", "/migrated", -1) /* resources/migrated/dir/file.fmp12 */

			// mkdir for migrated
			m := filepath.Dir(migratedPath) /* resources/migrated/dir */
			_, err := os.Stat(m)
			if err != nil {
				err = os.Mkdir(m,0777)
				if err != nil {
					return cli.NewExitError(err, 4)
				}
			}

			// run
			migration()

		} else {
			// error
			fmt.Printf("%s is not a .fmp 12 file.\n\n", clonePath)
		}
	}
	return nil
}

func migration() {
	// exec
	b, err := exec.Command(fmDataMigration, "-src_path", livePath, "-src_account", acc, "-src_pwd", pwd, "-clone_path", clonePath, "-clone_account", acc, "-clone_pwd", pwd, "-target_path", migratedPath).CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}

	// log replace
	logTemp := "__FILENAME__\n----------\n__LOG__\n"

	r := strings.NewReplacer("__FILENAME__", livePath, "__LOG__", string(b))
	l := r.Replace(logTemp)

	// file
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	f.WriteString(l)

	// log
	fmt.Println(l)
}

func main() {

	// app
	app := cli.NewApp()
	app.Name = "goFMDataMigration"
	app.Usage = "gofmdatamigration account password"
	app.Version = "0.1.0"
	app.Author = "frudens Inc. <https://frudens.com>"

	// action
	app.Action = func(c *cli.Context) error {
		acc = c.Args().Get(0)
		pwd = c.Args().Get(1)

		// args
		if c.NArg() == 0  {
			return cli.NewExitError("Argument required.", 1)
		}

		// base path
		cloneDir := filepath.Join("resources", "clone")

		err := getCloneDir(cloneDir)
		if err != nil {
			return err
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

