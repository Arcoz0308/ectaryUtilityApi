-- name: select-ranks
SELECT id, rank_name, rank_default
FROM Rank;

-- name: select-skin-by-username
SELECT game_skin_data
FROM Global
WHERE game_username = ?;

-- name: select-skin-by-xuid
SELECT game_skin_data
FROM Global
WHERE game_xuid = ?;

-- name: select-cape-by-username
SELECT game_cape_custom_data
FROM Global
WHERE game_username = ?;

-- name: select-cape-by-xuid
SELECT game_cape_custom_data
FROM Global
WHERE game_xuid = ?;

-- name: select-player-info-full-by-username
SELECT game_username,
       game_coins,
       game_gems,
       G.game_xuid,
       game_friends,
       game_cosmetics,
       game_rank,
       game_version,
       game_language,
       game_device_os,
       game_date_first_connection,
       game_date_last_connection,
       bedwars_win_count,
       bedwars_losses_number,
       bedwars_win_streak,
       bedwars_bed_broken_number,
       bedwars_final_kills,
       bedwars_kills,
       bedwars_ranked_points,
       practice_wins_count,
       practice_losses_number,
       practice_elo,
       practice_kills,
       practice_deaths
FROM Global AS G
         INNER JOIN Bedwars B on G.game_xuid = B.game_xuid
         INNER JOIN Practice P on B.game_xuid = P.game_xuid
WHERE G.game_username = ?;

-- name: select-player-info-full-by-xuid
SELECT game_username,
       game_coins,
       game_gems,
       G.game_xuid,
       game_friends,
       game_cosmetics,
       game_rank,
       game_version,
       game_language,
       game_device_os,
       game_date_first_connection,
       game_date_last_connection,
       bedwars_win_count,
       bedwars_losses_number,
       bedwars_win_streak,
       bedwars_bed_broken_number,
       bedwars_final_kills,
       bedwars_kills,
       bedwars_ranked_points,
       practice_wins_count,
       practice_losses_number,
       practice_elo,
       practice_kills,
       practice_deaths
FROM Global AS G
         INNER JOIN Bedwars B on G.game_xuid = B.game_xuid
         INNER JOIN Practice P on B.game_xuid = P.game_xuid
WHERE G.game_xuid = ?;

-- name: select-player-info-by-username
SELECT game_username,
       game_coins,
       game_gems,
       game_xuid,
       game_friends,
       game_cosmetics,
       game_rank,
       game_version,
       game_language,
       game_device_os,
       game_date_first_connection,
       game_date_last_connection
FROM Global
WHERE game_username = ?;

-- name: select-player-info-by-xuid
SELECT game_username,
       game_coins,
       game_gems,
       game_xuid,
       game_friends,
       game_cosmetics,
       game_rank,
       game_version,
       game_language,
       game_device_os,
       game_date_first_connection,
       game_date_last_connection
FROM Global
WHERE game_xuid = ?;


-- name: select-player-info-bedwars-by-username
SELECT G.game_xuid,
       bedwars_win_count,
       bedwars_losses_number,
       bedwars_win_streak,
       bedwars_bed_broken_number,
       bedwars_final_kills,
       bedwars_kills,
       bedwars_ranked_points
FROM Global AS G
    INNER JOIN Bedwars B on G.game_xuid = B.game_xuid
WHERE G.game_username = ?;

-- name: select-player-info-bedwars-by-xuid
SELECT game_xuid,
       bedwars_win_count,
       bedwars_losses_number,
       bedwars_win_streak,
       bedwars_bed_broken_number,
       bedwars_final_kills,
       bedwars_kills,
       bedwars_ranked_points
FROM Bedwars
WHERE game_xuid = ?;

-- name: select-player-info-practice-by-username
SELECT G.game_xuid,
       practice_wins_count,
       practice_losses_number,
       practice_elo,
       practice_kills,
       practice_deaths
FROM Global AS G
    INNER JOIN Practice P on G.game_xuid = P.game_xuid
WHERE G.game_username = ?;

-- name: select-player-info-practice-by-xuid
SELECT game_xuid,
       practice_wins_count,
       practice_losses_number,
       practice_elo,
       practice_kills,
       practice_deaths
FROM Practice
WHERE game_xuid = ?;