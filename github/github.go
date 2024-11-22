/**
 * Created by Goland
 * @file   github.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/11/22 16:06
 * @desc   github.go
 */

package github

import (
	"context"
	"github.com/google/go-github/v66/github"
)

// GetCommitHistory 获取仓库的提交历史
func GetCommitHistory(token string, owner string, repo string) (*github.RepositoryCommit, error) {
	client := github.NewClient(nil).WithAuthToken(token)
	options := github.CommitsListOptions{}
	options.Page = 1
	results, _, err := client.Repositories.ListCommits(context.Background(), owner, repo, &options)
	if err != nil {
		return nil, err
	}
	return results[0], nil
}
