begin;

insert into messages(id, body, created_at) values (
    1,
    "Hello World",
    NOW()
), (
    (select max(id) from messages as m)+1,
    "Jaljayo World",
    NOW()
);

insert into comments(id, message_id, body, created_at) values (
    1,
    1,
    "Hey, what'up?",
    NOW()
), (
    (select max(id) from comments as c)+1,
    1,
    "Thanks, my bro",
    NOW()
), (
    (select max(id) from comments as c)+1,
    2,
    "Good night!",
    NOW()
), (
    (select max(id) from comments as c)+1,
    2,
    "I'm sleepy now too. Jaljayo.",
    NOW()
);

commit;
