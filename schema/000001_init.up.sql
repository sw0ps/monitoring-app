CREATE TABLE channels_stats
(
    id               serial       not null unique,
    channel_id       varchar(255) not null,
    title            varchar(255) not null,
    view_count       varchar(255) not null,
    subscriber_count varchar(255) not null,
    video_count      varchar(255) not null
);