---
code: null
id: MDU6SXNzdWUzOTIwNzUwOTU=
databaseId: 392075095
number: 52
url: https://github.com/octopot/passport/issues/52
title: "wrap migration with transaction"
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
createdAt: 2018-12-18T09:51:49Z
updatedAt: 2018-12-20T17:28:25Z
lastEditedAt: null
closedAt: 2018-12-20T17:28:25Z
---

# wrap migration with transaction

```
begin;
...
commit;
```

<!-- 2018-12-20T17:28Z https://github.com/octopot/passport/issues/52#issuecomment-449075026
won't do
-->
