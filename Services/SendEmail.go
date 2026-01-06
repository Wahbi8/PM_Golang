
package Services

import "github.com/resend/resend-go/v2"

func SendEmail() {
    apiKey := "re_UgRSdtU7_74vHTKr3GAv3DcUcb2dBkJsp"

    client := resend.NewClient(apiKey)

    params := &resend.SendEmailRequest{
        From:    "onboarding@resend.dev",
        To:      []string{"wahbi.oussama08@gmail.com"},
        Subject: "Hello World",
        Html:    "<p>Congrats on sending your <strong>first email</strong>!</p>",
    }

    sent, err := client.Emails.Send(params)
}