(function (context, signer, logger, config){
    'use strict';
    let ctx = (context.Passport = context.Passport || {'debug': config.debug}),
        payload = {'fingerprint': undefined},
        counter = 0, synced = false, lock = false;

    function log(msg) { ctx.debug && logger(config.prefix + msg); }

    function stop() {
        if (!synced && counter > config.limit) {
            logger(config.prefix + 'critical: payload was not be sent');
        }
        return synced || counter > config.limit;
    }

    function notify(handle, informer) {
        if (stop()) {
            clearInterval(handle);
            log(informer + ' is done');
            return;
        }
        if (lock) {
            return; // try next time
        }
        lock = true;
        counter++;

        // TODO https://github.com/github/fetch
        let data = new FormData();
        data.append('fingerprint', payload.fingerprint);
        fetch(
            new Request(
                config.endpoint,
                {method: 'POST', body: data, credentials: 'include'}
            )
        )
            .then(response => {
                if (response.status === 200) {
                    synced = true;
                    ctx.fingerprint = payload.fingerprint;
                    log('sender has synced a payload');
                }
            })
            .finally(() => {
                lock = false;
                log(informer + ' has sent a notification to ' + config.endpoint);
            });
        // sender({
        //     type: 'POST',
        //     url: config.endpoint,
        //     data: JSON.stringify(payload),
        //     contentType: 'application/json; charset=utf-8',
        //     xhrFields: { withCredentials: true },
        //     success: function () {
        //         synced = true;
        //         ctx.fingerprint = payload.fingerprint;
        //         log('sender has synced a payload');
        //     },
        //     complete: function () { lock = false; log(informer + ' has sent a notification to ' + config.endpoint); }
        // });
    }

    let corrector = setInterval((function () {
        let threshold = 1;
        new signer().get(function(result) { payload.fingerprint = result; });
        return function () {
            new signer().get(function(result) {
                if (result !== payload.fingerprint) {
                    payload.fingerprint = result;
                    threshold = 0;
                    log('corrector has made a correction');
                }
                if (++threshold >= config.threshold) {
                    notify(corrector, 'corrector');
                }
            });
        }
    }()), config.correct);

    let watcher = setInterval(function () { notify(watcher, 'watcher'); }, config.watch);
}(window, window.Fingerprint2, window.console.log, {
    'endpoint': '{{.Endpoint}}', 'prefix': 'passport: ',
    'limit': 3, 'threshold': 3,
    'correct': 250, 'watch': 1000,
    'debug': true
}));
