begin;

insert into users(id, name, profile, email, created_at) values(
    1,
    "Adam",
    "Hello, I'm Adam.",
    "adam@example.com",
    NOW()
), (
    (select max(id) from users u)+1,
    "Bob",
    "Hello, I'm Bob.",
    "bob@example.com",
    NOW()
), (
    (select max(id) from users u)+1,
    "Choerry",
    "Hello, I'm Choerry From LOONA.",
    "choerry@example.com",
    NOW()
);

commit;
