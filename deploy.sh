#!/bin/sh
 
faas-cli deploy -f stack.yml -e APIKEY=$1
