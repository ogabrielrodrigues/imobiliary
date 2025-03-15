package main

import (
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pg/migrations",
		"--config",
		"./internal/store/pg/migrations/tern.conf",
		"-d",
		"-1",
	)

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
