package main

import (
	"fmt"
	"github.com/JackMaarek/Go-release-test/database"
	"github.com/JackMaarek/Go-release-test/mailer"
	"github.com/JackMaarek/Go-release-test/shared"
	"github.com/JackMaarek/Go-release-test/shared/repositories"
)

func main() {
	dbdetails := database.DBDetails{
		Username: "root",
		Password: "root",
		Name:     "symfony",
		Host:     "127.0.0.1",
		Port:     3308,
		Engine:   "mysql",
	}
	db, err := database.Connect(&dbdetails)
	if err != nil {
		fmt.Println("Cannot connect to db")
	}
	SBDetails := &mailer.SBDetails{
		Url:         "https://api.sendinblue.com/v3/smtp/email",
		ApiKey:      "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",
		SenderEmail: "riskandme@harmonie-technologie.com",
		SenderName:  "riskandme",
	}

	repo := repositories.Repository{DB: db}

	mailerService := shared.MailerService{
		SBDetails:      SBDetails,
		Repository: 	&repo,

	}

	err = mailerService.SendCompletedCampaignMail("e5651054-f54a-48ff-b25c-8235129d9e0d", "6a9e607a-cfb7-4ba1-a111-41129a125821")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("message sent !")
	}
}
