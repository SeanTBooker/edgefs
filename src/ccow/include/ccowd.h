/*
 * Copyright (c) 2015-2018 Nexenta Systems, inc.
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
#ifndef __CCOWD_H__
#define __CCOWD_H__

#ifdef	__cplusplus
extern "C" {
#endif

#define CCOWD_IPC_PATH "%s/var/run/ccowd.ipc"
#define CCOWD_IPC_ADDRESS "ipc://" CCOWD_IPC_PATH

struct ccowd_params {
	int log_flush;
	int log_recovery;
	int mdcache_cleanup;
	int plevel;
	char *name;
	int log_err;
	int relocate_mdoffload;
	int relocate_ttag;
};

extern int ccow_daemon_init(struct ccowd_params *p);
extern void ccow_daemon_term();
extern void ccow_daemon_print_cache_stats();

#ifdef	__cplusplus
}
#endif

#endif
