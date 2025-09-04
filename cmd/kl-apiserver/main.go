// Copyright 2025 杜杰 <dycloudlove@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ydcloud-dy/knowledgeBase.git

package main

import (
	"github.com/ydcloud-dy/knowledgeBase.git/cmd/kl-apiserver/app"
	"os"
)

func main() {

	command := app.NewKnowledgeBaseCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
