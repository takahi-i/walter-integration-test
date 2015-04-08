#!/bin/sh -e

# walter: a deployment pipeline template
# Copyright (C) 2014 Recruit Technologies Co., Ltd. and contributors
# (see CONTRIBUTORS.md)
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

. ./prepare

go test -i ./pipelines || exit 1
go test -v ./pipelines -race || exit 1

go test -i ./messengers || exit 1
go test -v ./messengers -race || exit 1

go test -i ./services || exit 1
go test -v ./services -race || exit 1
