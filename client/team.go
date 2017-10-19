// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "nmgapi": Team Resource Client
//
// Command:
// $ goagen
// --design=github.com/btoll/rest-go/design
// --out=$(GOPATH)/src/github.com/btoll/rest-go
// --regen=true
// --version=v1.3.0

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateTeamPath computes a request path to the create action of Team.
func CreateTeamPath() string {

	return fmt.Sprintf("/nmg/team/")
}

// Create a new sports team.
func (c *Client) CreateTeam(ctx context.Context, path string, payload *TeamPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateTeamRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateTeamRequest create the request corresponding to the create action endpoint of the Team resource.
func (c *Client) NewCreateTeamRequest(ctx context.Context, path string, payload *TeamPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteTeamPath computes a request path to the delete action of Team.
func DeleteTeamPath(id string) string {
	param0 := id

	return fmt.Sprintf("/nmg/team/%s", param0)
}

// Delete a sports team by id.
func (c *Client) DeleteTeam(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteTeamRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteTeamRequest create the request corresponding to the delete action endpoint of the Team resource.
func (c *Client) NewDeleteTeamRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListTeamPath computes a request path to the list action of Team.
func ListTeamPath() string {

	return fmt.Sprintf("/nmg/team/list")
}

// Get all teams
func (c *Client) ListTeam(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListTeamRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListTeamRequest create the request corresponding to the list action endpoint of the Team resource.
func (c *Client) NewListTeamRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowTeamPath computes a request path to the show action of Team.
func ShowTeamPath(id string) string {
	param0 := id

	return fmt.Sprintf("/nmg/team/%s", param0)
}

// Get a sports team by id.
func (c *Client) ShowTeam(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowTeamRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowTeamRequest create the request corresponding to the show action endpoint of the Team resource.
func (c *Client) NewShowTeamRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateTeamPath computes a request path to the update action of Team.
func UpdateTeamPath(id string) string {
	param0 := id

	return fmt.Sprintf("/nmg/team/%s", param0)
}

// Update a sports team by id.
func (c *Client) UpdateTeam(ctx context.Context, path string, payload *TeamPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateTeamRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateTeamRequest create the request corresponding to the update action endpoint of the Team resource.
func (c *Client) NewUpdateTeamRequest(ctx context.Context, path string, payload *TeamPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
