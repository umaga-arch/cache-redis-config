import logging
import os
from typing import Optional

class Config:
    def __init__(self, config_path: str = 'config.yaml'):
        self.config_path = config_path
        self.config = self._load_config()

    def _load_config(self) -> dict:
        import yaml
        with open(self.config_path, 'r') as config_file:
            return yaml.safe_load(config_file)

    @property
    def redis(self) -> dict:
        return self.config['redis']

    @property
    def cache(self) -> dict:
        return self.config['cache']

    @property
    def logger(self) -> logging.Logger:
        return logging.getLogger('cache-redis-config')

    def get_redis_config(self) -> dict:
        return {
            'host': self.redis.get('host', 'localhost'),
            'port': self.redis.get('port', 6379),
            'db': self.redis.get('db', 0),
            'password': self.redis.get('password', None),
            'max_connections': self.redis.get('max_connections', 10),
        }

    def get_cache_config(self) -> dict:
        return {
            'maxsize': self.cache.get('maxsize', 100),
            'default_timeout': self.cache.get('default_timeout', 60),
        }