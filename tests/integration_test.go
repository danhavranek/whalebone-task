package tests

import (
	"encoding/json"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/danhavranek/whalebone-task/database"
	"github.com/danhavranek/whalebone-task/models"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const executablePath string = "../bin/whalebone-task"
const httpAddress string = "http://localhost:8090/"

var serverProcessCmd *exec.Cmd
var DB *gorm.DB

func TestGetNonExistentPerson(t *testing.T) {
	// Act
	resp, err := http.Get(httpAddress + "test")
	// Assert
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}

func TestCreatePerson(t *testing.T) {
	// Arrange
	personToBeStored := models.Person{ExternalId: uuid.NewString(), Name: "Test Person", Email: "testperson@example.com", DateOfBirth: "2020-01-01T12:12:34+00:00"}
	jsonData, err := json.Marshal(personToBeStored)
	if err != nil {
		t.Fatal(err)
	}
	reader := strings.NewReader(string(jsonData))
	// Act
	resp, err := http.Post(httpAddress+"save", "application/json", reader)
	// Assert
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	// Check db
	var storedPerson models.Person
	err = DB.Where("external_id = ?", personToBeStored.ExternalId).First(&storedPerson).Error
	if err != nil {
		t.Fatal(err)
	}
	if personToBeStored.Name != storedPerson.Name || personToBeStored.Email != storedPerson.Email || personToBeStored.DateOfBirth != storedPerson.DateOfBirth {
		t.Fatal("person stored badly")
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

	// DB connection
	DB, err = gorm.Open(sqlite.Open(database.DbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
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
