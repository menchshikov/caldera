package generator

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/takama/caldera/pkg/config"
	"github.com/takama/caldera/pkg/helper"
)

// Run generator
func Run(cfg *config.Config) {
	helper.LogF("Copy base templates", copyTemplates(
		path.Join(cfg.Directories.Templates, config.Base),
		cfg.Directories.Service,
	))
	helper.LogF("Render templates", render(cfg))
	helper.LogF("Could not change directory", os.Chdir(cfg.Directories.Service))
	log.Println("Vendors initialization")

	log.Println("Git repository initialization")
	helper.LogF("Could not init git", Exec("git", "init"))
	helper.LogF("Could not add git files", Exec("git", "add", "--all"))
	helper.LogF("Could not commit git files", Exec("git", "commit", "-m", "'Initial commit'"))
	fmt.Printf("New repository was created, use command 'cd %s'", cfg.Directories.Service)
}

// Exec runs the commands
func Exec(command ...string) error {
	execCmd := exec.Command(command[0], command[1:]...) // nolint: gosec
	execCmd.Stderr = os.Stderr
	execCmd.Stdout = os.Stdout
	execCmd.Stdin = os.Stdin
	return execCmd.Run()
}
