#!/usr/bin/env node

var program = require('commander');

program
  .command('create <site>', 'create site site')
  .parse(process.argv); 