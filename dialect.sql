-- this are only here for ide autocompletion
create table Ban
(
    id                  int auto_increment
        primary key,
    game_ban_xuid       varchar(255)                           not null,
    game_ban_username   varchar(255)                           not null,
    game_ban_ip         varchar(255)                           not null,
    game_ban_device     varchar(255)                           not null,
    game_ban_region     varchar(255)                           not null,
    game_ban_model      varchar(255)                           not null,
    game_ban_controller varchar(255)                           not null,
    ban_author          varchar(255)                           null,
    ban_type            int          default 0                 null,
    ban_expire          timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    ban_reason          varchar(255) default 'No reason'       not null,
    unban               tinyint(1)   default 0                 not null,
    unban_author        varchar(255)                           null,
    unban_time          timestamp                              null
);

create table Bedwars
(
    id                        int auto_increment
        primary key,
    game_xuid                 varchar(255)  not null,
    bedwars_win_count         int default 0 not null,
    bedwars_losses_number     int default 0 not null,
    bedwars_win_streak        int default 0 not null,
    bedwars_bed_broken_number int default 0 not null,
    bedwars_final_kills       int default 0 not null,
    bedwars_kills             int default 0 not null,
    bedwars_ranked_points     int default 0 not null,
    bedwars_addon             longblob      null
);

create table Config
(
    id                   int auto_increment
        primary key,
    config_server_name   varchar(255)         not null,
    config_server_id     varchar(255)         not null,
    config_server_ip     varchar(255)         not null,
    config_server_port   int                  not null,
    config_server_motd   varchar(255)         not null,
    config_server_online tinyint(1) default 0 not null,
    config_server_update tinyint(1) default 0 not null
);

create table Global
(
    id                         int auto_increment
        primary key,
    game_username              varchar(255)                         not null,
    game_connect_on            text                                 null,
    game_coins                 int        default 50                not null,
    game_gems                  int        default 1                 not null,
    game_xuid                  varchar(255)                         not null,
    game_friends               longblob                             null,
    game_cosmetics             longblob                             null,
    game_skin_data             longblob                             null,
    game_cape_custom_data      longblob                             null,
    game_permissions           text                                 null,
    game_rank                  int                                  null,
    game_settings              longblob                             null,
    game_version               varchar(255)                         null,
    game_language              varchar(255)                         null,
    game_device_os             varchar(255)                         null,
    game_device_controller     varchar(255)                         null,
    game_device_id             varchar(255)                         null,
    game_device_model          text                                 null,
    game_device_ip             varchar(255)                         null,
    game_date_first_connection timestamp  default CURRENT_TIMESTAMP null,
    game_date_last_connection  timestamp                            null,
    game_on_save               tinyint(1) default 0                 null
);

create table Practice
(
    id                     int auto_increment
        primary key,
    game_xuid              varchar(255)  not null,
    practice_wins_count    int default 0 not null,
    practice_losses_number int default 0 not null,
    practice_elo           int default 0 not null,
    practice_kills         int default 0 not null,
    practice_deaths        int default 0 not null,
    practice_kits          longblob      null
);

create table `Rank`
(
    id              int auto_increment
        primary key,
    rank_name       varchar(255)         null,
    rank_permission varchar(255)         null,
    rank_default    tinyint(1) default 0 not null,
    rank_server     varchar(255)         not null
);

