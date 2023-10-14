import discord
from discord.ext import commands

# Remplacez ces valeurs par les vôtres
TOKEN = 'YOUR_BOT_TOKEN'
id_owner = 'YOUR_OWNER_ID'
id_admin = 'YOUR_ADMIN_ROLE_ID'

bot = commands.Bot(command_prefix='+')


@bot.event
async def on_ready():
    print(f'Logged in as {bot.user.name}')


@bot.command()
async def admin(ctx, action, member: discord.Member):
    if str(ctx.author.id) == id_owner:
        if action == 'add':
            role = discord.utils.get(ctx.guild.roles, id=id_admin)
            await member.add_roles(role)
            await ctx.send(f'Le rôle {role.name} a été ajouté à {member.mention}')
        elif action == 'rm':
            role = discord.utils.get(ctx.guild.roles, id=id_admin)
            await member.remove_roles(role)
            await ctx.send(f'Le rôle {role.name} a été enlevé à {member.mention}')
        else:
            await ctx.send('Utilisation incorrecte. Utilisez `+admin add` ou `+admin rm`')
    else:
        await ctx.send('Vous n\'avez pas la permission d\'utiliser cette commande.')


bot.run(TOKEN)
