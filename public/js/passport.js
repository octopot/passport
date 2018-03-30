(function (context, base, signer, sender, logger, debug){
    'use strict';
    context.Passport = {'fingerprint': undefined, 'metadata': undefined, 'counter': undefined, 'synced': undefined};

    var payload = {'fingerprint': undefined, 'metadata': undefined, 'counter': undefined}, synced = false, lock = false;

    function log(msg) { debug && logger('passport: ' + msg); }

    function notify(handle, informer) {
        log(informer + ' has sent a notification');
        var url = (base.substr(-1) === '/' ? base.substr(0, base.len - 1) : base) + '/api/v1/tracker/fingerprint';
        sender({
            type: 'POST',
            url: url,
            data: JSON.stringify(payload),
            contentType: 'application/json; charset=utf-8',
            success: function () {
                synced = true;
                clearInterval(handle);
                context.Passport.fingerprint = payload.fingerprint;
                context.Passport.metadata = payload.metadata;
                context.Passport.counter = payload.counter;
                log('sender has synced a payload');
                log(informer + ' is done');
            },
            complete: function () { lock = false; }
        });
    }

    function stop() {
        if (!synced && payload.counter > 5) {
            logger('passport: critical: payload was not be sent');
        }
        return synced || payload.counter > 5;
    }

    var corrector = setInterval(function () {
        !lock && new signer().get(function(result, components) {
            if (result !== payload.fingerprint) {
                payload.fingerprint = result;
                payload.metadata = components;
                payload.counter = 0;
                log('corrector has made a correction')
            }
            payload.counter++;
            if (stop()) {
                clearInterval(corrector);
                log('corrector is done');
                return;
            }
            if (payload.counter >= 3) {
                lock = true;
                notify(corrector, 'corrector');
            }
        })
    }, 100);

    var watcher = setInterval(function () {
        if (stop())  {
            clearInterval(watcher);
            log('watcher is done');
            return
        }
        if (!lock) {
            lock = true;
            notify(watcher, 'watcher');
        }
    }, 1000);
}(window, 'http://localhost:8080/', window.Fingerprint2, window.jQuery.ajax, window.console.log, true));
