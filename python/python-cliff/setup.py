PROJECT = 'ee'

# Change docs/sphinx/conf.py too!
VERSION = '0.1'

from setuptools import setup, find_packages

try:
    long_description = open('README.rst', 'rt').read()
except IOError:
    long_description = ''

setup(
    name=PROJECT,
    version=VERSION,

    description='Demo app for EasyEngine',
    long_description=long_description,

    author='Harshad Yeola',
    author_email='harshad.yeola@rtcamp.com',

    #url='https://github.com/dreamhost/cliff',
    #download_url='https://github.com/dreamhost/cliff/tarball/master',

    classifiers=['Development Status :: 3 - Alpha',
                 'License :: OSI Approved :: Apache Software License',
                 'Programming Language :: Python',
                 'Programming Language :: Python :: 2',
                 'Programming Language :: Python :: 2.7',
                 'Programming Language :: Python :: 3',
                 'Programming Language :: Python :: 3.2',
                 'Intended Audience :: Developers',
                 'Environment :: Console',
                 ],

    platforms=['Any'],

    scripts=[],

    provides=[],
    install_requires=['cliff'],

    namespace_packages=[],
    packages=find_packages(),

 include_package_data=True,

    entry_points={
        'console_scripts': [
            'ee = ee.main:main'
        ],
        'ee.commands': [
            'site = ee.site:Site',
            'stack = ee.site:Stack',
        ],
    },

    zip_safe=False,
)
