from setuptools import setup


setup(
		name='ee',
		version='0.1.0',
		py_modules=['ee'],
		install_requires=[
			'Click',
		],
		entry_points='''
			[console_scripts]
			ee=ee:cli
		''',
	)

