---
id: 20
database_id: 310337962
node_id: MDU6SXNzdWUzMTAzMzc5NjI=
status: closed
title: "improve demo-cross-origin test"
labels: ["type:refactoring"]
url: https://github.com/octopot/passport/issues/20
created_at: 2018-04-01T17:00:41Z
updated_at: 2018-10-17T14:14:35Z
---

# improve demo-cross-origin test

- up configured nginx (passport.local and demo.local)
- provide info to check hosts: `getent hosts passport.local`, it must exist and route to the nginx

related to #19
