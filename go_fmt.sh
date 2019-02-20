#!/usr/bin/env bash

files=$(git diff --cached --name-only)
for file in $files
do
  if [ -e "$file" ] && [[ $file == *.tf ]]; then
    go fmt $file
    git add $file
  fi
done
