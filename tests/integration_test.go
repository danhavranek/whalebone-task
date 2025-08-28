package tests

import (
	"fmt"
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
	externalId := uuid.NewString()
	name := "Test Person"
	email := "testperson@example.com"
	dateOfBirth := "2020-01-01T12:12:34+00:00"

	requestJson := fmt.Sprintf(`{"external_id":"%s","name":"%s","email":"%s","date_of_birth":"%s"}`, externalId, name, email, dateOfBirth)
	reader := strings.NewReader(requestJson)
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
	err = DB.Where("external_id = ?", externalId).First(&storedPerson).Error
	if err != nil {
		t.Fatal(err)
	}
	if name != storedPerson.Name || email != storedPerson.Email || dateOfBirth != storedPerson.DateOfBirth {
		t.Fatal("person stored badly")
	}
}

func TestGetPerson(t *testing.T) {
	// Arrange
	personToBeRecieved := models.Person{ExternalId: uuid.NewString(), Name: "Test Person", Email: "testperson@example.com", DateOfBirth: "2020-01-01T12:12:34+00:00"}
	err := DB.Create(personToBeRecieved).Error
	if err != nil {
		panic(err)
	}
	// Act
	var resp *http.Response
	resp, err = http.Get(httpAddress + personToBeRecieved.ExternalId)
	// Assert
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestGetPersonMethodNotAllowed(t *testing.T) {
	// Arrange
	externalId := uuid.NewString()
	// Act
	resp, _ := http.Post(httpAddress+externalId, "text/plain", nil)
	// Assert
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", resp.StatusCode)
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
