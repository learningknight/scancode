/*
* Copyright (C) 2021 THL A29 Limited, a Tencent company.  All rights reserved.
* This source code is licensed under the Apache License Version 2.0.
*/

package version

var (
	gitTag       = ""
	gitCommit    = "$Format:%H$"          // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState = "not a git tree"       // state of git tree, either "clean" or "dirty"
	buildDate    = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)
