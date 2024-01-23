package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
	"strings"
)

const (
	// AppName is the name of the application
	AppName        = "topo"
	Version        = "0.2.0"
	AppCtrDir      = "appctr"
	MiddlewaresDir = "cmd/middlewares"
	CommonDir      = "common"
	CmdDir         = "cmd"
	DockersDir     = "dockers"
	DockersBackDir = "dockers/backend"
	HttpDir        = "http-requests"
	InternalsDir   = "internal"
	PkgDir         = "pkg"
	TestDir        = "test"
	UploadsDir     = "uploads"
)

//go:embed appctr/config.go
var configGoContent string

//go:embed appctr/main.go
var mainGoContent string

//go:embed appctr/router.go
var routerGoContent string

//go:embed appctr/logger.go
var loggerGoContent string

//go:embed appctr/db.go
var dbGoContent string

//go:embed cmd/main.go
var cmdMainGoContent string

//go:embed cmd/ioc.go
var cmdIOCGoContent string

//go:embed cmd/middlewares/middleware.go
var middlewaresGoContent string

//go:embed common/enums.go
var enumsGoContent string

//go:embed common/queries.go
var queriesGoContent string

//go:embed common/securityUtils.go
var securityUtilsGoContent string

//go:embed common/timeToDB.go
var timeToDBGoContent string

//go:embed docker-compose.yml
var dockerComposeYMLContent string

//go:embed dockers/backend/Dockerfile
var dockerfileContent string

//go:embed go-b.mod
var goModContent string

//go:embed go-b.sum
var goSumContent string

//go:embed http-request/examples.http
var exampleHTTPContent string

//go:embed init.sh
var initSHContent string

//go:embed .air.toml
var airContent string

//go:embed internal/fixtures/main.fixtures.go
var fixturesMainContent string

//go:embed internal/fixtures/user.fixtures.go
var fixturesUserContent string

//go:embed internal/handlers/user.handler.go
var handlersUserContent string

//go:embed internal/migrations/migration.go
var migrationsContent string

//go:embed internal/models/user.model.go
var modelsUserContent string

//go:embed internal/models/user.entity.go
var entityUserContent string

//go:embed internal/repositories/user.repository.go
var repositoriesUserContent string

//go:embed internal/services/user.service.go
var servicesUserContent string

//go:embed pkg/CLI/askToAuthorize.go
var askToAuthorizeContent string

func main() {
	var rootCmd = &cobra.Command{Use: AppName}

	var createCmd = &cobra.Command{
		Use:   "new",
		Short: "Create folder and files structure for a new TOPO app.",
		Run: func(cmd *cobra.Command, args []string) {
			var appName string
			var dbUser string
			var dbPassword = "topo"

			fmt.Print("Enter the name for the new app: ")
			_, err := fmt.Scan(&appName)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Print("Enter the user for the database: ")
			_, err = fmt.Scan(&dbUser)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			err = createFolderStructure(appName, dbUser, dbPassword)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			// chmod +777 appName folder recursively
			err = exec.Command("chmod", "-R", "777", appName).Run()
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Println("New TOPO app successfully created. Run: \n" +
				"cd " + appName + " && topo start \n To start the app")
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of TOPO",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("TOPO v" + Version)
		},
	}

	var initCmd = &cobra.Command{
		Use:   "start",
		Short: "Initialize TOPO app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Initializing app...")
			err := startApp()
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startApp() error {
	// Check if Docker is installed
	if err := checkDockerInstalled(); err != nil {
		return err
	}

	// Run docker-compose up -d --remove-orphans
	cmd := exec.Command("docker-compose", "up", "--remove-orphans")

	// Get reader streams for stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error obtaining stdout pipe: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("error obtaining stderr pipe: %v", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting docker-compose: %v", err)
	}

	// Use goroutines to read from stdout and stderr in parallel
	go printOutput(stdout, "stdout")
	go printOutput(stderr, "stderr")

	// Wait for the command to complete
	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("error running docker-compose: %v", err)
	}

	return nil
}

func printOutput(reader io.Reader, prefix string) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("[%s] %s\n", prefix, scanner.Text())
	}
}

