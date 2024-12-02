
select
    children.name,
    wish_lists.wishes->>'first_choice' as primary_wish,
    wish_lists.wishes->>'second_choice' as backup_wish,
    wish_lists.wishes#>>'{colors, 0}' as favorite_color,
    json_array_length(wish_lists.wishes->'colors') as color_count,
    case toy_catalogue.difficulty_to_make
        when 1 then 'Simple Gift'
        when 2 then 'Moderate Gift'
        else 'Complex Gift'
    end,
    case toy_catalogue.category
        when 'educational' then 'Learning Workshop'
        when 'outdoor' then 'Outside Workshop'
        else 'General Workshop'
    end
from children
join wish_lists
    on children.child_id = wish_lists.child_id
join toy_catalogue
    on wish_lists.wishes->>'first_choice' = toy_catalogue.toy_name
order by name
limit 5
;
