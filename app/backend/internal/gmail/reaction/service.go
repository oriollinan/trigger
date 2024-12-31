package reaction

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/gmail"
	"trigger.com/trigger/internal/session"
	"trigger.com/trigger/internal/sync"
	"trigger.com/trigger/internal/user"

	"trigger.com/trigger/pkg/auth/oaclient"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/fetch"
)

func (m Model) MutlipleReactions(actionName string, ctx context.Context, action workspace.ActionNodeModel) error {
	accessToken, ok := ctx.Value(AccessTokenCtxKey).(string)
	if !ok {
		return errors.ErrAccessTokenCtx
	}

	switch actionName {
	case "send_email":
		return m.SendGmail(ctx, accessToken, action)
	}
	return nil
}

func createRawEmail(from string, to string, subject string, body string) (string, error) {
	var email bytes.Buffer
	email.WriteString(fmt.Sprintf("From: %s\r\n", from))
	email.WriteString(fmt.Sprintf("To: %s\r\n", to))
	email.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	email.WriteString("\r\n")
	email.WriteString(body)

	// Use URL encoding to avoid manual replacements of + and /
	rawMessage := base64.URLEncoding.EncodeToString(email.Bytes())

	return rawMessage, nil
}

func getGoogleUserByAccessToken(client *http.Client, accessToken string) (*GoogleUser, error) {
	res, err := fetch.Fetch(
		client,
		fetch.NewFetchRequest(
			http.MethodPost,
			fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken),
			nil,
			map[string]string{
				"Content-Type": "application/json",
			},
		),
	)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%v: %s", errors.ErrGmailSendEmail, bodyBytes)
	}
	googleUser, err := decode.Json[GoogleUser](res.Body)

	if err != nil {
		return nil, err
	}

	return &googleUser, nil
}

func (m Model) SendGmail(ctx context.Context, accessToken string, actionNode workspace.ActionNodeModel) error {
	userSsession, _, err := session.GetSessionByAccessTokenRequest(accessToken)
	if err != nil {
		return err
	}

	syncModel, _, err := sync.GetSyncAccessTokenRequest(accessToken, userSsession.UserId.Hex(), "google")
	if err != nil {
		return err
	}

	user, _, err := user.GetUserByIdRequest(accessToken, userSsession.UserId.Hex())
	if err != nil {
		return err
	}

	client, err := oaclient.New(ctx, gmail.Config(), syncModel)
	if err != nil {
		return err
	}

	googleUser, err := getGoogleUserByAccessToken(client, syncModel.AccessToken)

	if err == nil {
		// User must be extremely stupid
		if googleUser.Email == actionNode.Input["to"] {
			return errors.ErrSendingEmailToYourself
		}
	}

	rawEmail, err := createRawEmail(user.Email,
		actionNode.Input["to"], actionNode.Input["subject"], actionNode.Input["body"])

	if err != nil {
		return errors.ErrCreatingEmail
	}

	requestBody := map[string]string{
		"raw": rawEmail,
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	res, err := fetch.Fetch(
		client,
		fetch.NewFetchRequest(
			http.MethodPost,
			"https://gmail.googleapis.com/gmail/v1/users/me/messages/send",
			bytes.NewReader(body),
			map[string]string{
				"Content-Type": "application/json",
			},
		),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%v: %s", errors.ErrFailedToSendEmail, bodyBytes)
	}

	return nil
}
