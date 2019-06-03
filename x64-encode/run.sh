#!/bin/bash

set -e

echo "###############################################"
echo "test.s contents:"
echo "---------------------------------------------"
cat test.s
echo "###############################################"
as -msyntax=intel test.s -o test.obj
objdump -d test.obj
