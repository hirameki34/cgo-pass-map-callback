#include <map>

typedef void (*callback_t)(std::map<std::string, std::string> fields);
void registerCallback(callback_t callback);