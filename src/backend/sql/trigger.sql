create or replace function employee_passages_for_day(passage_document_id int, check_date date)
    returns table
            (
                document_id   int,
                entries_count bigint,
                exits_count   bigint
            ) as $$
begin
    if exists(
        select 1
        from passage p
        where
            p.document_id = passage_document_id
          and date(p.time) = check_date
        group by p.document_id
    ) then
        return query
            select
                p.document_id,
                count(case when p.type = 'Вход' then 1 end) as entries_count,
                count(case when p.type = 'Выход' then 1 end) as exits_count
            from passage p
            where
                p.document_id = passage_document_id
              and date(p.time) = check_date
            group by p.document_id;
    else
        return query
            select passage_document_id, 0::bigint, 0::bigint;
    end if;
end;
$$ language plpgsql;

create or replace function log_daily_employee_passages()
    returns trigger as
$$
declare
    document_id   int;
    entries_count bigint;
    exits_count   bigint;
begin
    select *
    into
        document_id,
        entries_count,
        exits_count
    from employee_passages_for_day(
            new.document_id,
            current_date
         );

    raise notice E'За % по документу с идентификатором % сотрудник: \n- вошел % раз(а)\n- вышел % раз(а)',
                current_date, document_id, entries_count, exits_count;
    return new;
end;
$$ language plpgsql;

create or replace trigger day_passages
    after insert on passage
    for each row
execute function log_daily_employee_passages();