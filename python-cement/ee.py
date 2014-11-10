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


class Create(controller.CementBaseController):
    """This controller commands are 'stacked' onto the site controller."""

    class Meta:
        label = 'create'
        interface = controller.IController
        stacked_on = 'site'
        stacked_type = 'nested'
        description = "create commands are listed under this controller"
        arguments = [
            (['site_name'], dict(help="domain name for your website like example.com")),
            ]

    @controller.expose(help="command under the base namespace")
    def default(self):
        print "Inside ee site create %s"%self.app.pargs.site_name



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
        description = "install commands  are listed under this controller"
        arguments = [
            (['--nginx'], dict(help="perform command on nginx package", action='store_true')),
            (['--php'], dict(help="perform command on php package", action='store_true')),
            (['--mysql'], dict(help="perform command on mysql package", action='store_true')),
            ]

    @controller.expose(help="default command for third_controller", hide=True)
    def default(self):
        print "Inside ee stack install"
        if self.app.pargs.nginx:
                print "Install nginx option selected"
        if self.app.pargs.php:
                print "Install php option selected"
        if self.app.pargs.mysql:
                print "Install mysql option selected"
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
