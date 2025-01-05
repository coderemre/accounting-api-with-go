package database

import (
	"log"
	"os"
	"os/exec"
)

func RunMigrations() {
	cmd := exec.Command(
		"migrate",
		"-path", "./migrations",
		"-database", "mysql://root:2646151Coder@/bank_app",
		"up",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}