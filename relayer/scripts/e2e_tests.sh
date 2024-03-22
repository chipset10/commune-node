#!/usr/bin/env bash
# The Licensed Work is (c) 2024 Commune
# SPDX-License-Identifier: LGPL-3.0-only

go test -timeout 30m -p=1 $(go list ./... | grep 'e2e')
