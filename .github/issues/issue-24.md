---
code: null
id: MDU6SXNzdWUzMTQ3MTE0NTU=
databaseId: 314711455
number: 24
url: https://github.com/octopot/passport/issues/24
title: "simplify JavaScript"
issueType: null
labels:
  - "type:refactoring"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-04-16T15:40:27Z
updatedAt: 2019-01-05T12:49:20Z
lastEditedAt: null
closedAt: 2019-01-05T12:49:20Z
---

# simplify JavaScript

refactoring:
- transport layer: not jQuery, just a function with fallback to jQuery.ajax
- calculate with a minimum amount of time
- send metadata only if fingerprint changed
