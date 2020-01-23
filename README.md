Valera Telegram Bot
===

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-8%25-brightgreen.svg?longCache=true&style=flat)</a>

Valera is the simple telegram bot that do one dead simple thing:

> If Valera see any JIRA task key in the chat, he send the short information about this task to the chat.

Table of contents
===

- [Valera Telegram Bot](#valera-telegram-bot)
- [Table of contents](#table-of-contents)
- [How to set up](#how-to-set-up)
  - [Create telegram bot](#create-telegram-bot)
  - [Create the config file](#create-the-config-file)
  - [Run the bot](#run-the-bot)
    - [Run executable file](#run-executable-file)
    - [Run docker image](#run-docker-image)
  - [Use the bot](#use-the-bot)

# How to set up

## Create telegram bot

Use telegram @BotFather bot to create new bot. 

Important things:
* Go to bot settings –> group privacy -> and turn it off. If you want the bot to work in group chats.
* Copy bot `TOKEN`.

## Create the config file

You can see example file `configs/config.yaml`.

Please, use yaml format.

```
PROXY_URL: url:port
PROXY_USER: test_user
PROXY_PASS: test_pass
TELEGRAM_BOT_TOKEN: test_token
JIRA_USER: test_user
JIRA_PASS: test_pass
JIRA_URL: https://jira.example-company.ru
```

Proxy information is needed only if you have blocking problems in you country.

For now only basic auth in jira is supported so you have to provide login and password.

## Run the bot

### Run executable file

In the `builds` directory you can find builds for the different OS: macOS, windows and linux.

You can run the bot like that:

`./valera -config config.yaml -allowed wl.txt`

`-allowed` path will be used by the bot to store the information about white list chats, if you need that.

### Run docker image

For example, if you place `config.yaml` file to the path `/tmp/configs` that you can run image like that:

`docker run --name=valera -d -v /tmp/configs:/configs mehanizm/valera`

See logs

`docker logs valera`

## Use the bot

By default, bot answers nobody. To have a talk with the bot you have to add chat to the white list.

To do that read the logs. Copy secret string from the line:

`To add chat to the white list, please, send the string to the Bot: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX`

and send it to the chat you want to add to the white list.

If you want Valera to work in the group simply add him to the group and send secret string to the chat.

By default, after restart Valera will forget all information about the white list chats. If you want Valera to save the progress tell him `save all data`. After that Valera create file with chat ID in the `-allowed` path (or in the `config` directory if we talk about docker image).

