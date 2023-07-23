import asyncio, discord, pathlib, json
from discord.ext import commands


bot = commands.Bot(command_prefix = "/", help_command = None, intents = discord.Intents.all())

with open("./config.json", "r", encoding = "UTF-8") as file:
	config = json.load(file)


@bot.event
async def on_ready():
	commands_dir = pathlib.Path(__file__).parent / "commands"
	for command_file in commands_dir.glob("*py"):
		if command_file != "__init__.py":
			await bot.load_extension(f"commands.{command_file.name[:-3]}")

	events_dir = pathlib.Path(__file__).parent / "events"
	for event_file in events_dir.glob("*py"):
		if event_file.name != "__init__.py" and event_file.name != "on_ready.py":
			await bot.load_extension(f"events.{event_file.name[:-3]}")
	try:
		synced = await bot.tree.sync()
		print(f"Synced {len(synced)} commands !")
	except Exception:
		print(Exception)

asyncio.run(bot.load_extension("events.on_ready"))

bot.run(config["other"]["token"])
