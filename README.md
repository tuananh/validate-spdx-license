validate-spdx-license
---------------------

Just a demo to validate SPDX license string using [ABNF](https://en.wikipedia.org/wiki/Augmented_Backus%E2%80%93Naur_form).

A SPDX license string ABNF looks like this. I cut the `license-id` and `license-exception-id` short but you can see what a valid license string looks like.

```
alpha  = %x41-5A / %x61-7A ; A-Z / a-z
digit  = %x30-39 ; 0-9
idstring = 1*(alpha / digit / "-" / "." )
license-id = "CC-BY-SA-2.1-JP" / "GPL-2.0-or-later" / "AMDPLPA"
license-exception-id = "openvpn-openssl-exception" / "LLVM-exception" / "Bison-exception-2.2"
license-ref = [ "DocumentRef-" 1*(idstring) ":" ] "LicenseRef-" 1*(idstring)
simple-exp = license-id / license-ref
compound-exp = 1*1(simple-exp / simple-exp "WITH" license-exception-id)
compound-exp2 = 1*1(compound-exp "AND" compound-exp / compound-exp "OR" compound-exp / "(" compound-exp ")")
license-exp = 1*1(simple-exp / compound-exp2)
```

Related issue in apko: https://github.com/chainguard-dev/apko/issues/250

## References

- License data: https://github.com/spdx/license-list-data/tree/master/json
- SPDX license list: https://spdx.github.io/spdx-spec/SPDX-license-list/
- SPDX license expression: https://spdx.github.io/spdx-spec/SPDX-license-expressions/
