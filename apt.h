struct Configuration;
extern struct Configuration *_config;

#ifdef __cplusplus
extern "C" {
#endif

void apt_init_config();
void apt_init_system();

char *apt_config_find(struct Configuration *c, const char *key);

#ifdef __cplusplus
}
#endif
