# Telegram Meme Bot

![Static Badge](https://img.shields.io/badge/golang-00ADD8?&style=plastic&logo=go&logoColor=white)


![Imgur](https://i.imgur.com/afqGtXf.jpg)

__This is a simple bot that can send you funny stories in a Telegram channel from the perspective of animals in the guise of very vivid characters from TV series. This is my first project where I used HTTP requests and API.I've added the Grok 4.1 fast model to it; if you want, you can try any other model — just change the model name.The API token works for all models.__

- [Api token](https://openrouter.ai/)

## What does the bot do?

**It sends text to the Telegram channel strictly on a timer, every 2 hours.
You need to add your bot's chat ID, API token, and the AI token**

- [Example](https://github.com/Fudziyosha/bot_meme_grok/blob/master/.env.example)

*You can also change the sending time. This is done here:* 
-
```go
_, errCron := c.AddFunc("*/120 * * * *", func()
```
**Where you need to change 120 minutes to any. But keep in mind that the free version assumes 50 requests per day. Keep in mind**
