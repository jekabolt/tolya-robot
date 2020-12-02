Tolya-robot [![Build Status](https://api.travis-ci.org/jekabolt/tolya-robot.svg?branch=master)](https://travis-ci.org/jekabolt/tolya-robot)

Goals
-----

* telegram bot for variable clothes sale subscription
* front-end for convenient submitting user data   

Environment variables
--------------
SERVER_PORT = default 8080  
TELEGRAM_BOT_TOKEN = default kektoken  

Using tolya-robot
--------------

``make local`` - replace all links to localhost  
``make dot`` - replace all links to dotmarket.io     
``make build`` - build for local test   
``make run`` - run locally  
``make image`` - build and tag docker image  

deploy via travis ci and webhook on pushed image   

TODO:  
How to setup deploy  