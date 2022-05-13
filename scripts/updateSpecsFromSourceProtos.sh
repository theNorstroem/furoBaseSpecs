#!/usr/bin/env bash

# enable recursion for /**/*.xxx
shopt -s globstar dotglob

cd sourceProtos

protoc -I./ --furo-muspecs_out=:../muspecs **/*.proto