#include "libdemo.h"
#include "wrapper.h"

#include <iostream>
#include <map>
using std::cout;
using std::endl;
using std::map;
using std::string;

void printMap(map<string, string> map) {
    for (auto &[k, v] : map) {
        cout << k << ": " << v << endl;
    }
}

int main() {
    registerCallback(printMap);
    testCallback();
    return 0;
}
