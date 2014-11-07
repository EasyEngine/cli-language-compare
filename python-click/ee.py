# EasyEngine Command parsing using click framework


import click
import os
@click.command()

def cli():
	print 'Hello World!'

@click.group()
def cli():
    pass


@cli.group()
def site():
		pass

#---------------------- ee site command group-----------------------------#
@site.group()
def create():
		pass
@site.command('create')
@click.argument('site_name', nargs=1)
@click.option('--html', 'transformation', flag_value='html', default=True)
@click.option('--php', 'transformation', flag_value='php')
@click.option('--mysql', 'transformation', flag_value='mysql')
@click.option('--wp', 'transformation', flag_value='wp')
@click.option('--wpsubdir', 'transformation', flag_value='wpsubdir')
@click.option('--wpsubdomain', 'transformation', flag_value='wpsubdomain')

def create(site_name,transformation):
	click.echo('create site %s %s'%(site_name, transformation))


@site.group()
def delete():
		pass
@site.command('delete')
@click.argument('site_name')
def delete(site_name):
		click.echo('Delete site %s :('%site_name)

@site.group()
def update():
		pass

@site.command('update')
@click.argument('site_name')
@click.option('--php', 'transformation', flag_value='php')
@click.option('--mysql', 'transformation', flag_value='mysql')
@click.option('--wp', 'transformation', flag_value='wp')
@click.option('--wpsubdir', 'transformation', flag_value='wpsubdir')
@click.option('--wpsubdomain', 'transformation', flag_value='wpsubdomain')
def update(site_name,transformation):
		click.echo('Update site %s %s'%(site_name, transformation))

#---------------------- ee stack command group-----------------------------#
@cli.group()
def stack():
		pass

@stack.group()
def install():
		pass

@stack.command('install')
def install():
	click.echo('Hello ee stack install')

#---------------------- ee debug command group-----------------------------#
@cli.group()
def debug():
		pass


