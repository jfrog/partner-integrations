package main

import (
	_ "jfrog.com/jfrog-partnership/example-go/docs"
	"jfrog.com/jfrog-partnership/example-go/server"
)

// @title Example Go API
// @version 1.0.0
// @description Example Go API to demo JFrog Partner Integrations
// @termsOfService https://jfrog.com/eula/

// @contact.name JFrog Support
// @contact.url https://jfrog.com/support/
// @contact.email partner-support@jfrog.com

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenurl /token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	// load the server
	server.StartServer()
}
