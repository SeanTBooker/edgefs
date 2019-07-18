/*
 * Copyright (c) 2015-2018 Nexenta Systems, Inc.
 *
 * This file is part of EdgeFS Project
 * (see https://github.com/Nexenta/edgefs).
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package object

/*
#include "ccow.h"
*/
import "C"

import (
	"fmt"
	"os"
	"strings"

	"../efsutil"
	"../validate"
	"github.com/spf13/cobra"
)

func objectDelete(opath string) error {
	s := strings.SplitN(opath, "/", 4)

	return efsutil.ObjectDelete(s[0], s[1], s[2], s[3])
}

func objectExpunge(opath string) error {
	s := strings.SplitN(opath, "/", 4)

	return efsutil.ObjectExpunge(s[0], s[1], s[2], s[3])
}

var (
	expunge bool

	deleteCmd = &cobra.Command{
		Use:   "delete  <cluster>/<tenant>/<bucket>/<object>",
		Short: "delete an existing object",
		Long:  "delete an existing object",
		Args:  validate.Object,
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			if expunge {
				err = objectExpunge(args[0])
			} else {
				err = objectDelete(args[0])
			}
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	deleteCmd.Flags().BoolVarP(&expunge, "expunge", "e", false, "Expunge the object")
	ObjectCmd.AddCommand(deleteCmd)
}
