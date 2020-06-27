package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"
)

// path
var clonePath string
var prodPath string
var migratedPath string
var toolPath string

// option
var account string
var pwd string
var key string
var force bool
var ignore_valuelists bool
var ignore_accounts bool
var ignore_fonts bool
var verbose bool
var quiet bool

// value
var force_string string
var ignore_valuelists_string string
var ignore_accounts_string string
var ignore_fonts_string string
var verbose_string string
var quiet_string string

func getCloneDir(dir string) error {

	// prod dir
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
		ext := filepath.Ext(clonePath)              /* .fmp12 */

		// check .fmp12
		if ext == ".fmp12" {
			r := strings.NewReplacer(" Clone", "", " クローン", "")
			clonePathTrim := r.Replace(clonePath) /* resources/clone/dir/file.fmp12 */

			// prodpath migratedPath
			prodPath = strings.Replace(clonePathTrim, "clone", "prod", -1)         /* resources/prod/dir/file.fmp12 */
			migratedPath = strings.Replace(clonePathTrim, "clone", "migrated", -1) /* resources/migrated/dir/file.fmp12 */

			// mkdir for migrated
			m := filepath.Dir(migratedPath) /* resources/migrated/dir */
			_, err := os.Stat(m)
			if err != nil {
				err = os.MkdirAll(m, 0777)
				if err != nil {
					return cli.NewExitError(err, 4)
				}
			}

			// run
			migration()

		} else {
			// error
			fmt.Printf("%s is not a .fmp12 file.\n\n", clonePath)
		}
	}
	return nil
}

func migration() {
	// exec
	cmd := exec.Command(toolPath, "-src_path", prodPath, "-src_account", account, "-src_pwd", pwd, "-src_key", key, "-clone_path", clonePath, "-clone_account", account, "-clone_pwd", pwd, "-clone_key", key, "-target_path", migratedPath, ignore_valuelists_string, ignore_accounts_string, ignore_fonts_string, force_string, verbose_string, quiet_string)

	b, err := cmd.Output()
	if err != nil {
		log.Fatalf("FMDataMigration error:\n%s\n", string(b))
	}

	// log replace
	logTemp := "__FILENAME__\n----------\n__LOG__\n"

	r := strings.NewReplacer("__FILENAME__", prodPath, "__LOG__", string(b))
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

func getToolPath() string {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// check toolName
	var toolName string
	var toolDir string
	if runtime.GOOS == "windows" {
		toolName = "FMDataMigration.exe"
	} else {
		toolName = "FMDataMigration"
	}
	toolDir = "FMDataMigration"

	var path string
	if existDir(filepath.Join(d, toolDir)) {
		// check tool dir
		path = filepath.Join(d, toolDir, toolName)
	} else {
		// check current dir
		path = filepath.Join(d, toolName)
	}

	if existFile(path) {
		return path
	} else {
		return ""
	}
}

func existFile(filename string) bool {
	f, err := os.Stat(filename)
	if os.IsNotExist(err) || f.IsDir() {
		return false
	} else {
		return true
	}
}

func existDir(dir string) bool {
	f, err := os.Stat(dir)
	if os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}

func main() {

	// app
	app := cli.NewApp()
	app.Name = "goFMDataMigration"
	app.Usage = "goFMDataMigration is a command line tool for easy data migration using the FMDataMigration tool."
	app.Version = "0.3.0"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "frudens Inc.",
			Email: "komaki@frudens.com",
		},
	}

	// global option
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "account",
			Aliases:     []string{"a"},
			Usage:       "source Full Access or FMMigration account",
			Destination: &account,
		},
		&cli.StringFlag{
			Name:        "pwd",
			Aliases:     []string{"p"},
			Usage:       "source Full Access or FMMigration password",
			Destination: &pwd,
		},
		&cli.StringFlag{
			Name:        "key",
			Aliases:     []string{"k"},
			Usage:       "source decryption key",
			Destination: &key,
		},
		&cli.BoolFlag{
			Name:        "force",
			Usage:       "overwrite existing target file",
			Destination: &force,
		},
		&cli.BoolFlag{
			Name:        "ignore_valuelists",
			Usage:       "use custom value lists from clone instead of source",
			Destination: &ignore_valuelists,
		},
		&cli.BoolFlag{
			Name:        "ignore_accounts",
			Usage:       "use accounts and decryption key from clone instead of source",
			Destination: &ignore_accounts,
		},
		&cli.BoolFlag{
			Name:        "ignore_fonts",
			Usage:       "assume no font mapping required for field contents",
			Destination: &ignore_fonts,
		},
		&cli.BoolFlag{
			Name:        "verbose",
			Usage:       "verbose mode",
			Destination: &verbose,
		},
		&cli.BoolFlag{
			Name:        "quiet",
			Usage:       "quiet mode",
			Destination: &quiet,
		},
	}

	// action
	app.Action = func(c *cli.Context) error {

		// check toolPath
		toolPath = getToolPath()
		if toolPath == "" {
			return cli.NewExitError("FMDataMigration not found in current directory.", 5)
		}

		// check account
		if account == "" {
			return cli.NewExitError("account id required.", 1)
		}

		// base path
		cloneDir := filepath.Join("resources", "clone")

		// set option
		if force {
			force_string = "-force"
		}
		if ignore_valuelists {
			ignore_valuelists_string = "-ignore_valuelists"
		}
		if ignore_accounts {
			ignore_accounts_string = "-ignore_accounts"
		}
		if ignore_fonts {
			ignore_fonts_string = "-ignore_fonts"
		}
		if verbose {
			verbose_string = "-v"
		}
		if quiet {
			quiet_string = "-q"
		}

		// run
		err := getCloneDir(cloneDir)
		if err != nil {
			return err
		}
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

