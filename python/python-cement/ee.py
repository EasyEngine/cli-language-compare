#!/usr/bin/env python3.4

from cement.core import foundation, controller, handler

# define application controllers
class EEBaseController(controller.CementBaseController):
		class Meta:
				label = 'base'
				interface = controller.IController
				description = "EasyEngine Does Amazing Things"

		@controller.expose(help="EayEngine command", hide=True)
		def default(self):
				self.app.args.parse_args(['--help'])

class Site(controller.CementBaseController):
		"""This controller commands are 'stacked' onto the base controller."""

		class Meta:
				label = 'site'
				interface = controller.IController
				stacked_on = 'base'
				stacked_type = 'nested'
				description = "Site commands are work under this controller"

## argument selection is an issue
class Create(controller.CementBaseController):
		"""This controller commands are 'stacked' onto the site controller."""

		class Meta:
				label = 'create'
				interface = controller.IController
				stacked_on = 'site'
				stacked_type = 'nested'
				description = "create commands are listed under this controller"
				arguments = [
						(['site_name'], dict(help="site domain name like example.com")),
						(['--html'], dict(help="create html site example.com", action='store_true')),
						(['--php'], dict(help="create php site example.com", action='store_true')),
						(['--mysql'], dict(help="create mysql site example.com", action='store_true')),
						(['--wp'], dict(help="create wp site example.com", action='store_true')),
						(['--wpsubdir'], dict(help="create wpsubdir site example.com", action='store_true')),
						(['--wpsubdomain'], dict(help="create wpsubdomain site example.com", action='store_true')),
						]

		@controller.expose(help="command under the base namespace")
		def default(self):
				if self.app.pargs.html:
						print ("site", self.app.pargs.site_name ,"created with html option")
				if self.app.pargs.php:
						print ("site", self.app.pargs.site_name, "created with php option")
				if self.app.pargs.mysql:
						print ("site", self.app.pargs.site_name, "created with mysql option")
				if self.app.pargs.wp:
						print ("site", self.app.pargs.site_name, "created with wp option")
				if self.app.pargs.wpsubdir:
						print ("site", self.app.pargs.site_name, "created with wpsubdir option")
				if self.app.pargs.wpsubdomain:
						print ("site", self.app.pargs.site_name, "created with wpsubdomain option")


class Stack(controller.CementBaseController):
		"""This controller commands are *not* 'stacked' onto the base controller."""

		class Meta:
				label = 'stack'
				interface = controller.IController
				stacked_on = 'base'
				stacked_type = 'nested'
				description = "Stack commands are listed under this controller"



class Install(controller.CementBaseController):
		"""This controller commands are *not* 'stacked' onto the base controller."""

		class Meta:
				label = 'install'
				interface = controller.IController
				stacked_on = 'stack'
				stacked_type = 'nested'
				description = "install commands	are listed under this controller"
				arguments = [
						(['--web'], dict(help="Install web stack", action='store_true')),
						(['--admin'], dict(help="Install admin stack", action='store_true')),
						(['--mail'], dict(help="Install mail stack", action='store_true')),
						(['--mailscanner'], dict(help="Install mailscanner stack", action='store_true')),
						(['--all'], dict(help="Install all stack", action='store_true')),
						]

		@controller.expose(help="default command for third_controller", hide=True)
		def default(self):
				
				if self.app.pargs.web:
								print("Installed stack : web")
				if self.app.pargs.admin:
								print( "Installed stack : admin")
				if self.app.pargs.mail:
								print( "Installed stack : mail")
				if self.app.pargs.mailscanner:
								print( "Installed stack : mailscanner")
				if self.app.pargs.all:
								print( "Installed stack : all")

try:
		# create the application
		app = foundation.CementApp('ee')

		# register non-base controllers
		handler.register(EEBaseController)
		handler.register(Site)
		handler.register(Stack)
		handler.register(Create)
		handler.register(Install)


		# setup the application
		app.setup()

		app.run()
finally:
		app.close()
