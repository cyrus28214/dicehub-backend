drop table if exists "user" cascade;
create table "user" (
    id serial primary key,
    openid varchar(256) not null unique,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

drop table if exists "game" cascade;
create table "game" (
    id serial primary key,
    "name" varchar(100) not null,
    description text,
    image text,
    rating float,
    likes_count int default 0,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

drop table if exists "tag" cascade;
create table "tag" (
    id serial primary key,
    "name" varchar(100) not null,
    description text,
    image text,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

drop table if exists "game_tag_relation" cascade;
create table "game_tag_relation" (
    game_id int,
    tag_id int,
    primary key (game_id, tag_id),
    foreign key (game_id) references game(id),
    foreign key (tag_id) references tag(id),
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

drop table if exists "like" cascade;
create table "like" (
    user_id int not null,
    game_id int not null,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp,
    primary key (user_id, game_id),
    foreign key (user_id) references "user"(id),
    foreign key (game_id) references game(id)
);

-- 插入游戏标签
insert into tag (id, name, description) values
    (1, '推理', '蛛丝马迹间的真相博弈，下一个福尔摩斯就是你'),
    (2, '社交', '打破社交距离的快乐结界'),
    (3, '聚会', '制造欢笑的破冰神器'),
    (4, '家庭', '三代人共享的温馨时光制造机'),
    (5, '卡牌', '方寸之间扭转乾坤的战略艺术'),
    (6, '德式', '精密如钟表的策略交响曲'),
    (7, '策略', '需要战略思维和战术规划的游戏类型'),
    (8, '资源管理', '在稀缺中缔造丰饶'),
    (9, '益智', '需要解开各种谜题的游戏类型'),
    (10, '解谜', '与设计者脑洞的终极对决'),
    (11, '抽象', '极简的思维角斗场'),
    (12, '毛线', '三分钟笑出腹肌的快乐法宝'),
    (13, '美式', '肾上腺素与骰子齐飞的狂想曲'),
    (14, '剧情', '亲手改写命运的叙事盛宴'),
    (15, '恐怖', '以恐怖氛围和惊悚元素为主的游戏类型'),
    (16, '反应', '手脑协调极限挑战赛'),
    (17, '儿童', '熊孩子能量转化装置'),
    (18, '中式', '东方智慧与美学的现代重构'),
    (19, '蒸蒸日上', '我们的游戏正在蒸蒸日上哦'),
    (20, '经济', '华尔街之狼速成模拟器'),
    (21, 'trpg', '第二人生体验卡已激活'),
    (22, '角色扮演', '解锁人生第N种可能的平行世界'),
    (23, '冒险', '在客厅完成史诗远征的魔法'),
    (24, '战棋', '沙盘上的百万雄兵'),
    (25, '新手', '推开桌游世界的大门');

-- 插入游戏数据
insert into game (name, description, image, rating, likes_count) values
    ('阿瓦隆', '圆桌骑士与邪恶爪牙的终极博弈，每一次眼神交汇都暗藏玄机的身份谜局', 'https://pic1.imgdb.cn/item/679b7bb9d0e0a243d4f8b0e4.jpg', 9.2, 2348),
    ('狼人杀', '月夜狼嚎与村民智慧的生死博弈，用谎言编织真相的社交修罗场', 'https://pic1.imgdb.cn/item/67ab9736d0e0a243d4fe75a5.jpg', 8.8, 7623),
    ('uno', '彩虹色卡牌旋风，反转+4的魔法时刻让聚会秒变尖叫现场', 'https://pic1.imgdb.cn/item/67ab9693d0e0a243d4fe7596.jpg', 8.5, 4521),
    ('卡坦岛', '拓荒者的经济学盛宴，用羊毛矿石搭建属于你的海上贸易帝国', 'https://pic1.imgdb.cn/item/67ab9693d0e0a243d4fe7595.jpg', 9.3, 3876),
    ('达芬奇密码', '数字矩阵中的头脑风暴，用排除法破解对手的密码铠甲', 'https://pic1.imgdb.cn/item/67ab9735d0e0a243d4fe75a1.jpg', 8.6, 1892),
    ('卡卡颂', '中世纪版图拼图大师，用城墙与修道院绘制法兰西风情画卷', 'https://pic1.imgdb.cn/item/67ab9736d0e0a243d4fe75a4.jpg', 9.0, 2967),
    ('谁是牛头人', '真话假话大乱斗，在夸张表演中揪出说谎的米诺陶洛斯', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75a8.jpg', 8.7, 1543),
    ('山屋惊魂', '古宅幽深走廊中的诅咒谜团，每一次骰子滚动都在叩响命运之门', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75ab.jpg', 9.1, 986),
    ('德国心脏病', '水果警报器狂响时刻，手速与眼力的终极试炼场', 'https://pic1.imgdb.cn/item/67ab9735d0e0a243d4fe75a3.jpg', 8.4, 3254),
    ('三国杀', '青梅煮酒论英雄，锦囊妙计定乾坤的东方权谋盛宴', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75aa.jpg', 8.9, 8965),
    ('怒海求生', '惊涛骇浪中的生存博弈，道德与利益的暴风抉择', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75a9.jpg', 9.2, 1234),
    ('大富翁', '地产大亨的财富狂想曲，用骰子丈量你的商业版图', 'https://pic1.imgdb.cn/item/67ab9735d0e0a243d4fe75a2.jpg', 8.5, 4567),
    ('龙与地下城', '剑与魔法的史诗旅程，每一个骰点都在书写你的英雄传说', 'https://pic1.imgdb.cn/item/67ab9693d0e0a243d4fe7594.jpg', 9.3, 3421),
    ('克苏鲁的呼唤', '直面不可名状的恐惧，在疯狂边缘探寻禁忌真相', 'https://pic1.imgdb.cn/item/67ab9692d0e0a243d4fe7593.jpg', 8.7, 2789),
    ('战锤40k', '银河战火永不熄灭，用战术棋子演绎星辰大海的征服史诗', 'https://pic1.imgdb.cn/item/67ab9692d0e0a243d4fe7592.jpg', 8.7, 1678);

-- 插入游戏和标签的关联关系
insert into game_tag_relation (game_id, tag_id) 
select g.id, t.id
from game g, tag t
where 
    (g.name = '阿瓦隆' and t.name in ('推理', '社交', '聚会')) or
    (g.name = '狼人杀' and t.name in ('推理', '社交', '聚会')) or
    (g.name = 'uno' and t.name in ('毛线', '家庭', '聚会', '卡牌', '新手')) or
    (g.name = '卡坦岛' and t.name in ('德式', '策略', '资源管理')) or
    (g.name = '达芬奇密码' and t.name in ('益智', '家庭', '解谜')) or
    (g.name = '卡卡颂' and t.name in ('德式', '策略', '抽象')) or
    (g.name = '谁是牛头人' and t.name in ('毛线', '聚会', '社交', '新手')) or
    (g.name = '山屋惊魂' and t.name in ('美式', '剧情', '推理', '恐怖')) or
    (g.name = '德国心脏病' and t.name in ('毛线', '家庭', '反应', '儿童')) or
    (g.name = '三国杀' and t.name in ('中式', '卡牌', '策略', '蒸蒸日上')) or
    (g.name = '怒海求生' and t.name in ('美式', '合作', '剧情')) or
    (g.name = '大富翁' and t.name in ('家庭', '策略', '经济', '新手')) or
    (g.name = '龙与地下城' and t.name in ('trpg', '角色扮演', '冒险')) or
    (g.name = '克苏鲁的呼唤' and t.name in ('trpg', '角色扮演', '冒险')) or
    (g.name = '战锤40k' and t.name in ('战棋', '策略', '角色扮演'));