func createFolderStructure(appName string, user string, password string) error {
	// Define the folder structure
	folders := []string{
		appName + "/" + AppCtrDir,
		appName + "/" + MiddlewaresDir,
		appName + "/" + CommonDir,
		appName + "/" + CmdDir,
		appName + "/" + DockersDir,
		appName + "/" + DockersBackDir,
		appName + "/" + HttpDir,
		appName + "/" + InternalsDir,
		appName + "/" + InternalsDir + "/fixtures",
		appName + "/" + InternalsDir + "/migrations",
		appName + "/" + InternalsDir + "/handlers",
		appName + "/" + InternalsDir + "/models",
		appName + "/" + InternalsDir + "/repositories",
		appName + "/" + InternalsDir + "/services",
		appName + "/" + PkgDir,
		appName + "/" + PkgDir + "/CLI",
		appName + "/" + TestDir,
		appName + "/" + UploadsDir,
	}

	// Create folders
	for _, folder := range folders {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create files
	err := createFilesFromExamples(appName, user, password)
	if err != nil {
		return err
	}

	return nil
}

func createFilesFromExamples(appName string, user string, password string) error {
	exampleFiles := map[string]string{
		appName + "/" + AppCtrDir + "/config.go":                          configGoContent,
		appName + "/" + AppCtrDir + "/main.go":                            mainGoContent,
		appName + "/" + AppCtrDir + "/router.go":                          routerGoContent,
		appName + "/" + AppCtrDir + "/logger.go":                          loggerGoContent,
		appName + "/" + AppCtrDir + "/db.go":                              dbGoContent,
		appName + "/" + CmdDir + "/main.go":                               cmdMainGoContent,
		appName + "/" + CmdDir + "/ioc.go":                                cmdIOCGoContent,
		appName + "/" + MiddlewaresDir + "/middleware.go":                 middlewaresGoContent,
		appName + "/" + CommonDir + "/enums.go":                           enumsGoContent,
		appName + "/" + CommonDir + "/queries.go":                         queriesGoContent,
		appName + "/" + CommonDir + "/securityUtils.go":                   securityUtilsGoContent,
		appName + "/" + CommonDir + "/timeToDB.go":                        timeToDBGoContent,
		appName + "/" + "docker-compose.yml":                              dockerComposeYMLContent,
		appName + "/" + DockersBackDir + "/Dockerfile":                    dockerfileContent,
		appName + "/" + "go.mod":                                          goModContent,
		appName + "/" + "go.sum":                                          goSumContent,
		appName + "/" + HttpDir + "/example.http":                         exampleHTTPContent,
		appName + "/" + "init.sh":                                         initSHContent,
		appName + "/" + ".air.toml":                                       airContent,
		appName + "/" + InternalsDir + "/fixtures/main.fixtures.go":       fixturesMainContent,
		appName + "/" + InternalsDir + "/fixtures/user.fixtures.go":       fixturesUserContent,
		appName + "/" + InternalsDir + "/handlers/user.handler.go":        handlersUserContent,
		appName + "/" + InternalsDir + "/migrations/migration.go":         migrationsContent,
		appName + "/" + InternalsDir + "/models/user.model.go":            modelsUserContent,
		appName + "/" + InternalsDir + "/models/user.entity.go":           entityUserContent,
		appName + "/" + InternalsDir + "/repositories/user.repository.go": repositoriesUserContent,
		appName + "/" + InternalsDir + "/services/user.service.go":        servicesUserContent,
		appName + "/" + PkgDir + "/CLI/askToAuthorize.go":                 askToAuthorizeContent,
	}

	for targetPath, content := range exampleFiles {
		// TOPO or topo content is replaced with the new app name

		// Replace POSTGRES_USER and POSTGRES_PASSWORD variables on docker compose
		// with the new database user and password
		if targetPath == appName+"/docker-compose.yml" {
			fmt.Printf("Replacing POSTGRES_USER and POSTGRES_PASSWORD variables on %s\n", targetPath)
			fmt.Printf("Replacing POSTGRES_USER with %s\n", user)
			fmt.Printf("Replacing POSTGRES_PASWORD with %s\n", password)
			content = replaceDatabaseCredentials(content, user, password)
		}
		content = strings.ReplaceAll(content, "topo", appName)
		content = strings.ReplaceAll(content, "TOPO", strings.ToUpper(appName))

		err := writeToFile(targetPath, content)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func checkDockerInstalled() error {
	// Check if Docker is installed
	cmd := exec.Command("docker", "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error checking Docker installation: %v\n%s", err, output)
	}

	fmt.Println("Docker is installed:", string(output))
	return nil
}

// Function to replace database credentials in docker-compose.yml content
func replaceDatabaseCredentials(content, user, password string) string {
	content = strings.ReplaceAll(content, "dbuser", user)
	content = strings.ReplaceAll(content, "dbpassword", password)
	return content
}
