#!/usr/bin/env bash

buf generate --path $(glob dist/proto/furo/* | tr ' ' ,)