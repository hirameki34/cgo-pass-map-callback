package main

/*
#define _CALLBACK_PARAMS int size, char **keys, char **values, void *fp

typedef void (*_callback_t) (_CALLBACK_PARAMS);
void __makeCallback(_callback_t cb, _CALLBACK_PARAMS) {
	cb(size, keys, values, fp);
}
*/
import "C"
import "log/slog"

type Callback func(map[string]interface{})

var callback Callback

func register(cb Callback) {
	callback = cb
}

func makeCallback(fields map[string]interface{}) {
	callback(fields)
	slog.Info("MakeCallback", "fields", fields)
}
