-- auto-generated definition
create table crm360_person
(
    id                              bigserial
        primary key,
    first_name                      jsonb,
    last_name                       jsonb,
    middle_name                     jsonb,
    iin                             varchar,
    rnn                             varchar,
    birth_date                      date,
    death_date                      date,
    resident                        varchar,
    phone_number                    varchar,
    asp                             jsonb,
    esp                             jsonb,
    accident                        jsonb,
    avg_income                      jsonb,
    avg_pension_income              jsonb,
    avg_social_income               jsonb,
    birth_info                      jsonb,
    citizenship                     jsonb,
    education                       jsonb,
    health                          jsonb,
    farm_animal                     jsonb,
    scoring                         jsonb,
    financing_terr_extr_list        jsonb,
    opv                             jsonb,
    gender                          jsonb,
    nationality                     jsonb,
    income                          jsonb,
    job                             jsonb,
    income_refund                   jsonb,
    criminal_record                 jsonb,
    kdn                             jsonb,
    tax_notification                jsonb,
    narco_registry                  jsonb,
    photo                           jsonb,
    bmg                             jsonb,
    psycho_registry                 jsonb,
    registration_address            jsonb,
    tub_registry                    jsonb,
    vaccination_kz                  jsonb,
    real_estate_object_registration jsonb,
    portal_refresh_date             date,
    wanted                          jsonb,
    real_estate_queue               jsonb,
    real_estate_object_encumbrance  jsonb,
    bankruptcy_application          jsonb,
    debtor_mu                       jsonb,
    unemployed_reg                  varchar,
    unemployed_reg_refresh_date     date
);

alter table crm360_person
    owner to root;

create index idx_iin
    on crm360_person (iin);

