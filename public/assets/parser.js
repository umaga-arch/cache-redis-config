const fs = require('fs');
const path = require('path');

class ConfigParser {
  constructor(filePath) {
    this.filePath = filePath;
    this.config = {};
  }

  parse() {
    try {
      const data = fs.readFileSync(this.filePath, 'utf8');
      const jsonConfig = JSON.parse(data);
      this.config = jsonConfig;
    } catch (err) {
      console.error(`Error parsing config file: ${err}`);
      process.exit(1);
    }
  }

  getConfig() {
    return this.config;
  }
}

module.exports = ConfigParser;