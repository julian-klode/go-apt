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
