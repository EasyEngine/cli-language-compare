#!/usr/bin/env node

var program = require('commander');

program
  .command('foo', 'list packages installed')
  .option('-a, --all', 'hidden files listing')
  .parse(process.argv);
