#include <apt-pkg/init.h>
#include "apt.h"

void apt_init_config()
{
    pkgInitConfig(*_config);
}

void apt_init_system()
{
    pkgInitSystem(*_config, _system);
}

char *apt_config_find(Configuration *cfg, const char *key)
{
    std::string val = cfg->Find(key, 0);

    if (val.empty()) {
        return strdup("");
    } else {
        return strdup(val.c_str());
    }
}
