drop table if exists company CASCADE;
drop table if exists company_posts CASCADE;
drop table if exists document CASCADE;
drop table if exists employee CASCADE;
drop table if exists field CASCADE;
drop table if exists info_card CASCADE;
drop table if exists photo CASCADE;
drop table if exists checkpoint CASCADE;
drop table if exists passage CASCADE;

drop function if exists employee_passages_for_day(passage_document_id integer, check_date date);
drop function if exists employee_passages_for_day(passage_document_id text, check_date date);
drop function if exists log_daily_employee_passages();
drop trigger if exists day_passages on passage;