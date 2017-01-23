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

struct apt_config *apt_config_get_default(void)
{
    return (struct apt_config *) _config;
}

char *apt_config_find(struct apt_config *cfg, const char *key)
{
    std::string val = ((Configuration*)cfg)->Find(key, 0);

    if (val.empty()) {
        return NULL;
    } else {
        return strdup(val.c_str());
    }
}
