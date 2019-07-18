#!/bin/bash
##
## Copyright (c) 2015-2018 Nexenta Systems, inc.
##
## This file is part of EdgeFS Project
## (see https://github.com/Nexenta/edgefs).
##
## Licensed to the Apache Software Foundation (ASF) under one
## or more contributor license agreements.  See the NOTICE file
## distributed with this work for additional information
## regarding copyright ownership.  The ASF licenses this file
## to you under the Apache License, Version 2.0 (the
## "License"); you may not use this file except in compliance
## with the License.  You may obtain a copy of the License at
##
##   http://www.apache.org/licenses/LICENSE-2.0
##
## Unless required by applicable law or agreed to in writing,
## software distributed under the License is distributed on an
## "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
## KIND, either express or implied.  See the License for the
## specific language governing permissions and limitations
## under the License.
##

TEST_DIR=/opt/nedge/src/ccow/test

rm -rf /data/ata-Sa*/{nameindex,chunks}/

echo "RUNNING CLUSTER TESTS"

$TEST_DIR/cluster_test
sleep 2

$TEST_DIR/put_test
sleep 2

$TEST_DIR/iter_test
sleep 2

$TEST_DIR/vol_test
sleep 2

$TEST_DIR/clone_test
sleep 2

$TEST_DIR/get_test

echo "COMPLETE"
