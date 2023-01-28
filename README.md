# Simple Telegram bot

Base template of Telegram bot. Simple catalog of items.

### Supported methods

- Add `/addprod Name_Description of item_9.99` returns id
- Get `/get <id>` returns item
- Edit `/editprod <id>_NewName_NewDescription_7.99` returns oldItem and newItem
- Delete `/delprod <id>` returns deletedItem (other products id will be shifted up)
- Find `/findprod foo` returns all items with Title contains "foo"
- Help `/help` returns list of commands available
- List `/list` returns list of items using pagination

### Item struct:
```
  Title       string
  Description string
  Price       float64
```

## Usage

Put your Telegram bot TOKEN in .env `TOKEN="1234567890:AaSsDd..."` or enter it after starting app: `input telegram bot unique token:`

To make new TOKEN for new bot find @BotFather in Telegram `/newbot`


### ✨Technical features

Easy to add new command.

Just add new myNew.go file in package "commands" with code contains init() and method will be operates next start
```
func init() {
	registeredCommands["myNew"] = (*Commander).MyNew
}


func (c *Commander) MyNew(inputMessage *tgbotapi.Message) {
...
}
```
