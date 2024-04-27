package app

import (
	"fmt"
	"os/exec"
)

var encKey []byte

func init() {
	cmd := []string{"java", "-cp", "/mattermost/bin/keyReceiver.jar", "App"}
	output, err := exec.Command(cmd[0], cmd[1:]...).Output()

	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(output)
	outputString := string(output)

	encKey = []byte(outputString)
}
