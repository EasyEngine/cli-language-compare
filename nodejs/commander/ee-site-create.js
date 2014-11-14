#!/usr/bin/env node

var program = require('commander');

program
  .option('--php', 'create php site')
  .option('--html', 'create html site')
  .option('--mysql', 'create mysql site')
  .option('--wp', 'create wordpress site no-cache')
  .option('--wpsubdir', 'wordpress multisite with subdir setup')
  .option('--wpsubdomain', 'wordpress multisite with subdomain setup')
  .parse(process.argv);
var pkgs = program.args;
if (program.html) console.log('site %s created with html option', pkgs);
if (program.php) console.log('site %s with php option', pkgs);
if (program.mysql) console.log('site %s created with mysql option', pkg);
if (program.wp) console.log('site %s created with wp option', pkgs);
if (program.wpsubdir) console.log('site %s createed with wpsubdir option', pkgs);
if (program.wpsubdomain) console.log('site %s created with wpsubdomain option', pkgs);
console.log();
