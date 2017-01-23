#ifdef __cplusplus
extern "C" {
#endif

void apt_init_config();
void apt_init_system();

struct apt_cache;

struct apt_config;

struct apt_config *apt_config_get_default(void);
char *apt_config_find(struct apt_config *c, const char *key);

#ifdef __cplusplus
}
#endif
