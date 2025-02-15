drop table if exists "user" cascade;
create table "user" (
    id serial primary key,
    openid varchar(256) not null unique,
    name varchar(100),
    avatar text,
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
    extra_info jsonb,
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

-- 添加评论表
drop table if exists "comment" cascade;
create table "comment" (
    id serial primary key,
    user_id int not null,
    game_id int not null,
    content text not null,
    rating float not null check (rating >= 0 and rating <= 10),
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp,
    foreign key (user_id) references "user"(id),
    foreign key (game_id) references game(id)
);

-- 添加触发器更新游戏的平均评分
create or replace function update_game_rating()
returns trigger as $$
begin
    update game
    set rating = (
        select avg(rating)
        from "comment"
        where game_id = new.game_id
    )
    where id = new.game_id;
    return new;
end;
$$ language plpgsql;

create trigger update_game_rating_trigger
after insert or update or delete on "comment"
for each row execute function update_game_rating();

-- -- 插入游戏标签
-- insert into tag (id, name, description) values
--     (1, '推理', '蛛丝马迹间的真相博弈，下一个福尔摩斯就是你'),
--     (2, '社交', '打破社交距离的快乐结界'),
--     (3, '聚会', '制造欢笑的破冰神器'),
--     (4, '家庭', '三代人共享的温馨时光制造机'),
--     (5, '卡牌', '方寸之间扭转乾坤的战略艺术'),
--     (6, '德式', '精密如钟表的策略交响曲'),
--     (7, '策略', '需要战略思维和战术规划的游戏类型'),
--     (8, '资源管理', '在稀缺中缔造丰饶'),
--     (9, '益智', '需要解开各种谜题的游戏类型'),
--     (10, '解谜', '与设计者脑洞的终极对决'),
--     (11, '抽象', '极简的思维角斗场'),
--     (12, '毛线', '三分钟笑出腹肌的快乐法宝'),
--     (13, '美式', '肾上腺素与骰子齐飞的狂想曲'),
--     (14, '剧情', '亲手改写命运的叙事盛宴'),
--     (15, '恐怖', '以恐怖氛围和惊悚元素为主的游戏类型'),
--     (16, '反应', '手脑协调极限挑战赛'),
--     (17, '儿童', '熊孩子能量转化装置'),
--     (18, '中式', '东方智慧与美学的现代重构'),
--     (19, '蒸蒸日上', '我们的游戏正在蒸蒸日上哦'),
--     (20, '经济', '华尔街之狼速成模拟器'),
--     (21, 'trpg', '第二人生体验卡已激活'),
--     (22, '角色扮演', '解锁人生第N种可能的平行世界'),
--     (23, '冒险', '在客厅完成史诗远征的魔法'),
--     (24, '战棋', '沙盘上的百万雄兵'),
--     (25, '新手', '推开桌游世界的大门');

-- -- 插入游戏数据
-- -- insert into game (name, description, image, rating, likes_count) values
-- --     ('阿瓦隆', '圆桌骑士与邪恶爪牙的终极博弈，每一次眼神交汇都暗藏玄机的身份谜局', 'https://pic1.imgdb.cn/item/679b7bb9d0e0a243d4f8b0e4.jpg', 9.2, 2348),
-- --     ('狼人杀', '月夜狼嚎与村民智慧的生死博弈，用谎言编织真相的社交修罗场', 'https://pic1.imgdb.cn/item/67ab9736d0e0a243d4fe75a5.jpg', 8.8, 7623),
-- --     ('uno', '彩虹色卡牌旋风，反转+4的魔法时刻让聚会秒变尖叫现场', 'https://pic1.imgdb.cn/item/67ab9693d0e0a243d4fe7596.jpg', 8.5, 4521),
-- --     ('卡坦岛', '拓荒者的经济学盛宴，用羊毛矿石搭建属于你的海上贸易帝国', 'https://pic1.imgdb.cn/item/67ab9693d0e0a243d4fe7595.jpg', 9.3, 3876),
-- --     ('达芬奇密码', '数字矩阵中的头脑风暴，用排除法破解对手的密码铠甲', 'https://pic1.imgdb.cn/item/67ab9735d0e0a243d4fe75a1.jpg', 8.6, 1892),
-- --     ('卡卡颂', '中世纪版图拼图大师，用城墙与修道院绘制法兰西风情画卷', 'https://pic1.imgdb.cn/item/67ab9736d0e0a243d4fe75a4.jpg', 9.0, 2967),
-- --     ('谁是牛头人', '真话假话大乱斗，在夸张表演中揪出说谎的米诺陶洛斯', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75a8.jpg', 8.7, 1543),
-- --     ('山屋惊魂', '古宅幽深走廊中的诅咒谜团，每一次骰子滚动都在叩响命运之门', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75ab.jpg', 9.1, 986),
-- --     ('德国心脏病', '水果警报器狂响时刻，手速与眼力的终极试炼场', 'https://pic1.imgdb.cn/item/67ab9735d0e0a243d4fe75a3.jpg', 8.4, 3254),
-- --     ('三国杀', '青梅煮酒论英雄，锦囊妙计定乾坤的东方权谋盛宴', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75aa.jpg', 8.9, 8965),
-- --     ('怒海求生', '惊涛骇浪中的生存博弈，道德与利益的暴风抉择', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75a9.jpg', 9.2, 1234),
-- --     ('大富翁', '地产大亨的财富狂想曲，用骰子丈量你的商业版图', 'https://pic1.imgdb.cn/item/67ab9735d0e0a243d4fe75a2.jpg', 8.5, 4567),
-- --     ('龙与地下城', '剑与魔法的史诗旅程，每一个骰点都在书写你的英雄传说', 'https://pic1.imgdb.cn/item/67ab9693d0e0a243d4fe7594.jpg', 9.3, 3421),
-- --     ('克苏鲁的呼唤', '直面不可名状的恐惧，在疯狂边缘探寻禁忌真相', 'https://pic1.imgdb.cn/item/67ab9692d0e0a243d4fe7593.jpg', 8.7, 2789),
-- --     ('战锤40k', '银河战火永不熄灭，用战术棋子演绎星辰大海的征服史诗', 'https://pic1.imgdb.cn/item/67ab9692d0e0a243d4fe7592.jpg', 8.7, 1678);

-- -- 插入游戏和标签的关联关系
-- insert into game_tag_relation (game_id, tag_id) 
-- select g.id, t.id
-- from game g, tag t
-- where 
--     (g.name = '阿瓦隆' and t.name in ('推理', '社交', '聚会')) or
--     (g.name = '狼人杀' and t.name in ('推理', '社交', '聚会')) or
--     (g.name = 'uno' and t.name in ('毛线', '家庭', '聚会', '卡牌', '新手')) or
--     (g.name = '卡坦岛' and t.name in ('德式', '策略', '资源管理')) or
--     (g.name = '达芬奇密码' and t.name in ('益智', '家庭', '解谜')) or
--     (g.name = '卡卡颂' and t.name in ('德式', '策略', '抽象')) or
--     (g.name = '谁是牛头人' and t.name in ('毛线', '聚会', '社交', '新手')) or
--     (g.name = '山屋惊魂' and t.name in ('美式', '剧情', '推理', '恐怖')) or
--     (g.name = '德国心脏病' and t.name in ('毛线', '家庭', '反应', '儿童')) or
--     (g.name = '三国杀' and t.name in ('中式', '卡牌', '策略', '蒸蒸日上')) or
--     (g.name = '怒海求生' and t.name in ('美式', '合作', '剧情')) or
--     (g.name = '大富翁' and t.name in ('家庭', '策略', '经济', '新手')) or
--     (g.name = '龙与地下城' and t.name in ('trpg', '角色扮演', '冒险')) or
--     (g.name = '克苏鲁的呼唤' and t.name in ('trpg', '角色扮演', '冒险')) or
--     (g.name = '战锤40k' and t.name in ('战棋', '策略', '角色扮演'));
insert into game (id, "name", description, image, rating, likes_count, extra_info) values
(1, '阿瓦隆', '正义与邪恶阵营对抗的社交推理游戏，通过投票完成任务决定胜负。', 'https://pic1.imgdb.cn/item/679b7bb9d0e0a243d4f8b0e4.jpg', 9.2, 23421, '{"minPlayers":5,"maxPlayers":10,"duration":45,"processSteps":[{"title":"确定角色","content":"6人局: \n好人阵营: 梅林, 派西维尔, 忠臣×2;\n坏人阵营: 莫德雷德, 刺客\n特别规则: 莫甘娜替换奥伯伦:cite[4]\n10人局: \n好人阵营: 梅林, 派西维尔, 忠臣×4, 湖中仙女; \n坏人阵营: 莫德雷德, 刺客, 莫甘娜, 奥伯伦\n特殊规则: 新增双面人角色，第3轮可能变阵营"},{"title":"角色分配","content":"随机分配角色卡，包含梅林、派西维尔、忠臣和莫德雷德阵营","image":"/images/process/avalon1.jpg"},{"title":"任务投票","content":"每轮由队长选择队员，全员投票决定是否执行该队伍。投票失败3次后进入强制轮:cite[4]","image":"/images/process/avalon2.jpg"},{"title":"任务执行","content":"被选队员暗投任务卡，若有1张反对票则任务失败（第4轮需2张反对票失败）:cite[7]","image":"/images/process/avalon3.jpg"},{"title":"刺杀阶段","content":"蓝方累计3次任务成功时，刺客需在限定时间内指认梅林身份:cite[4]"}],"roles":[{"id":1,"name":"梅林","team":"好人","description":"知晓所有坏人身份（除莫德雷德），需隐藏自己","complexity":4},{"id":2,"name":"派西维尔","team":"好人","description":"能识别梅林和莫甘娜，需判断真先知","complexity":3},{"id":3,"name":"莫德雷德","team":"坏人","description":"红方首领，梅林无法识别其身份:cite[7]","complexity":5},{"id":4,"name":"刺客","team":"坏人","description":"在蓝方获胜后可通过刺杀梅林翻盘","complexity":4},{"id":5,"name":"奥伯伦","team":"坏人","description":"隐狼角色，队友无法确认其身份","complexity":2}],"errataList":[{"id":1,"question":"刺客在游戏结束后是否可以查看梅林身份？","answer":"只有游戏失败时刺客可以指认梅林，成功后需公开验证","status":"confirmed","source":"官方规则书P23","reportCount":12},{"id":2,"question":"梅林能否直接透露坏人身份？","answer":"过度暴露会被刺客刺杀，需通过隐晦暗示引导队友","status":"confirmed","source":"社区FAQ","reportCount":34},{"id":3,"question":"8人局新增角色规则","answer":"加入湖中仙女角色，第二轮后可查验他人阵营（可撒谎）:cite[7]","status":"confirmed","source":"扩展规则v2.1"},{"id":4,"question":"新增规则王者之剑应该如何游玩","answer":"队长可指定王者之剑持有者，任务卡可被强制更换结果:cite[9]","status":"confirmed","source":"扩展规则v2.1"}]}'),
(2, '狼人杀', '狼人隐藏身份猎杀村民，村民通过推理找出狼人的社交游戏。', 'https://pic1.imgdb.cn/item/67ab9736d0e0a243d4fe75a5.jpg', 8.8, 45231, '{"minPlayers":6,"maxPlayers":18,"duration":60,"processSteps":[],"roles":[],"errataList":[]}'),
(3, 'UNO', '全球畅销的快速卡牌游戏，通过颜色/数字匹配与功能牌组合制造连锁反应，+2、反转、万能牌等特殊机制带来瞬息万变的局势。', 'https://pic1.imgdb.cn/item/67ab9693d0e0a243d4fe7596.jpg', 8.5, 34152, '{"minPlayers":2,"maxPlayers":10,"duration":15,"processSteps":[{"title":"分发起始","content":"每人发7张手牌，剩余牌堆作为抽牌区"},{"title":"核心回合","content":"匹配上家卡牌颜色/数字，或使用功能牌改变局势"},{"title":"特殊机制","content":"+4黑牌可指定颜色并让下家抽4张，反转牌改变出牌顺序"}],"roles":[],"errataList":[]}'),
(4, '卡坦岛：开拓者版', '2025新版重塑经典，新增航海扩展与3D地形模块，支持动态版图拼接与资源交易策略。文明建设游戏标杆，荣获2025年度最佳焕新桌游奖:cite[2]:cite[4]:cite[6]', 'https://pic1.imgdb.cn/item/67ab9693d0e0a243d4fe7595.jpg', 9.6, 45678, '{"minPlayers":3,"maxPlayers":6,"duration":90,"processSteps":[{"title":"版图创建","content":"随机拼接六边形地形板块，生成独特岛屿","image":"/images/catan_map.jpg"},{"title":"资源循环","content":"通过骰子结算资源产出，建造道路/村庄/城市"},{"title":"贸易博弈","content":"港口交易与玩家间资源谈判策略"}],"roles":[],"errataList":[{"id":401,"question":"新版航海扩展是否兼容旧版组件？","answer":"完全兼容，新增船只单位可替换原有道路模块","status":"confirmed","source":"官方FAQ"}]}'),
(5, '达芬奇密码', '数字推理巅峰之作，通过对手牌排列的试探性猜测，结合排除法破解4位数字密码。2023年推出双人快速对战变体规则', 'https://pic1.imgdb.cn/item/67ab9735d0e0a243d4fe75a1.jpg', 8.6, 12543, '{"minPlayers":2,"maxPlayers":4,"duration":20,"processSteps":[],"roles":[],"errataList":[]}'),
(6, '卡卡颂20周年纪念版', '新增公会与修道院扩展，磁吸式地形板块提升拼图体验。支持进阶计分规则：农夫占领区域计分方式优化', 'https://pic1.imgdb.cn/item/67ab9736d0e0a243d4fe75a4.jpg', 9.2, 32890, '{"minPlayers":2,"maxPlayers":5,"duration":45,"processSteps":[],"roles":[],"errataList":[]}'),
(7, '谁是牛头人', '虚实博弈社交游戏，通过道具卡牌制造信息差，2024年新增「黑暗交易」扩展包引入双重身份机制', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75a8.jpg', 8.7, 19234, '{"minPlayers":4,"maxPlayers":12,"duration":30,"processSteps":[],"roles":[],"errataList":[]}'),
(8, '山屋惊魂（新版）', '沉浸式恐怖剧情游戏，新增AR线索扫描功能，通过手机摄像头解谜。2024年推出「血色遗产」扩展包', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75ab.jpg', 9.3, 24567, '{"minPlayers":3,"maxPlayers":6,"duration":120,"processSteps":[{"title":"序幕阶段","content":"随机选择剧本，分配初始物品和秘密任务","image":"/images/betrayal_setup.jpg"},{"title":"探索阶段","content":"移动探索房间，触发预兆事件和超自然现象"},{"title":"转折时刻","content":"当预兆条件满足时，开启专属剧本的最终对抗"}],"roles":[],"errataList":[{"id":801,"question":"新版AR线索是否影响原有剧本平衡？","answer":"AR功能为可选模块，不影响核心规则平衡性","status":"confirmed","source":"设计师访谈"}]}'),
(9, '德国心脏病：动物乐园版', '新增30种濒危动物卡牌，配套生态知识问答机制，适合亲子环保教育', 'https://pic1.imgdb.cn/item/67ab9735d0e0a243d4fe75a3.jpg', 8.7, 34567, '{"minPlayers":2,"maxPlayers":8,"duration":15,"processSteps":[{"title":"卡牌分发","content":"每人获得5张动物卡，中央放置铃铛"},{"title":"快速匹配","content":"当出现相同动物时立即抢铃，正确者获得卡牌"},{"title":"知识问答","content":"特殊标记卡触发生态保护知识挑战"}],"roles":[],"errataList":[]}'),
(10, '三国杀：界限突破版', '重构经典武将技能体系，新增军争篇扩展和3D立体卡牌。支持在线身份场匹配功能', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75aa.jpg', 9, 89234, '{"minPlayers":4,"maxPlayers":10,"duration":45,"processSteps":[],"roles":[{"id":1001,"name":"界关羽","team":"坏人","description":"武圣：可将红色牌当【杀】使用或打出","complexity":3},{"id":1002,"name":"SP貂蝉","team":"中立","description":"离魂：出牌阶段可交换两名其他角色手牌","complexity":4}],"errataList":[{"id":1001,"question":"新版【闪电】判定规则是否有变化？","answer":"保持原有判定流程，新增数字传感器自动判定功能","status":"confirmed","reportCount":28}]}'),
(11, '怒海求生：风暴来袭扩展', '新增动态天气系统和救援直升机机制，支持6-12人大型团队合作模式', 'https://pic1.imgdb.cn/item/67ab977ed0e0a243d4fe75a9.jpg', 9.4, 23456, '{"minPlayers":4,"maxPlayers":12,"duration":90,"processSteps":[{"title":"角色分配","content":"随机抽取船员/叛徒身份"},{"title":"资源争夺","content":"每回合争夺有限淡水和食物"},{"title":"事件阶段","content":"触发随机海难事件（鲨鱼袭击/暴风雨等）"}],"roles":[],"errataList":[]}'),
(12, '大富翁：元宇宙版', '结合NFT地产交易系统，支持AR虚拟建筑投影。新增股票市场和加密货币玩法', 'https://pic1.imgdb.cn/item/67ab9735d0e0a243d4fe75a2.jpg', 8.8, 45678, '{"minPlayers":2,"maxPlayers":6,"duration":120,"processSteps":[],"roles":[],"errataList":[{"id":1201,"question":"实体版是否兼容虚拟资产？","answer":"通过扫码实现实体卡牌与区块链资产绑定","status":"pending","reportCount":15}]}'),
(13, '龙与地下城：龙焰传承', '2024年核心规则更新版，新增龙裔血脉系统和动态遭遇生成APP', 'https://pic1.imgdb.cn/item/67ab9693d0e0a243d4fe7594.jpg', 9.5, 123456, '{"minPlayers":3,"maxPlayers":8,"duration":240,"processSteps":[{"title":"创建角色","content":"选择种族职业，分配属性点"},{"title":"冒险准备","content":"DM设置剧情线索和遭遇事件"},{"title":"遭遇阶段","content":"投掷20面骰进行战斗/交涉判定"}],"roles":[],"errataList":[]}'),
(14, '克苏鲁的呼唤：电子之秘', '支持VR沉浸式模组，新增赛博朋克世界观扩展。包含AI守秘人辅助系统', 'https://pic1.imgdb.cn/item/67ab9692d0e0a243d4fe7593.jpg', 9, 34567, '{"minPlayers":2,"maxPlayers":6,"duration":180,"processSteps":[],"roles":[],"errataList":[]}'),
(15, '战锤40k：第十版', '全面更新战斗规则，简化AP计算系统。新增泰伦虫族基因吞噬者单位', 'https://pic1.imgdb.cn/item/67ab9692d0e0a243d4fe7592.jpg', 9.1, 45678, '{"minPlayers":2,"maxPlayers":2,"duration":180,"processSteps":[],"roles":[],"errataList":[{"id":1501,"question":"新单位移动距离计算方式？","answer":"采用新版测距轮系统，基础移动+骰子修正","status":"confirmed","source":"核心规则v10.1.2"}]}')
on conflict (id) do update set "name" = excluded."name", description = excluded.description, image = excluded.image, rating = excluded.rating, likes_count = excluded.likes_count, extra_info = excluded.extra_info;

insert into tag (id, "name", description) values
(1, '推理', ''),
(2, '社交', ''),
(3, '聚会', ''),
(4, '毛线', ''),
(5, '家庭', ''),
(6, '卡牌', ''),
(7, '德式', ''),
(8, '策略', ''),
(9, '资源管理', ''),
(10, '扩展', ''),
(11, '益智', ''),
(12, '解密', ''),
(13, '数字', ''),
(14, '抽象', ''),
(15, '拼图', ''),
(16, '角色', ''),
(17, '美式', ''),
(18, '剧情', ''),
(19, '恐怖', ''),
(20, 'AR', ''),
(21, '反应', ''),
(22, '教育', ''),
(23, '中式', ''),
(24, '电子化', ''),
(25, '合作', ''),
(26, '经济', ''),
(27, 'TRPG', ''),
(28, '角色扮演', ''),
(29, '冒险', ''),
(30, '科技', ''),
(31, '战棋', ''),
(32, '科幻', ''),
(33, '模型', '');

insert into game_tag_relation (game_id, tag_id) values
(1, 0),
(1, 1),
(1, 2),
(2, 0),
(2, 1),
(2, 2),
(3, 3),
(3, 4),
(3, 2),
(3, 5),
(4, 6),
(4, 7),
(4, 8),
(4, 9),
(5, 10),
(5, 4),
(5, 11),
(5, 12),
(6, 6),
(6, 7),
(6, 13),
(6, 14),
(7, 3),
(7, 2),
(7, 1),
(7, 15),
(8, 16),
(8, 17),
(8, 0),
(8, 18),
(8, 19),
(9, 3),
(9, 4),
(9, 20),
(9, 21),
(10, 22),
(10, 5),
(10, 7),
(10, 23),
(11, 16),
(11, 24),
(11, 17),
(11, 9),
(12, 4),
(12, 7),
(12, 25),
(12, 12),
(13, 26),
(13, 27),
(13, 28),
(13, 12),
(14, 26),
(14, 27),
(14, 18),
(14, 29),
(15, 30),
(15, 7),
(15, 31),
(15, 32);

