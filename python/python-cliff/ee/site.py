import os
import sys
import logging
from cliff.app import App
from cliff.commandmanager import CommandManager
from cliff.command import Command
from cliff.show import ShowOne
from cliff.lister import Lister

class Site(Command):
    "A simple command that prints a message."

    log = logging.getLogger(__name__)

    def get_parser(self, prog_name):
        parser = super(Site, self).get_parser(prog_name)
        subparser1 = parser.add_subparsers(help='sub-command help')

        # add a subcommand
        parser_a = subparser1.add_parser('create', help='create help')
        #parser_a.add_argument('value', type=str, help='value help')


        # add a subcommand of a subcommand
        #subparser2 = parser_a.add_subparsers(help='child sub-command help')
        #parser_b = subparser2.add_parser('ipv4', help='ipv4 help')
        #parser_b.add_argument('ipval', type=str, help='value help')

        return parser

    def take_action(self, parsed_args):
        self.log.info('sending greeting')
        self.log.debug('debugging')
        self.app.stdout.write('Site create command called \n')

class Stack(Command):
    "Performs stack related activities"

    log = logging.getLogger(__name__)

    def get_parser(self, prog_name):
        parser = super(Stack, self).get_parser(prog_name)
        subparser1 = parser.add_subparsers(help='sub-command help')

        # add a subcommand
        parser_a = subparser1.add_parser('install', help='install help')
        #parser_b = subparser1.add_parser('remove', help='remove help')
        #parser_a.add_argument('value', type=str, help='value help')


        # add a subcommand of a subcommand
        #subparser2 = parser_a.add_subparsers(help='child sub-command help')
        #parser_b = subparser2.add_parser('ipv4', help='ipv4 help')
        #parser_b.add_argument('ipval', type=str, help='value help')

        return parser

    def take_action(self, parsed_args):
        self.log.info('ee stack command called')                                                                                                                        1,0-1         Top
        self.log.debug('debugging')
        self.app.stdout.write('Hi ee stack install\n')
