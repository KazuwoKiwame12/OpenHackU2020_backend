CREATE table responses (
    id SERIAL NOT NULL,
    user_id NOT NULL,
    comment_id NOT NULL,
    comment TEXT,
    date_time datetime NOT NULL,
    PRIMARY KEY (id), 
    foreign key (user_id) references users(id)
        on delete cascade
        on update no action
    foreign key (comment_id) references comments(id)
        on delete no action
        on update no action
);