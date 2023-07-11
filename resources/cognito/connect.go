package cognito

import (
	"fmt"
	"log"

	"github.com/fmoral2/mux-api/resources"
)

func Cognito() {
	config, err := resources.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// randomString := resources.RandString()
	// randEmail := fmt.Sprintf("%s@%s.com", "test+automation+"+randomString, "sharklasers")
	cognito := NewCognitoClient(config.AwsRegion, config.AppID, config.PoolID)
	// err, sign := cognito.Signup(randEmail, config.Password)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Result  " + fmt.Sprintf("%s \n", sign))

	// err, confirm := cognito.ConfirmSignup("franciscomoral@hotmail.com", "253350")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Result  " + fmt.Sprintf("%s \n", confirm))

	err, _, initiateAuthOutput := cognito.SignIn("franciscomoral@hotmail.com", config.Password)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token  " + fmt.Sprintf("%s \n", *initiateAuthOutput.AuthenticationResult.IdToken))
}
