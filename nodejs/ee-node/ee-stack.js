#!/usr/bin/env node

var program = require('commander');

program
  .command('install', 'packages install')
  .parse(process.argv);
