#!/bin/bash

file=${1:-"out.db"}

curl -s https://api.github.com/repos/karelnagel/avaandmed/releases/latest | jq -r '.assets[0].browser_download_url' | xargs -I {} curl -s -L {} | gunzip -c > $file