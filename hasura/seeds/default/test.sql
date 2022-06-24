insert into users (id, name) values (1, 'test');
insert into labels 
  (id, name) 
values 
  (1, 'label1'),
  (2, 'label2'),
  (3, 'label3'),
  (4, 'label4'),
  (5, 'label5'),
  (6, 'label6'),
  (7, 'label7'),
  (8, 'label8'),
  (9, 'label9');
insert into priorities 
  (id, name) 
values 
  (1, '高'),
  (2, '中'),
  (3, '低');
insert into statuses 
  (id, name) 
values 
  (1, '未完'),
  (2, '実行中'),
  (3, '完了');
insert into todos (id, title,description, user_id, status_id, priority_id) values (1, 'test','test', 1, 1, 1);
insert into todos_labels_relations
  (id, todo_id, label_id) 
values 
  (1, 1, 1),
  (2, 1, 2);
