---
code: null
id: MDU6SXNzdWUzMTAzMzc5NjI=
databaseId: 310337962
number: 20
url: https://github.com/octopot/passport/issues/20
title: "improve demo-cross-origin test"
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
createdAt: 2018-04-01T17:00:41Z
updatedAt: 2018-10-17T14:14:35Z
lastEditedAt: null
closedAt: 2018-04-02T10:38:03Z
---

# improve demo-cross-origin test

- up configured nginx (passport.local and demo.local)
- provide info to check hosts: `getent hosts passport.local`, it must exist and route to the nginx

related to [[issue-19]]
