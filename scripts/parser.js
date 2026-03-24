const fs = require('fs');
const path = require('path');
const { promisify } = require('util');

const readFile = promisify(fs.readFile);

class ConfigParser {
  constructor(configPath) {
    this.configPath = path.resolve(configPath);
  }

  async parse() {
    try {
      const configData = await readFile(this.configPath, 'utf8');
      const config = JSON.parse(configData);
      
      if (!config.host || !config.port) {
        throw new Error('Invalid configuration: host and port are required');
      }

      return {
        host: config.host,
        port: config.port,
        password: config.password || null,
        db: config.db || 0,
        tls: config.tls || false,
        connectTimeout: config.connectTimeout || 5000,
        maxRetries: config.maxRetries || 3
      };
    } catch (error) {
      throw new Error(`Failed to parse config: ${error.message}`);
    }
  }
}

module.exports = ConfigParser;