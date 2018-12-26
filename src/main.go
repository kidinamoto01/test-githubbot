// Copyright 2017 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The simple command demonstrates a simple functionality which
// prompts the user for a GitHub username and lists all the public
// organization memberships of the specified username.
package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Fetch all the public organizations' membership of a user.
//
func FetchOrganizations(username string) ([]*github.Organization, error) {
	ctx := context.Background()
//curl https://api.github.com/kidinamoto01/repos?access_token=2c9c72a10ba7990cd12ad40b6b94517d5ff628e9
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "2c9c72a10ba7990cd12ad40b6b94517d5ff628e9"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	orgs, _, err := client.Organizations.List(context.Background(), username, nil)
	return orgs, err
}

// Fetch all the public organizations' membership of a user.
//
func FetchRepos(username string) ([]*github.Repository, error) {
	ctx := context.Background()
	//curl https://api.github.com/kidinamoto01/repos?access_token=2c9c72a10ba7990cd12ad40b6b94517d5ff628e9
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "2c9c72a10ba7990cd12ad40b6b94517d5ff628e9"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List(context.Background(), username, nil)
	return repos, err
}

func FetchPullRequest(ownername string,reporname string) ([]*github.PullRequest, error) {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "2c9c72a10ba7990cd12ad40b6b94517d5ff628e9"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	prs, _, err := client.PullRequests.List(context.Background(), ownername,reporname, nil)
	return prs, err
}

func main() {
	var username string
	fmt.Print("Enter GitHub username: ")
	fmt.Scanf("%s", &username)

	orgs, err := FetchOrganizations(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, org := range orgs {
		fmt.Printf("%v. %v\n", i+1,org.GetLogin())
	}

	repos, err := FetchRepos(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, repo := range repos {
		fmt.Printf("%v. %v\n", i+1,repo.GetFullName())
	}
	prs ,err:= FetchPullRequest("irisnet","irishub")
	for i, pr := range prs {
		fmt.Printf("%v. %v\n", i+1,pr.GetTitle())
	}
}