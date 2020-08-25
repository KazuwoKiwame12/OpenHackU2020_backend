CREATE table comments (
    id SERIAL NOT NULL,
    user_id NOT NULL,
    emotion_id NOT NULL,
    prefecture varchar(20) NOT NULL,
    latitude double precision NOT NULL,
    longtitude double precision NOT NULL,
    comment TEXT,
    date_time datetime NOT NULL,
    PRIMARY KEY (id), 
    foreign key (user_id) references users(id)
        on delete cascade
        on update no action
    foreign key (emotion_id) references emotions(id)
        on delete no action
        on update no action
);