//go:build userfetcher

package main

import userfetchers "lunarisnia/go-build-tags-experiment/internal/user_fetchers"

func init() {
	handler = userfetchers.UserFetcherHandler{}
}
