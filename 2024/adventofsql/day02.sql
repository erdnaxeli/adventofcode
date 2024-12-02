select string_agg(char, '') from (
    select id, chr(value) as char
    from letters_a
    where
        value between 65 and 90
        or value between 97 and 122
        or value in (20, 21, 27, 28, 29, 44, 45, 46, 58, 59, 63)

    union

    select (select max(id) from letters_a) + id, chr(value) as char
    from letters_b
    where
        value between 65 and 90
        or value between 97 and 122
        or value in (21, 27, 28, 29, 32, 33, 44, 45, 46, 58, 59, 63)

    order by id
) as letters
;
