# Mai Configuration

Mai stores a config file in JSON format on your disk. By default, it's stored in location `$USER_HOME/.mai/config.json` but you can specify a different location by setting an environment variable `MAI_CONFIG_DIR`.

## Initial Config
After installing Mai, you can configure it by running `mai config init` to generate an empty config. If a config already exists, this command will do nothing. You have to run `mai config nuke` to start fresh.

```
$ mai config init
```

Before you can do anything useful, you need to add at least a model to the config by running `mai config add-model` command.

## Add Model
You can add a model either by editing the `config.json` manually (which is faster if you know what you are doing) or use `mai config add-model`. This command interactively walks you through entering all required information to make sure the added model will really work. Your chosen name of the model needs to be unique among all specified models.

```
$ mai config add-model

Select your model provider (openai): openai
Select the name of the provider's model (gpt-3.5-turbo, gpt-4): gpt-4
Choose a name for your model: codegen
Enter the openai API key for accessing this model: XXXXXXXXXXX..

Trying a test prompt....
Prompt successful! You are good to go.
```

## Default Model
Mai needs a default model to run many of the "fast" operations. If no default is specified, we assume the first model from the config as the default. If you want to change that, you can set default as follows:

```
$ mai config set-default --model codegen
```

## Nuke Config
If you wish to cleanup all Mai config from your machine, you can run `mai config nuke` command.

## Other Operations
Other operations against the config such as updating and deleting models are not supported through the CLI at this point. You can simply edit the config file for that.