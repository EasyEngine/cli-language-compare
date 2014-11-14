#!/usr/bin/env node

var program = require('commander');

program
  .parse(process.argv);

var pkgs = program.args;

if (!pkgs.length) {
  console.error('listing what!');
  process.exit(1);
}

pkgs.forEach(function(pkg){
  console.log('FOO ' ,  pkg);
});


console.log();
