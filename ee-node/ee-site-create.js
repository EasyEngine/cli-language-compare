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
if (program.html) console.log('create site %s with --html', pkgs);
if (program.php) console.log('create site %s with --php', pkgs);
if (program.mysql) console.log('create site %s with --mysql', pkg);
if (program.wp) console.log('create site %s with --wp', pkgs);
if (program.wpsubdir) console.log('create site %s with --wpsubdir', pkgs);
if (program.wpsubdomain) console.log('create site %s with --wpsubdomain', pkgs);
console.log();
