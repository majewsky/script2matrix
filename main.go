/*******************************************************************************
*
* Copyright 2019 Stefan Majewsky <majewsky@gmx.net>
*
* This program is free software: you can redistribute it and/or modify it under
* the terms of the GNU General Public License as published by the Free Software
* Foundation, either version 3 of the License, or (at your option) any later
* version.
*
* This program is distributed in the hope that it will be useful, but WITHOUT ANY
* WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
* A PARTICULAR PURPOSE. See the GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License along with
* this program. If not, see <http://www.gnu.org/licenses/>.
*
*******************************************************************************/

package main

import (
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/matrix-org/gomatrix"
	"github.com/sapcc/go-bits/logg"
)

var userIDRx = regexp.MustCompile(`^@(.+):(.+)$`)

func main() {
	if len(os.Args) < 2 {
		logg.Fatal("usage: %s <command> [<command-arg>...]", os.Args[0])
	}

	match := userIDRx.FindStringSubmatch(getenv("MATRIX_USER"))
	if match == nil {
		logg.Fatal("malformed environment variable: MATRIX_USER should look like @username:domain.name")
	}

	//login to Matrix account
	client, err := gomatrix.NewClient("https://"+match[2], "", "")
	failIf(err, "create Matrix client")
	resp, err := client.Login(&gomatrix.ReqLogin{
		Type:     "m.login.password",
		User:     match[1],
		Password: getenv("MATRIX_PASSWORD"),
	})
	failIf(err, "login to Matrix account")
	client.SetCredentials(resp.UserID, resp.AccessToken)

	//read remaining env vars now to detect presence
	targetRoomID := getenv("MATRIX_TARGET")

	//remove secrets from subcommand environment
	var commandEnvironment []string
	for _, env := range os.Environ() {
		if !strings.HasPrefix(env, "MATRIX_") {
			commandEnvironment = append(commandEnvironment, env)
		}
	}

	//collect output from subcommand
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Env = commandEnvironment
	outputBytes, err := cmd.CombinedOutput()
	output := string(outputBytes)
	if err != nil {
		if !strings.HasSuffix(output, "\n") {
			output += "\n"
		}
		output += err.Error()
	}

	_, err = client.SendMessageEvent(targetRoomID, "m.room.message", gomatrix.TextMessage{
		MsgType: "m.text",
		Body:    output,
	})
	failIf(err, "send message to "+targetRoomID)

	_, err = client.Logout()
	failIf(err, "logout from Matrix account")
}

func getenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		logg.Fatal("missing environment variable: " + key)
	}
	return val
}

func failIf(err error, msg string) {
	if err != nil {
		logg.Fatal("%s: %s", msg, err.Error())
	}
}
