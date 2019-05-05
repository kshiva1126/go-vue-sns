begin;

insert into messages(id, user_id, mention_id, body, created_at) values(
    1,
    1,
    NULL,
    "Hi, I am Adam! Nice to meecha!",
    NOW()
), (
    (select max(id) from messages m)+1,
    1,
    3,
    "I'm boring... What are u doing?",
    NOW()
), (
    (select max(id) from messages m)+1,
    2,
    NULL,
    "Hi, I am Bob!",
    NOW()
), (
    (select max(id) from messages m)+1,
    2,
    2,
    "I'M TWEETING NOW",
    NOW()
), (
    (select max(id) from messages m)+1,
    3,
    NULL,
    "Hi! I'm Choerry!",
    NOW()
), (
    (select max(id) from messages m)+1,
    3,
    2,
    "I practice singing and dancing!",
    NOW()
);

commit;
