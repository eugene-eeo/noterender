from distutils.core import setup

setup(
    name='noterender',
    version='0.0.1',
    packages=['noterender'],
    package_data={'noterender': ['data/*.html']},
    scripts=['bin/noterender'],
    install_requires=[
        'docopt==0.6.2',
        'path.py==8.2.1',
        'markdown2==2.3.1',
        'chevron==0.9.0',
    ],
)
