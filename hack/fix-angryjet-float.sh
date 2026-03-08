#!/usr/bin/env bash
# Workaround for a bug in crossplane-tools/angryjet where single-resolution
# float pointer fields are generated using pkg/ptr instead of pkg/convert.
# See: https://github.com/crossplane/crossplane-tools/issues (ptrPkgPath vs convertPkgPath)
#
# The bug: angryjet emits
#   ptr.FromFloatPtrValue(x, "")  -> should be  reference.FromFloatPtrValue(x)
#   ptr.ToFloatPtrValue(x)        -> should be  reference.ToFloatPtrValue(x)
#
# These functions exist in crossplane-runtime/pkg/reference, not k8s.io/utils/ptr.

set -euo pipefail

find . -name 'zz_generated.resolvers.go' | while read -r f; do
  if grep -q 'ptr\.FromFloatPtrValue\|ptr\.ToFloatPtrValue' "$f"; then
    sed -i.bak \
      -e 's/ptr\.FromFloatPtrValue(\([^,]*\), "")/reference.FromFloatPtrValue(\1)/g' \
      -e 's/ptr\.ToFloatPtrValue/reference.ToFloatPtrValue/g' \
      "$f"
    rm -f "${f}.bak"
    echo "fixed: $f"
  fi
done
