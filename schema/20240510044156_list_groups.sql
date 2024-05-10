-- +goose Up
-- +goose StatementBegin
create table list_group
(
    list_group_id TEXT PRIMARY KEY,
    list_id       text      not null,
    name          text      not null,
    created_at    timestamp not null,
    updated_at    timestamp not null,

    foreign key (list_id) references list (list_id)
);

create table item_group
(
    item_group_id TEXT      not null,
    list_item_id  text      not null,
    created_at    timestamp not null,
    updated_at    timestamp not null,

    primary key (item_group_id, list_item_id),
    foreign key (item_group_id) references list_group (list_group_id),
    foreign key (list_item_id) references list_item (list_item_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table list_group;
drop table item_group;
-- +goose StatementEnd
