package up

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Execute Installation for server
func InstallServices() {
	file, _ := os.Open("./config.json")

	b1 := make([]byte, 32*1024)
	n1, e := file.Read(b1)

	if e != nil {
		log.Fatal(e)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(b1[:n1])), &result)

	for w, v := range result {

		switch w {
		case "mongodb":
			if v == true {
				errAdd := runCommands("sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 4B7C549A058F8B6B")

				if errAdd != nil {
					fmt.Fprintln(os.Stderr, errAdd)
				}

				errEcho := runCommands("echo \"deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse\" | sudo tee /etc/apt/sources.list.d/mongodb.list")

				if errEcho != nil {
					fmt.Fprintln(os.Stderr, errEcho)
				}

				errUpdateApt := runCommands("sudo apt update")

				if errUpdateApt != nil {
					fmt.Fprintln(os.Stderr, errUpdateApt)
				}

				errInstallMongo := runCommands("sudo apt install mongodb-org")

				if errInstallMongo != nil {
					fmt.Fprintln(os.Stderr, errInstallMongo)
				}

				errEnabledMongo := runCommands("sudo systemctl enable mongod")

				if errEnabledMongo != nil {
					fmt.Fprintln(os.Stderr, errEnabledMongo)
				}

				errStartMongo := runCommands("sudo systemctl start mongod")

				if errStartMongo != nil {
					fmt.Fprintln(os.Stderr, errStartMongo)
				}
			}
		case "mysql":
			if v == true {
				errUpdateApt := runCommands("sudo apt-get update")

				if errUpdateApt != nil {
					fmt.Fprintln(os.Stderr, errUpdateApt)
				}

				errInstallMysql := runCommands("sudo apt-get install mysql-server")

				if errInstallMysql != nil {
					fmt.Fprintln(os.Stderr, errInstallMysql)
				}
			}
		case "crystalLang":
			if v == true {
				errAptConfig := runCommands("curl -sSL https://dist.crystal-lang.org/apt/setup.sh | sudo bash")

				if errAptConfig != nil {
					fmt.Fprintln(os.Stderr, errAptConfig)
				}

				errKeyCrystal := runCommands("curl -sL \"https://keybase.io/crystal/pgp_keys.asc\" | sudo apt-key add - echo \"deb https://dist.crystal-lang.org/apt crystal main\" | sudo tee /etc/apt/sources.list.d/crystal.list")

				if errKeyCrystal != nil {
					fmt.Fprintln(os.Stderr, errKeyCrystal)
				}

				errUpdateApt := runCommands("sudo apt-get update")

				if errUpdateApt != nil {
					fmt.Fprintln(os.Stderr, errUpdateApt)
				}

				errInstallCrystal := runCommands("sudo apt install crystal")

				if errInstallCrystal != nil {
					fmt.Fprintln(os.Stderr, errInstallCrystal)
				}

				errInstallLibSsl := runCommands("sudo apt install libssl-dev")

				if errInstallLibSsl != nil {
					fmt.Fprintln(os.Stderr, errInstallLibSsl)
				}

				errinstallLibXML := runCommands("sudo apt install libxml2-dev")

				if errinstallLibXML != nil {
					fmt.Fprintln(os.Stderr, errinstallLibXML)
				}

				errInstallLibYaml := runCommands("sudo apt install libyaml-dev")

				if errInstallLibYaml != nil {
					fmt.Fprintln(os.Stderr, errInstallLibYaml)
				}

				errInstallLibGmp := runCommands("sudo apt install libgmp-dev")

				if errInstallLibGmp != nil {
					fmt.Fprintln(os.Stderr, errInstallLibGmp)
				}

				errInstallLibReadline := runCommands("sudo apt install libreadline-dev")

				if errInstallLibReadline != nil {
					fmt.Fprintln(os.Stderr, errInstallLibReadline)
				}

				errInstallLibZ := runCommands("sudo apt install libz-dev")

				if errInstallLibZ != nil {
					fmt.Fprintln(os.Stderr, errInstallLibZ)
				}

			}
		case "driverMongo":
			if v == true {
				errUpdateApt := runCommands("sudo apt-get update")

				if errUpdateApt != nil {
					fmt.Fprintln(os.Stderr, errUpdateApt)
				}

				errInstallDriverMongoC := runCommands("sudo apt install libmongoc-dev libmongoc-1.0-0 libmongoclient-dev")

				if errInstallDriverMongoC != nil {
					fmt.Fprintln(os.Stderr, errInstallDriverMongoC)
				}
			}
		case "redis":
			errUpdateApt := runCommands("sudo apt-get update")

			if errUpdateApt != nil {
				fmt.Fprintln(os.Stderr, errUpdateApt)
			}

			errInstallRedis := runCommands("sudo apt-get install redis-server")

			if errInstallRedis != nil {
				fmt.Fprintln(os.Stderr, errInstallRedis)
			}

			errEnableRedis := runCommands("sudo systemctl enable redis-server.service")

			if errEnableRedis != nil {
				fmt.Fprintln(os.Stderr, errEnableRedis)
			}

		case "gitclone":
			value := fmt.Sprintf("git clone  %v", v)

			err := runCommands(value)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}

func runCommands(command string) error {
	command = strings.TrimSuffix(command, "\n")
	arrCommandStr := strings.Fields(command)

	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
		// add another case here for custom commands.
	}

	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
