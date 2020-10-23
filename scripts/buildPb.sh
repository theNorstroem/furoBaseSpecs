#!/usr/bin/env bash
TARGETDIR="../pb/"

# enable recursion for /**/*.xxx
shopt -s globstar dotglob

cd dist/proto


# clear
rm -rf $TARGETDIR
mkdir $TARGETDIR


FILES=./furo/*.proto

protoc --proto_path=./ \
-I. \
--go_out=\
:$TARGETDIR $FILES

# move the files
cd $TARGETDIR
mv github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo furo
rm -rf github.com
