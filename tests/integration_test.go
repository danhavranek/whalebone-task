package tests

import (
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
)

const executablePath string = "../bin/whalebone-task"
const httpAddress string = "http://localhost:8090/"

var serverProcessCmd *exec.Cmd

func TestGetNonExistentPerson(t *testing.T) {
	// Get non-existent person
	resp, err := http.Get(httpAddress + "test")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}

func setup() {
	// Service start
	serverProcessCmd = exec.Command(executablePath)
	serverProcessCmd.Stdout = os.Stdout
	serverProcessCmd.Stderr = os.Stderr
	// TODO: setup tmp env vars

	err := serverProcessCmd.Start()
	if err != nil {
		panic(err)
	}

	// Let it boot
	time.Sleep(2 * time.Second)
}

func teardown() {
	if serverProcessCmd != nil && serverProcessCmd.Process != nil {
		_ = serverProcessCmd.Process.Kill()
		// Wait for the process to exit
		_ = serverProcessCmd.Wait()
	}
	// Remove test database
	_ = os.RemoveAll("app")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
