with wishes as (
    select
        *,
        row_number() over (partition by child_id order by wish_lists.submitted_date desc) as rank
    from wish_lists
)
select
    children.name,
    wishes.wishes->>'first_choice' as primary_wish,
    wishes.wishes->>'second_choice' as backup_wish,
    wishes.wishes#>>'{colors, 0}' as favorite_color,
    json_array_length(wishes.wishes->'colors') as color_count,
    case toy_catalogue.difficulty_to_make
        when 1 then 'Simple Gift'
        when 2 then 'Moderate Gift'
        else 'Complexe Gift'
    end,
    case toy_catalogue.category
        when 'educational' then 'Learning Workshop'
        when 'outdoor' then 'Outside Workshop'
        else 'General Workshop'
    end
from children
join wishes
    on children.child_id = wishes.child_id
    and wishes.rank = 1
join toy_catalogue
    on wishes.wishes->>'first_choice' = toy_catalogue.toy_name
order by name
limit 5
;
