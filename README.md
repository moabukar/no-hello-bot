# Discord Bot Go

No-Hello-Bot is a Discord bot for Discord is an open-source bot that sends "no-hello" or gives you a prompt to give more context to your question!

This Discord bot is easy to set up and use. All you need to do is add it to your server and set up the words you'd like it to recognize. It will then let you know when it sees one of those words in chat and respond to it with a reply or a link to the no-hello site. 

## Inspiration

```sh
"Imagine calling someone on the phone, going hello! then putting them on hold... 🤦"

When asking for help (coding or something technical) from people generally or co-workers, it’s important to ask directly what you want rather than have 5 minutes of conversation by going back and forth with the question. This no-hello bot reminds people to be efficient with their commnication via chat.

For example, rather than saying:

"Hello, how are you? Can you help me with some code?” OR "Hello, who knows python here?” 

No! Too vague. The reply will possibly be “Hello, I’m fine, what is your issue exactly?”

So you see you just lost some productivity by making the person ask YOU what is the question. 

Instead, say something direct like:

“Hello, I’m running this Terraform script that creates some EC2 instance AND I’m getting permission issue errors, can anyone help?”

We get that you’re trying to be nice by doing this but chat/text is not like real-life conversations so it’s slower. You are actually making the person wait for you longer to ask the question. As a result, losing time and productivity. 

Don’t get me wrong, I’m not saying don’t ask your friends “how are you?” etc etc. There’s a time and place for that. But in such scenarios where you need help on some code, it makes it easier for others to answer your question.

```

## Getting Started

To get started with No-Hello-Bot, you'll need to add it to your Discord server.

1. Go to [this page](https://discord.com/api/oauth2/authorize?client_id=717596406498967564&permissions=8&scope=bot) and click "Add to Server".

2. Select the server you'd like to add the bot to and click "Authorize".

3. Once the bot has been added to your server, you can set up the words you'd like it to recognize. To do this, type `!hello-set <word>` in chat, where `<word>` is the word you'd like the bot to recognize. You can also type `!hello-set all` to have the bot respond to any word it sees.

Once you've set up the words you'd like the bot to recognize, it will start responding to them with fun gifs. Enjoy!

## How to run bot locally (with from app)

Pre-requisites:

- Make sure to have Golang installed.

1. Get your token ready into `docs/config.json`
2. `go run cmd/discord-bot-go.go`

## How to run bot locally (using containers - Docker)

Pre-requisites:

- Make sure to have Docker installed.

1. Build image using Dockerfile >> `docker build -t nohello Dockerfile`
2. Run container using `docker run nohello`

## Libraries

No Hello Bot for Discord was built using the Go libraries listed in go.mod

## Contributing

We welcome contributions to No-Hello-Bot! If you have an idea for a new feature or have found a bug, please open an issue or submit a pull request.
## License

No-Hello-Bot is licensed under the MIT license. See [LICENSE](LICENSE) for more details.

## TODO

- Add images/screenshots and more testing.
- Add a detailed step-by-step deployment guide
- Add support for slack too