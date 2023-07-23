## <p align="center">SkinAPI</p>
> ### Here is the optimized player's head image processing solution from Ectary Network!   
> It allows you to get the head of the player with the help of his XUID (Xbox) or his name !  

## Paths
| METHOD | PATH                          | RequireAuth | Description                                                                 | Params                      |
|--------|-------------------------------|-------------|-----------------------------------------------------------------------------|-----------------------------|
| GET    | /players/{name}/skin/head     | false       | get the head of a skin, you can chose the size of the output and the border | -xuid<br/>-size<br/>-border |
| GET    | /players/{name}/skin/full     | false       | get full skin of a player                                                   | -xuid                       |
| GET    | /players/{name}/info/full     | true        | get full player infos                                                       | -xuid                       |
| GET    | /players/{name}/info/bedwars  | true        | get player stats about bedwars                                              | -xuid                       |
| GET    | /players/{name}/info/practice | true        | get player stats about practice                                             | -xuid                       |
| GET    | /players/{name}/info          | true        | get base player infos                                                       | -xuid                       |
| GET    | /players/{name}/cape          | false       | get cape of the player                                                      | -xuid                       |
| GET    | /servers/{server}             | false       | get one/multiple server query result                                        |



## Params

| ParamName | Description                                                                                                 | Possible values                             | Default value |
|-----------|-------------------------------------------------------------------------------------------------------------|---------------------------------------------|---------------|
| name      | username/xuid of the player                                                                                 | any string                                  | No default    |
| xuid      | if param 'name' are a username or a xuid                                                                    | true, false                                 | false         |
| size      | resize the image result to specific size (0 = default size)                                                 | 0, 8, 16, 32, 64, 128, 256, 512, 1024, 2048 | 0             |
| border    | the size of the transparent border of the image                                                             | [0-100]                                     | 0             |
| server    | a address of one/multiple servers (if multiple, address are separate with %), port optional : default 19132 | any string                                  | No default    |


## Auth

for authorise first you need to add a token in token.json then when you make a request you have the choice between :
1. add header 'Authorization' with the token 
2. add param named 'token' with the token as value on the request


## Headers

- ARC-DefaultContent : if the content are a default content
- ARC-CacheDelay : delay of the cache

## Examples JSON Success Responses

GET `http://127.0.0.1:8000/players/MatroxMc/info/full`
```json
{
    "username": "MatroxMC",
    "coins": 50,
    "gems": 1,
    "xuid": "2535426682755720",
    "friends": null,
    "cosmetics": null,
    "rank": "VIP",
    "game_version": "1.18.2",
    "language": "fr_FR",
    "device_os": "Windows 10",
    "first_connection": "2021-12-30T18:40:21Z",
    "last_connection": "2021-12-30T18:40:21Z",
    "bedwars_stats": {
        "wins": 1,
        "losses": 2,
        "win_streak": 0,
        "broken_beds": 12,
        "kills": 0,
        "final_kills": 1,
        "ranked_points": 0
    },
    "practice_stats": {
        "wins": 0,
        "losses": 0,
        "elo": 100,
        "kills": 3,
        "deaths": 3
    }
}
```


GET `http://127.0.0.1:8000/players/MatroxMc/info/bedwars`
```json
{
    "wins": 1,
    "losses": 2,
    "win_streak": 0,
    "broken_beds": 12,
    "kills": 0,
    "final_kills": 1,
    "ranked_points": 0
}
```


GET `http://127.0.0.1:8000/players/MatroxMc/info/practice`
```json
{
    "wins": 0,
    "losses": 0,
    "elo": 100,
    "kills": 3,
    "deaths": 3
}
```


GET `http://127.0.0.1:8000/players/MatroxMc/info/`
```json
{
    "username": "MatroxMC",
    "coins": 50,
    "gems": 1,
    "xuid": "2535426682755720",
    "friends": null,
    "cosmetics": null,
    "rank": "VIP",
    "game_version": "1.18.2",
    "language": "fr_FR",
    "device_os": "Windows 10",
    "first_connection": "2021-12-30T18:40:21Z",
    "last_connection": "2021-12-30T18:40:21Z"
}
```

GET `http://localhost:8000/servers/ectary.club:19132%beta.ectary.club%unknow.ectary.club%ectary.club:19133`
```json
{
    "beta.ectary.club:19132": {
        "game_id": "MINECRAFTPE",
        "gametype": "SMP",
        "hostip": "0.0.0.0",
        "hostname": "�5�lEctary�r�e [New]",
        "hostport": "19132",
        "map": "WaterdogPE",
        "maxplayers": "1000",
        "numplayers": "0",
        "players": "",
        "plugins": "",
        "version": "",
        "whitelist": "off"
    },
    "ectary.club:19132": {
        "game_id": "MINECRAFTPE",
        "gametype": "SMP",
        "hostip": "0.0.0.0",
        "hostname": "§5§lEctary§r [1.18]",
        "hostport": "19132",
        "map": "lobby",
        "maxplayers": "1000",
        "numplayers": "78",
        "players": "Jl2u, oKqyc, Sugu1102, NOT dorito3380, Hattan Basferr",
        "plugins": "PocketMine-MP 3.28.0",
        "server_engine": "PocketMine-MP 3.28.0",
        "version": "v1.18.11",
        "whitelist": "off"
    },
    "ectary.club:19133": {
        "code": 11,
        "message": "someting whrong with query ectary.club:19133, error = read udp 192.168.235.55:61361->141.94.37.242:19133: i/o timeout",
        "status_code": 500
    },
    "unknow.ectary.club:19132": {
        "code": 11,
        "message": "someting whrong with query unknow.ectary.club:19132, error = error dialing UDP conn: dial udp: lookup unknow.ectary.club: no such host",
        "status_code": 500
    }
}
```

GET `http://localhost:8000/servers/ectary.club`
```json
{
    "game_id": "MINECRAFTPE",
    "gametype": "SMP",
    "hostip": "0.0.0.0",
    "hostname": "§5§lEctary§r [1.18]",
    "hostport": "19132",
    "map": "lobby",
    "maxplayers": "1000",
    "numplayers": "76",
    "players": "Jl2u, oKqyc, Sugu1102, Caleb4j",
    "plugins": "PocketMine-MP 3.28.0",
    "server_engine": "PocketMine-MP 3.28.0",
    "version": "v1.18.11",
    "whitelist": "off"
}
```

## Information's: 

| TYPE            | PROTOCOL      | VERSION |
|-----------------|---------------|---------|
| api             | HTTP          | 2.0.0   |

Make with ♥ by Arcoz
