#!/bin/sh

deadcode . > deadcode-output.txt

# Since deadcode report files that are not called from tha main package,
# we need to remove these outputs that are only (but certainly) used in tests.
# For this, we'd need to follow the pattern to name these files like
# test_.*_util(s).go
# Surely, if either of these files contain a really dead code,
# it will be omitted from the report thus this approach is potentially unsafe,
# but it is the best we can do for now, until the deadcode team won't introduce
# a flag to ignore files by pattern.
sed -i '/test_.*util.go\|test_.*utils.go/d' ./deadcode-output.txt

if [ -s deadcode-output.txt ]; then
    echo "The following function(s) is/are not used in the code:"
    cat deadcode-output.txt
    exit 1
else
    echo "No dead code found!"
    exit 0
fi
