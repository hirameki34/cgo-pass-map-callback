#include "libdemo.h"

#include <map>
using std::map;
using std::string;

typedef void (*callback_t)(map<string, string> fields);

void _callbackWrapper(_CALLBACK_PARAMS) {
    map<string, string> fields;
    for (int i = 0; i < size; i++) {
        fields[string(keys[i])] = string(values[i]);
    }
    ((callback_t)fp)(fields);
};

void registerCallback(callback_t callback) {
    __registerCallback(_callbackWrapper, (void *)callback);
}