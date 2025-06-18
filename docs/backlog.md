# 요구사항 짬뽕 docs

## user

### u_user
- uid
- cert
- password
- registered_datetime
- last_login_datetime


### u_user detail
- email
- name
- birthday
- user_type (admin, user, creator)

<br>

## know

### kn_forest
- uid
- knowledge_title
- owner_uid (u_user_uid foreign key)
- registered_datetime

### kn_tree
- uid
- forest_uid ( foreign key forest uid )
- comment_uid (foreign key comment uid)
- registered_datetime

### kn_comment
- uid
- tree_uid
- user_uid
- content
- registered_datetime


<br>
