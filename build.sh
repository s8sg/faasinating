#!/bin/sh

faas-cli build -f stack.yml --parallel 4
#./faas-cli build -f stack.yml --parallel 4 --arg test=testy
