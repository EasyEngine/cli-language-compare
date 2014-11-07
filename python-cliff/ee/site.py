import logging

from cliff.command import Command


class Site(Command):
    "site command manages site related functions"

    log = logging.getLogger(__name__)

    def take_action(self, parsed_args):
        self.log.info('ee site command called')
        self.log.debug('debugging')
        self.app.stdout.write('ee stack command called!\n')


class Stack(Command):
    "Performs stack related activities"

    log = logging.getLogger(__name__)

    def take_action(self, parsed_args):
        self.log.info('ee stack command called')
        self.log.debug('debugging')
        self.app.stdout.write('higfadsggdfgdfgdfshgdhsdhdh!\n')
