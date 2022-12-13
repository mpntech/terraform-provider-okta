package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"net/http"
)

type ErrorPageCustomization struct {
	PageContent string `json:"pageContent,omitempty"`
}

type SignInPageCustomization struct {
	PageContent          string                         `json:"pageContent,omitempty"`
	WidgetCustomizations SignInPageWidgetCustomizations `json:"widgetCustomizations,omitempty"`
	WidgetVersion        string                         `json:"widgetVersion,omitempty"`
}

type SignInPageWidgetCustomizations struct {
	SignInLabel                             string `json:"signInLabel,omitempty"`
	UsernameLabel                           string `json:"usernameLabel,omitempty"`
	UsernameInfoTip                         string `json:"usernameInfoTip,omitempty"`
	PasswordLabel                           string `json:"passwordLabel,omitempty"`
	PasswordInfoTip                         string `json:"passwordInfoTip,omitempty"`
	ForgotPasswordLabel                     string `json:"forgotPasswordLabel,omitempty"`
	ForgotPasswordURL                       string `json:"forgotPasswordUrl,omitempty"`
	UnlockAccountLabel                      string `json:"unlockAccountLabel,omitempty"`
	UnlockAccountURL                        string `json:"unlockAccountUrl,omitempty"`
	HelpLabel                               string `json:"helpLabel,omitempty"`
	HelpURL                                 string `json:"helpUrl,omitempty"`
	CustomLink1Label                        string `json:"customLink1Label,omitempty"`
	CustomLink1URL                          string `json:"customLink1Url,omitempty"`
	CustomLink2Label                        string `json:"customLink2Label,omitempty"`
	CustomLink2URL                          string `json:"customLink2Url,omitempty"`
	AuthenticatorPageCustomLinkLabel        string `json:"authenticatorPageCustomLinkLabel,omitempty"`
	AuthenticatorPageCustomLinkURL          string `json:"authenticatorPageCustomLinkUrl,omitempty"`
	ClassicRecoveryFlowEmailOrUsernameLabel string `json:"classicRecoveryFlowEmailOrUsernameLabel,omitempty"`
	ShowPasswordVisibilityToggle            bool   `json:"showPasswordVisibilityToggle"`
	ShowUserIdentifier                      bool   `json:"showUserIdentifier"`
}

func getPageURL(brandID, page string) string {
	return "/api/v1/brands/" + brandID + "/pages/" + page + "/customized"
}

func (m *APISupplement) SetPageSignIn(ctx context.Context, brandID string, obj *SignInPageCustomization) (*okta.Response, error) {
	url := getPageURL(brandID, "sign-in")
	marsh, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	req, err := m.RequestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest(http.MethodPut, url, marsh)
	if err != nil {
		return nil, err
	}
	resp, err := m.RequestExecutor.Do(ctx, req, nil)
	if err != nil || resp.StatusCode != 200 {
		err = fmt.Errorf("%s %s: %w %d", http.MethodPut, req.URL.String(), err, resp.StatusCode)
	}
	return resp, err
}

func (m *APISupplement) GetPageSignIn(ctx context.Context, brandID string) (*SignInPageCustomization, *okta.Response, error) {
	obj := &SignInPageCustomization{}
	url := getPageURL(brandID, "sign-in")
	clone := *m.RequestExecutor
	exec := &clone
	req, err := exec.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := m.RequestExecutor.Do(ctx, req, obj)
	if err != nil {
		return nil, nil, err
	}
	return obj, resp, err
}

func (m *APISupplement) DeletePageSignIn(ctx context.Context, brandID string) (*okta.Response, error) {
	template := &SignInPageCustomization{}
	url := getPageURL(brandID, "sign-in")
	req, err := m.RequestExecutor.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	return m.RequestExecutor.Do(ctx, req, template)
}

func (m *APISupplement) SetPageError(ctx context.Context, brandID string, obj *ErrorPageCustomization) (*okta.Response, error) {
	url := getPageURL(brandID, "error")
	marsh, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	req, err := m.RequestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest(http.MethodPut, url, marsh)
	if err != nil {
		return nil, err
	}
	resp, err := m.RequestExecutor.Do(ctx, req, nil)
	if err != nil || resp.StatusCode != 200 {
		err = fmt.Errorf("%s %s: %w %d", http.MethodPut, req.URL.String(), err, resp.StatusCode)
	}
	return resp, err
}

func (m *APISupplement) GetPageError(ctx context.Context, brandID string) (*ErrorPageCustomization, *okta.Response, error) {
	obj := &ErrorPageCustomization{}
	url := getPageURL(brandID, "error")
	clone := *m.RequestExecutor
	exec := &clone
	req, err := exec.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := m.RequestExecutor.Do(ctx, req, obj)
	if err != nil {
		return nil, nil, err
	}
	return obj, resp, err
}

func (m *APISupplement) DeletePageError(ctx context.Context, brandID string) (*okta.Response, error) {
	template := &ErrorPageCustomization{}
	url := getPageURL(brandID, "error")
	req, err := m.RequestExecutor.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	return m.RequestExecutor.Do(ctx, req, template)
}
