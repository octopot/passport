---
code: null
id: MDU6SXNzdWUzMTAwODAzMDA=
databaseId: 310080300
number: 15
url: https://github.com/octopot/passport/issues/15
title: "cross-origin problems"
issueType: null
labels:
  - "type:critical"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-03-30T14:31:19Z
updatedAt: 2018-03-30T18:41:50Z
lastEditedAt: null
closedAt: 2018-03-30T18:41:49Z
---

# cross-origin problems

for example in `file://` and different domains.

<!-- 2018-03-30T15:33Z https://github.com/octopot/passport/issues/15#issuecomment-377550936
```
Failed to load https://id.octolab.net/api/v1/tracker/fingerprint: Response to preflight request doesn't
pass access control check: No 'Access-Control-Allow-Origin' header is present on the requested
resource. Origin 'https://kamilsk.github.io' is therefore not allowed access. The response had HTTP
status code 415.
```
-->

<!-- 2018-03-30T15:34Z https://github.com/octopot/passport/issues/15#issuecomment-377551053
```
Request URL: https://id.octolab.net/api/v1/tracker/fingerprint
Request Method: OPTIONS
Status Code: 415
Remote Address: 37.139.24.208:443
Referrer Policy: no-referrer-when-downgrade
```
-->

<!-- 2018-03-30T15:34Z https://github.com/octopot/passport/issues/15#issuecomment-377551109
```
2018/03/30 15:32:52 "OPTIONS http://id.octolab.net/api/v1/tracker/fingerprint HTTP/1.0" from 46.188.123.85 - 415 0B in 27.776µs
```
-->

<!-- 2018-03-30T15:35Z https://github.com/octopot/passport/issues/15#issuecomment-377551238
```
172.18.0.9 - - [30/Mar/2018:15:32:52 +0000] "OPTIONS /api/v1/tracker/fingerprint HTTP/1.0" 415 0 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36" "46.188.123.85"
```
-->
