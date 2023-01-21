package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

var NewModule = &cobra.Command{
	Use:     "newmodule",
	Short:   "Create new module",
	Long:    "Create new module in app directory",
	Example: "bosque newmodule",
	Run: func(cmd *cobra.Command, args []string) {
		CreateNewModule()
	},
}

func CreateNewModule() {
	newModule := "newModule"
	newpath := filepath.Join(".", os.Getenv("MODULE_PATH")+newModule)
	checkDir, err := os.Stat(newpath)
	if checkDir == nil {
		err := os.MkdirAll(newpath, 0755)
		if err != nil {
			fmr.Println("Can't create directory")
			return
		}

	}
}