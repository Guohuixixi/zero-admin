create table cms_topic_category
(
    id            bigint auto_increment
        primary key,
    name          varchar(100) not null,
    icon          varchar(500) not null comment '分类图标',
    subject_count int          not null comment '专题数量',
    show_status tinyint not null,
    sort          int          not null
)
    comment '话题分类表' charset = utf8;

