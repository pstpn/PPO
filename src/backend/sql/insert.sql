insert into company(id, name, city)
values ('1', 'Yandex', 'Moscow'),
       ('2', 'Tilt', 'SPB');

select *
from company;

insert into employee(id, phone_number, full_name, company_id, post, password, date_of_birth)
values ('1', '123', 'aaa', '1', 'Сотрудник', '123', '24.04.2004'),
       ('2', '321', 'bbb', '2', 'Сотрудник', '123', '24.04.2204'),
       ('3', '444', 'aaa', '1', 'Сотрудник СБ', '123', '21.04.2004');

select *
from employee;

insert into info_card(id, created_employee_id, is_confirmed, created_date)
values ('123', '1', false, '21.02.2003'),
       ('321', '2', false, '22.02.2003'),
       ('444', '3', true, '11.02.2003');

select *
from info_card;

insert into document(id, serial_number, info_card_id, type)
values ('123', '123', '123', 'Паспорт'),
       ('321', '321', '321', 'СНИЛС'),
       ('333', '333', '444', 'Паспорт');

select *
from document;

insert into photo(id, document_id, key)
values ('123', '123', 'owipeciuwc'),
       ('321', '321', 'edwedw'),
       ('333', '333', 'pokopj');

select *
from photo;

insert into field(id, document_id, type, value)
values ('123', '123', '1', 'test1'),
       ('321', '321', '0', 'test2'),
       ('333', '333', '1', 'test3');

select *
from field;

insert into checkpoint(id, phone_number)
values ('123', '123'),
       ('321', '321'),
       ('222', '222'),
       ('111', '111'),
       ('323', '323'),
       ('444', '444');

select *
from checkpoint;

insert into passage(id, checkpoint_id, document_id, type, time)
values ('123', '222', '123', 'Вход', now()),
       ('13', '222', '123', 'Выход', now()),
       ('1234', '321', '321', 'Вход', now()),
       ('1235', '111', '321', 'Выход', now()),
       ('1236', '222', '333', 'Вход', now()),
       ('1237', '222', '333', 'Выход', now()),
       ('1232', '222', '333', 'Вход', now());

select *
from passage;

insert into passage(id, checkpoint_id, document_id, type, time)
values ('1231213', '222', '333', 'Выход', now());
