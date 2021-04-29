import sys

# Get from https://my.telegram.org/apps.
API_ID, API_HASH = 123, "CHANGE_ME"


def main() -> None:
    from telethon.sync import TelegramClient, functions

    client = TelegramClient("set online", api_id=API_ID, api_hash=API_HASH)
    client.start()

    set_online = functions.account.UpdateStatusRequest(offline=False)
    client(set_online)


try:
    main()
except KeyboardInterrupt:
    sys.exit(2)
