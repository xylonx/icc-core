create table image(
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null,
    deleted_at timestamp with time zone,
    image_id text not null primary key,
    external_id text,
    md5_sum text not null,
    unique(md5_sum)
);
create index idx_image_deleted_at on image(deleted_at);
create table tag (
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null,
    deleted_at timestamp with time zone,
    id integer not null,
    tag_name_en text,
    tag_name_cn text,
    tag_name_jp text,
    primary key(id)
);
create index idx_tag_deleted_at on tag(deleted_at);
-- using logical constraint to speed up
create table image_tag_relation(
    image_id text not null,
    tag_id integer not null,
    primary key (image_id, tag_id)
);
create view rich_image as
select image.created_at,
    image.updated_at,
    image.deleted_at,
    image.image_id,
    external_id,
    md5_sum,
    array_remove(
        array_agg(
            DISTINCT tag_id
            order by tag_id
        ),
        NULL
    ) as tag_ids
from image
    left join image_tag_relation on image.image_id = image_tag_relation.image_id
group by image.image_id;
create table auth_token (
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null,
    deleted_at timestamp with time zone,
    token text not null primary key,
    uploading_bytes int8 not null default(0)
);
create index idx_auth_token_deleted_at on auth_token(deleted_at);
-- - for using tablesampling to select random rows
-- CREATE EXTENSION tsm_system_rows;