# Utility Belt

A library of various packages of pure functions used in production across the Rhythmic Sound stack (including Simple Data Service).

The library uses no custom types or has any side effects other that writing to Std.out (see Depreciation and modification process). Therefore is useful for a wide array of other purposes outside of Rhythmic Sound. We've open sourced it for wider consumption and feedback. 

Other open source libraries are used within the functions. For stability, some are copied into this library with comments explaining code origin and any alterations on the original code.

## Organisation
Functions are grouped and organised within packages of relavent topics for easier use and more appropriate self documenting naming conventions. 

## Depreciation and modificaction process
Once published in a release, the function's API is fixed for that major release and can only be depreciated or modified in the next major release. 

Where modified, the function should `log` the potentially unexpected behavior from previous releases to std.out to alert users. 

Depreciation log `warnings` should also be added to the previous release via std.out

Developers should check for this warnings when updating to a new major version.

> There will be some movement of functions between packages until the release of major version 1.0.0

## ToDo

- [ ] Add tests in each package
- [ ] Fix directory structure for v1.0.0
- [ ] Alpha release