# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: build
build:
	go build ./v8/driver/...
	go build ./v9/driver/...

.PHONY: test
test:
	go test github.com/atomix/atomix/drivers/redis/v8/driver/...
	go test github.com/atomix/atomix/drivers/redis/v9/driver/...
