(function (context, signer, sender, logger, config){
    'use strict';
    var ctx = (context.Passport = context.Passport || {'debug': config.debug}),
        payload = {'fingerprint': undefined, 'metadata': undefined},
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
        sender({
            type: 'POST',
            url: config.endpoint,
            data: JSON.stringify(payload),
            contentType: 'application/json; charset=utf-8',
            xhrFields: { withCredentials: true },
            success: function () {
                synced = true;
                ctx.fingerprint = payload.fingerprint;
                ctx.metadata = payload.metadata;
                log('sender has synced a payload');
            },
            complete: function () { lock = false; log(informer + ' has sent a notification to ' + config.endpoint); }
        });
    }

    var corrector = setInterval((function () {
        var threshold = 1;
        new signer().get(function(result, components) { payload.fingerprint = result; payload.metadata = components; });
        return function () {
            new signer().get(function(result, components) {
                if (result !== payload.fingerprint) {
                    payload.fingerprint = result;
                    payload.metadata = components;
                    threshold = 0;
                    log('corrector has made a correction');
                }
                if (++threshold >= config.threshold) {
                    notify(corrector, 'corrector');
                }
            });
        }
    }()), config.correct);

    var watcher = setInterval(function () { notify(watcher, 'watcher'); }, config.watch);
}(window, window.Fingerprint2, window.jQuery.ajax, window.console.log, {
    'endpoint': '{{.Endpoint}}', 'prefix': 'passport: ',
    'limit': 3, 'threshold': 3,
    'correct': 250, 'watch': 1000,
    'debug': false
}));
