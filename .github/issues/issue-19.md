---
code: null
id: MDU6SXNzdWUzMTAyNTgyMTE=
databaseId: 310258211
number: 19
url: https://github.com/octopot/passport/issues/19
title: "safari issue"
issueType: null
labels:
  - "status:research"
  - "type:critical"
  - "help wanted"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-03-31T17:42:25Z
updatedAt: 2019-01-05T11:19:25Z
lastEditedAt: 2018-03-31T17:44:49Z
closedAt: 2019-01-05T11:19:25Z
---

# safari issue

Safari didn't send cookies
<img width="845" alt="safari bug" src="https://user-images.githubusercontent.com/1165416/38165921-53812e96-3524-11e8-8bc7-cd35a940e1ca.png">

<!-- 2018-03-31T19:55Z https://github.com/octopot/passport/issues/19#issuecomment-377719001
so, fallback is triggered `cookie not found, skip this check ("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/604.5.6 (KHTML, like Gecko) Version/11.0.3 Safari/604.5.6": "")`, but I have to find the right way
-->

<!-- 2018-04-01T13:11Z https://github.com/octopot/passport/issues/19#issuecomment-377785879
notes:
- https://github.com/github/fetch#sending-cookies
-->

<!-- 2018-04-02T10:21Z https://github.com/octopot/passport/issues/19#issuecomment-377901934
how to reproduce with #20:

1. `make cross-origin-up cross-origin-status && docker logs -f crossorigin_service_1`
2. `open http://demo.local/`
3. watch log
```
2018/04/02 10:19:16 - set cookie
2018/04/02 10:19:16 marker=cross-origin <nil>
2018/04/02 10:19:16 -
2018/04/02 10:19:16 "GET http://passport.local/js/passport.js HTTP/1.0" from 172.18.0.1 - 200 2506B in 321.3µs
...
[cors] 2018/04/02 10:19:17 Handler: Preflight request
[cors] 2018/04/02 10:19:17   Preflight response headers: map[Access-Control-Allow-Credentials:[true] Vary:[Origin Access-Control-Request-Method Access-Control-Request-Headers] Access-Control-Allow-Origin:[http://demo.local] Access-Control-Allow-Methods:[POST] Access-Control-Allow-Headers:[Content-Type]]
2018/04/02 10:19:17 "OPTIONS http://passport.local/api/v1/tracker/fingerprint HTTP/1.0" from 172.18.0.1 - 200 0B in 297.2µs
[cors] 2018/04/02 10:19:17 Handler: Actual request
[cors] 2018/04/02 10:19:17   Actual response added headers: map[Vary:[Origin] Access-Control-Allow-Origin:[http://demo.local] Access-Control-Allow-Credentials:[true]]
2018/04/02 10:19:17 - get cookie
2018/04/02 10:19:17 marker=cross-origin <nil>
2018/04/02 10:19:17 -
2018/04/02 10:19:18 "POST http://passport.local/api/v1/tracker/fingerprint HTTP/1.0" from 172.18.0.1 - 000 0B in 251.6425ms
```
-->

<!-- 2018-04-09T19:04Z https://github.com/octopot/passport/issues/19#issuecomment-379859781
- https://medium.com/@bluepnume/safaris-new-tracking-rules-and-enabling-cross-domain-data-storage-85241eea7483
- https://medium.com/@bluepnume/introducing-paypals-open-source-cross-domain-javascript-suite-95f991b2731d
- https://webkit.org/blog/7675/intelligent-tracking-prevention/
-->
