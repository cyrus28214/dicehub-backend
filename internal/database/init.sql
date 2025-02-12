drop table if exists "user";
create table "user" (
    id serial primary key,
    openid varchar(256) not null unique,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

drop table if exists "game";
create table "game" (
    id serial primary key,
    "name" varchar(100) not null,
    description text,
    image text,
    rating float,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

drop table if exists "tag";
create table "tag" (
    id serial primary key,
    "name" varchar(100) not null,
    description text,
    image text,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

drop table if exists "game_tag_relation";
create table "game_tag_relation" (
    game_id int,
    tag_id int,
    primary key (game_id, tag_id),
    foreign key (game_id) references game(id),
    foreign key (tag_id) references tag(id),
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

-- 插入游戏标签
insert into tag (name, description) values
    ('推理', '需要玩家运用逻辑思维，解开谜题的游戏类型'),
    ('社交', '强调玩家之间互动和交流的游戏类型'),
    ('聚会', '适合多人一起游玩的休闲游戏类型'),
    ('家庭', '适合家庭成员一起游玩的游戏类型'),
    ('卡牌', '以卡牌为主要游戏机制的类型'),
    ('德式', '以德国风格桌游为特色的游戏类型'),
    ('策略', '需要战略思维和战术规划的游戏类型'),
    ('资源管理', '需要合理规划和交易资源的游戏类型'),
    ('益智', '需要解开各种谜题的游戏类型'),
    ('解密', '需要通过推理和逻辑找出对手密码排列的益智游戏'),
    ('抽象', '需要通过放置地形卡牌建造中世纪城堡和道路的策略游戏'),
    ('毛线', '玩法简单，可以用来活跃气氛的游戏'),
    ('美式', '以美国风格桌游为特色的游戏类型'),
    ('剧情', '以故事情节为主要重点的游戏类型'),
    ('恐怖', '以恐怖氛围和惊悚元素为主的游戏类型'),
    ('反应', '需要快速反应和操作的游戏类型'),
    ('儿童', '适合儿童游玩的游戏类型'),
    ('中式', '以中国风格桌游为特色的游戏类型'),
    ('蒸蒸日上', '我们的游戏正在蒸蒸日上哦'),
    ('经济', '需要玩家通过商业竞争和投资获得胜利的游戏类型'),
    ('trpg', '需要玩家通过角色扮演和冒险的游戏类型'),
    ('角色扮演', '玩家需要扮演特定角色进行游戏的类型'),
    ('冒险', '以探索和冒险为主题的游戏类型'),
    ('战棋', '需要玩家通过战略和战术规划的游戏类型');

-- 插入游戏数据
insert into game (name, description, image, rating) values
    ('阿瓦隆', '正义与邪恶阵营对抗的社交推理游戏，通过投票完成任务决定胜负。', 'https://pic1.imgdb.cn/item/679b7bb9d0e0a243d4f8b0e4.jpg', 9.2),
    ('狼人杀', '狼人隐藏身份猎杀村民，村民通过推理找出狼人的社交游戏。', 'https://pic1.imgdb.cn/item/67ab9736d0e0a243d4fe75a5.jpg', 8.8),
    ('uno', '经典的卡牌配对游戏，通过出牌干扰对手，最先出完手牌获胜。', 'https://pic1.imgdb.cn/item/67ab9693d0e0a243d4fe7596.jpg', 8.5),
    ('卡坦岛', '在岛屿上收集资源、建设城市的策略游戏，需要合理规划和交易。', 'https://pic1.imgdb.cn/item/67ab9693d0e0a243d4fe7595.jpg', 9.3),
    ('达芬奇密码', '通过推理和逻辑找出对手密码排列的益智游戏。', 'https://pic1.imgdb.cn/item/67ab9735d0e0a243d4fe75a1.jpg', 8.6),
    ('卡卡颂', '通过放置地形卡牌建造中世纪城堡和道路的策略游戏。', 'https://pic1.imgdb.cn/item/67ab9736d0e0a243d4fe75a4.jpg', 9.0),
    ('谁是牛头人', '一款充满欢乐的吹牛与识破游戏，玩家需要在虚实之间找到平衡。', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75a8.jpg', 8.7),
    ('山屋惊魂', '在恐怖氛围中寻找线索，揭示真相的剧情推理游戏。', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75ab.jpg', 9.1),
    ('德国心脏病', '快节奏的图案匹配游戏，考验反应速度和观察力。', 'https://pic1.imgdb.cn/item/67ab9735d0e0a243d4fe75a3.jpg', 8.4),
    ('三国杀', '以三国为背景的卡牌对战游戏，玩家扮演不同的三国角色进行对抗。', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75aa.jpg', 8.9),
    ('怒海求生', '在沉船场景中合作逃生，随机事件带来紧张刺激的体验。', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75a9.jpg', 9.2),
    ('大富翁', '经典的房地产交易与投资游戏，体验商业竞争的乐趣。', 'https://pic1.imgdb.cn/item/67ab9735d0e0a243d4fe75a2.jpg', 8.5),
    ('龙与地下城', '在奇幻世界中冒险的角色扮演桌游，体验史诗般的故事。', 'https://pic1.imgdb.cn/item/67ab9693d0e0a243d4fe7594.jpg', 9.3),
    ('克苏鲁的呼唤', '在神秘的克苏鲁世界中探索，揭露隐藏的真相。', 'https://pic1.imgdb.cn/item/67ab9692d0e0a243d4fe7593.jpg', 8.7),
    ('战锤40k', '指挥遥远黑暗的四十个千年之后的军队，为人类二战。', 'https://pic1.imgdb.cn/item/67ab9692d0e0a243d4fe7592.jpg', 8.7);

-- 插入游戏和标签的关联关系
insert into game_tag_relation (game_id, tag_id) 
select g.id, t.id
from game g, tag t
where 
    (g.name = '阿瓦隆' and t.name in ('推理', '社交', '聚会')) or
    (g.name = '狼人杀' and t.name in ('推理', '社交', '聚会')) or
    (g.name = 'uno' and t.name in ('毛线', '家庭', '聚会', '卡牌')) or
    (g.name = '卡坦岛' and t.name in ('德式', '策略', '资源管理')) or
    (g.name = '达芬奇密码' and t.name in ('益智', '家庭', '解密')) or
    (g.name = '卡卡颂' and t.name in ('德式', '策略', '抽象')) or
    (g.name = '谁是牛头人' and t.name in ('毛线', '聚会', '社交')) or
    (g.name = '山屋惊魂' and t.name in ('美式', '剧情', '推理', '恐怖')) or
    (g.name = '德国心脏病' and t.name in ('毛线', '家庭', '反应', '儿童')) or
    (g.name = '三国杀' and t.name in ('中式', '卡牌', '策略', '蒸蒸日上')) or
    (g.name = '怒海求生' and t.name in ('美式', '合作', '剧情')) or
    (g.name = '大富翁' and t.name in ('家庭', '策略', '经济')) or
    (g.name = '龙与地下城' and t.name in ('trpg', '角色扮演', '冒险')) or
    (g.name = '克苏鲁的呼唤' and t.name in ('trpg', '角色扮演', '冒险')) or
    (g.name = '战锤40k' and t.name in ('战棋', '策略', '角色扮演'));

