package cognito

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoClient interface {
	Signup(email, password string) (error, string)
	ConfirmSignup(email, code string) (error, string)
	SignIn(email, password string) (error, string, *cognito.InitiateAuthOutput)
}

type awsCognitoClient struct {
	cognitoClient *cognito.CognitoIdentityProvider
	appClientID   string
	poolID        string
}

func NewCognitoClient(cognitoRegion string, cognitoAppClientID string, cognitoPoolID string) CognitoClient {
	conf := &aws.Config{Region: aws.String(cognitoRegion)}
	session, err := session.NewSession(conf)
	if err != nil {

		panic(err)
	}
	client := cognito.New(session)

	return &awsCognitoClient{
		cognitoClient: client,
		appClientID:   cognitoAppClientID,
		poolID:        cognitoPoolID,
	}
}

func (c *awsCognitoClient) Signup(email, password string) (error, string) {

	user := &cognito.SignUpInput{
		Username: aws.String(email),
		Password: aws.String(password),
		ClientId: aws.String(c.appClientID),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	}

	res, err := c.cognitoClient.SignUp(user)
	if err != nil {
		return err, ""
	}

	return nil, res.String()
}
func (c *awsCognitoClient) ConfirmSignup(email, code string) (error, string) {

	confirm := &cognito.ConfirmSignUpInput{
		Username:         aws.String(email),
		ClientId:         aws.String(c.appClientID),
		ConfirmationCode: aws.String(code),
	}

	res, err := c.cognitoClient.ConfirmSignUp(confirm)
	if err != nil {
		return err, ""
	}

	return nil, res.String()
}

func (c *awsCognitoClient) SignIn(email, password string) (error, string, *cognito.InitiateAuthOutput) {

	auth := &cognito.AdminInitiateAuthInput{
		AuthFlow: aws.String("ADMIN_USER_PASSWORD_AUTH"),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		}),
		UserPoolId: aws.String(c.poolID),
		ClientId:   aws.String(c.appClientID),
	}

	res, err := c.cognitoClient.AdminInitiateAuth(auth)
	if err != nil {
		return err, "", nil
	}

	return nil, res.String(), (*cognito.InitiateAuthOutput)(res)
}
