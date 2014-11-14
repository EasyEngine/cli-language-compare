#!/usr/bin/env node

var program = require('commander');

program
  .option('--web', 'install web package')
  .option('--admin', 'install admin package')
  .option('--mail', 'install mail package')
  .option('--mailscanner', 'install mailscanner package')
  .option('--all', 'install all package')
  .parse(process.argv);

console.log();

if (program.web) console.log('Installed stack : web');
if (program.admin) console.log('Installed stack : admin');
if (program.mail) console.log('Installed stack : mail');
if (program.mailscanner) console.log('Installed stack : mailscanner');
if (program.all) console.log('Installed stack : all');