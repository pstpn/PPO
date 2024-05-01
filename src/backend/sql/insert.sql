insert into company(name, city)
values ('Yandex', 'Moscow'),
       ('Tilt', 'SPB');

select *
from company;

insert into employee(phone_number, full_name, company_id, post, password, date_of_birth)
values ('123', 'aaa', 1, 'Сотрудник', '123', '24.04.2004'),
       ('321', 'bbb', 2, 'Сотрудник', '123', '24.04.2204'),
       ('444', 'aaa', 1, 'Сотрудник СБ', '123', '21.04.2004');

select *
from employee;

insert into info_card(created_employee_phone_number, is_confirmed, created_date)
values ('123', false, '21.02.2003'),
       ('321', false, '22.02.2003'),
       ('444', true, '11.02.2003'),
       ('444', true, '20.02.2003');

select *
from info_card;

insert into document(serial_number, info_card_id, type)
values ('123', 1, 'Паспорт'),
       ('321', 2, 'СНИЛС'),
       ('333', 2, 'Паспорт'),
       ('222', 3, 'Паспорт'),
       ('111123', 4, 'СНИЛС');

select *
from document;

insert into checkpoint(phone_number)
values ('123'),
       ('321'),
       ('222'),
       ('111'),
       ('323'),
       ('444');

select *
from checkpoint;

insert into passage(checkpoint_id, document_id, type, time)
values (1, 1, 'Вход', now()),
       (1, 1, 'Выход', now()),
       (2, 2, 'Вход', now()),
       (3, 2, 'Выход', now()),
       (1, 3, 'Вход', now()),
       (1, 3, 'Вход', now());

select *
from passage;

insert into passage(checkpoint_id, document_id, type, time)
values (1, 1, 'Выход', now());
