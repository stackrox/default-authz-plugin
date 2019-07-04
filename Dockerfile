# Copyright 2019 StackRox Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.12.5-alpine3.9 AS build

ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/stackrox/default-authz-plugin
COPY . .
RUN go build -o /default-authz-plugin .

FROM scratch
COPY --from=build /default-authz-plugin /

COPY LICENSE /
COPY THIRD_PARTY_NOTICES/ /
COPY NOTICE /

ENTRYPOINT ["/default-authz-plugin"]
