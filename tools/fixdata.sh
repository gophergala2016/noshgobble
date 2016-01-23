#!/bin/bash

# Recursive file convertion windows-1251 --> utf-8

for file in data/*.txt; do
  echo " $file"
  mv $file $file.icv
  iconv -f WINDOWS-1251 -t UTF-8 $file.icv > $file
  rm -f $file.icv
done
