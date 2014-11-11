#!/usr/bin/env node

var program = require('commander');

program
  .command('create <site>', 'create site site')
  .parse(process.argv); 
#!/usr/bin/env node

var program = require('commander');

program
  .option('--nginx', 'install nginx package')
  .option('--php', 'install php package')
  .option('--mysql', 'install mysql package')
  .parse(process.argv);

console.log();

if (program.nginx) console.log('  ee stack : install nginx');
if (program.php) console.log('  ee stack : install php');
if (program.mysql) console.log('  ee stack: install mysql');

console.log();
