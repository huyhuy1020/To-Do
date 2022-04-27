CREATE TABLE IF NOT EXISTS Employee(
    employee_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50),
);

CREATE TABLE IF NOT EXISTS Tasks(
    task_id     SERIAL  PRIMARY KEY,
    name_task   VARCHAR(50) NOT NULL,
);

CREATE TABLE IF NOT EXISTS Employee_Task(
    task_id     SERIAL,
    employee_id          SERIAL,
    PRIMARY KEY(task_id,id)
    CONSTRAINT FK_employee foreign key(employee_id) references Employee(employee_id)
    CONSTRAINT FK_task foreign key(task_id) references Employee(task_id)
)